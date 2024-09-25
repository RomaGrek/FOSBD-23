package bucket

import (
	"debildb/internal/parser"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/edsrzf/mmap-go"
)

const (
	pageSize = 4096

	lenKV = 1365
)

var (
	ErrBucketIsFull = errors.New("bucket is full, need resize") 
	ErrKeyNotFound  = errors.New("key not found")
)

type Bucket struct {
	offset int
	pathDB string
}

type KV struct {
	Key string
	Val string
}

// Функция создания бакета.
func CreateBucket(pathToDB string, endOffset int) (*Bucket, int, error) {
	file, err := os.OpenFile(pathToDB, os.O_RDWR|os.O_CREATE, 0755) // открытие файла
	if err != nil {
		log.Fatalf("CreateBucket open file: %s", err.Error())
	}
	defer file.Close()

	reserv := make([]byte, os.Getpagesize())
	_, err = file.WriteAt(reserv, int64(endOffset)) // запись нулевых байтов (резервируем бакет)
	if err != nil {
		return nil, -1, fmt.Errorf("create bucket - write at: %w", err)
	}

	newEndOffset := endOffset + pageSize // считаем новый указатель на конец бд

	return &Bucket{
		offset: endOffset,
		pathDB: pathToDB,
	}, newEndOffset, nil
}

// Функция получения значения из бакета по ключу
func (b *Bucket) GetValue(key string) (*KV, error) {
	bktData, err := b.getBucket()				// получаем сам бакет
	if err != nil {
		return nil, fmt.Errorf("error bucket Get Value: %w", err)
	}

	for i := 0; i < int(bktData[0]); i++ {	// начинаем итерироваться по бакту. Первый байт бакета содержит количество элементов в бакете на данный момент
		offset := 1 + i*lenKV				// считаем смещение
		curKV := bktData[offset : offset+lenKV] // получаем текущий слайс бакт относящийся к нужной записи

		curKey, curVal, err := parser.UnmarshalKV(curKV) // анмаршалим его
		if err != nil {
			return nil, fmt.Errorf("error bucket Get Value: %w", err)
		}

		if key == curKey {
			return &KV{
				Key: curKey,
				Val: curVal,
			}, nil
		}
	}

	return nil, fmt.Errorf("error bucket Get Value: %w", ErrKeyNotFound)
}

// Функция на загрузку значения в бакет
func (b *Bucket) PutValue(kv *KV) error {
	bktData, err := b.getBucket()	// Получаем бакет
	if err != nil {
		return fmt.Errorf("error bucket Put Value: %w", err)
	}

	if bktData[0] > 2 {	// Проверяем что в бакете есть место
		return ErrBucketIsFull
	}

	kvData, err := parser.MarshalKV(kv.Key, kv.Val)	// маршалим запись
	if err != nil {
		return fmt.Errorf("error bucket Put Value: %w", err)
	}

	err = b.setBucket(int(bktData[0]), kvData) // кладем в бакет
	if err != nil {
		return fmt.Errorf("error bucket Put Value: %w", err)
	}

	return nil
}

// Функция получения всех значений внутри бакета. Используется при сплите бакета, когда нужно перераспределить значнеия между двумя бакетами после разделения.
func (b *Bucket) GetBucketValues() ([]KV, error) {
	bktData, err := b.getBucket() // Получаем бакет
	if err != nil {
		return nil, fmt.Errorf("error bucket Get Bucket Values: %w", err)
	}

	result := make([]KV, 0, int(bktData[0]))
	for i := 0; i < int(bktData[0]); i++ { // начинаем итерироваться по каждому элементу
		offset := 1 + i*lenKV // считаем смещение
		curKV := bktData[offset : offset+lenKV] // достаем слайс байт опередленной записи

		curKey, curVal, err := parser.UnmarshalKV(curKV) // анмаршаллим запись
		if err != nil {
			return nil, fmt.Errorf("error bucket Get Value: %w", err)
		}

		result = append(result, KV{Key: curKey, Val: curVal})
	}

	return result, nil
}

// Функция обнуления бакета (нужно при сплите бакета), когда после того как достали элементы нужно его почистить.
func (b *Bucket) SetBucketIsEmpty() error {
	db, err := os.OpenFile(b.pathDB, os.O_RDWR|os.O_CREATE, 0755)	// Открываем файл
	if err != nil {
		return fmt.Errorf("get bucket - open file: %w", err)
	}
	defer db.Close()

	data, err := mmap.MapRegion(db, pageSize, mmap.RDWR, 0, int64(b.offset)) // мапим нужную страницу (страницу - потому что размер бакета равен размеру страницы)
	if err != nil {
		return fmt.Errorf("create bucket - map region: %w", err)
	}

	copy(data, make([]byte, 4096))	// заполняем слайс байт нулевыми байтами

	if err = data.Flush(); err != nil { // флашим все на диск
		return fmt.Errorf("create bucket - flush: %w", err)
	}

	if err := data.Unmap(); err != nil { // размапливаем страницу
		log.Fatalf("CreateBucket Map: %s", err.Error())
	}

	return nil
}

// Функция получения байт бакета
func (b *Bucket) getBucket() ([]byte, error) {
	db, err := os.OpenFile(b.pathDB, os.O_RDWR|os.O_CREATE, 0755) // Открываем файл
	if err != nil {
		return nil, fmt.Errorf("get bucket - open file: %w", err)
	}
	defer db.Close()

	data, err := mmap.MapRegion(db, os.Getpagesize(), mmap.RDWR, 0, int64(b.offset)) // мапим нужную страницу (страницу - потому что размер бакета равен размеру страницы)
	if err != nil {
		return nil, fmt.Errorf("get bucket - map region: %w", err)
	}

	resultData := make([]byte, len(data))
	copy(resultData, data) // копируем данные к себе в буфер

	if err := data.Unmap(); err != nil { // размапливаем страницу
		return nil, fmt.Errorf("get bucket - unmap: %w", err)
	}

	return resultData, nil
}

// Функция сохранения бакета на диск (обновленного)
func (b *Bucket) setBucket(index int, dataKV []byte) error {
	db, err := os.OpenFile(b.pathDB, os.O_RDWR|os.O_CREATE, 0755) // открываем файл
	if err != nil {
		return fmt.Errorf("get bucket - open file: %w", err)
	}
	defer db.Close()

	data, err := mmap.MapRegion(db, pageSize, mmap.RDWR, 0, int64(b.offset)) // мапим нужную страницу (страницу - потому что размер бакета равен размеру страницы)
	if err != nil {
		return fmt.Errorf("create bucket - map region: %w", err)
	}

	offset := 1 + index*lenKV // считаем смещение для обновления определенного элемента
	copy(data[offset:offset+lenKV], dataKV) // переносим значения
	data[0] = data[0] + 1 // увеличиваем счетчик количества элементов в бакете

	if err = data.Flush(); err != nil { // флашим данные на диск
		return fmt.Errorf("create bucket - flush: %w", err)
	}

	if err := data.Unmap(); err != nil { // размапливаем память
		log.Fatalf("CreateBucket Map: %s", err.Error())
	}

	return nil
}

// Функция расчета бакет ID (по факту индекс бакета)
func (b *Bucket) GetBucketID() int {
	return b.offset / pageSize
}