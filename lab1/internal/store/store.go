package store

import (
	"errors"
	"fmt"
	"math"
	"os"

	bkt "debildb/internal/bucket"

	"go.uber.org/zap"
)

const (
	defaultGlobalDepth int = 1
	defaultLocalDepth  int = 1
	lenKV                  = 1365
)

var (
	pageSize = os.Getpagesize()
)

// Главня аструктура хранилища
type Store struct {
	dirList     []Directory
	globalDepth int
	pathToDB    string
	endOffset   int
	log         *zap.Logger
}

// NewStore - инициализирует хранилище с базовыми значениями
func NewStore(pathDB string, log *zap.Logger) *Store {
	store := &Store{
		pathToDB:    pathDB,
		globalDepth: defaultGlobalDepth,
		log:         log,
	}

	err := store.InitDefaultDirectoryList() // инициализация начального списка из двух директорий и двух бакетов
	if err != nil {
		log.Fatal("new store", zap.Error(err))
	}

	return store
}

// Структура директории
type Directory struct {
	index      byte
	bucket     *bkt.Bucket
	localDepth int
}

// Функция начальной инициализации списка диреткорий и бакетов
func (s *Store) InitDefaultDirectoryList() error {
	bkt1, endOffset, err := bkt.CreateBucket(s.pathToDB, s.endOffset)	// Создание первого бакета
	if err != nil {
		return fmt.Errorf("new default directory list: %w", err)
	}
	s.endOffset = endOffset // обновляем указатель на конец бд

	bkt2, endOffset, err := bkt.CreateBucket(s.pathToDB, s.endOffset) // Создание второго бакета
	if err != nil {
		return fmt.Errorf("new default directory list: %w", err)
	}
	s.endOffset = endOffset // обновляем указатель на конец бд

	s.dirList = []Directory{
		{
			index:      0,
			bucket:     bkt1,
			localDepth: defaultLocalDepth,
		},
		{
			index:      1,
			bucket:     bkt2,
			localDepth: defaultLocalDepth,
		},
	}

	s.log.Info("Successful init store")

	return nil
}

// Функция загрузки значения
func (s *Store) SetValue(key, value string) error {
	index := getDirID(key, s.globalDepth) // получаем id директории по ключу

	if int(index) > len(s.dirList) { // проверка на то, что id директории валидный
		return fmt.Errorf("invalid index")
	}

	dir := s.dirList[int(index)]
	err := dir.bucket.PutValue(&bkt.KV{Key: key, Val: value}) // Пытаемся положить значение
	if err != nil {
		if errors.Is(err, bkt.ErrBucketIsFull) { // Если получаем ошибку того, что бакет переполнен, значит нужен или глобальный ресайз или сплит
			if dir.localDepth < s.globalDepth { // если local depth меньше чем global depth - значит можем просто сплитануть бакет без глобального ресайза
				err := s.splitBucket(dir, dir.bucket) // сплитуем бакет
				if err != nil {
					return fmt.Errorf("store - SetValue: %w", err)
				}

				err = s.SetValue(key, value) // Снова пытаемся положить значение
				if err != nil {
					return fmt.Errorf("recircive call set value 1: %w", err)
				}
				return nil
			}
			// если global depth == local depth значит требуется глобальный ресайз
			s.globalResize() // выполняем глобальный ресайз
			err = s.SetValue(key, value) // Заново пытаемся положить значнеие (на практике будет опять ошибка и уже в этот раз мы попадем на сплит бакета, в процессе которого уже значение положиться нормально)
			if err != nil {
				return fmt.Errorf("recircive call set value 2: %w", err)
			}
			return nil
		}
		return fmt.Errorf("store - SetValue: %w", err)
	}

	s.log.Info("Save data", zap.Int("directory", int(dir.index)), zap.Int("bucket", dir.bucket.GetBucketID()), zap.String("key", key), zap.String("value", value))

	return nil
}

// Функция глобального рейсайза директорий
func (s *Store) globalResize() {
	s.log.Info("global resize")
	newGlobalDepth := s.globalDepth + 1                     // увеличиваем globalDepth
	countDir := int(math.Pow(2.0, float64(newGlobalDepth))) // считываем кол-во директорий, которое будет после ресайза
	newDirList := make([]Directory, countDir)

	for i := 0; i < countDir; i++ { // цикл формирования новых директорий
		index := byte(i)                  // получаем индекс директории (идентификатор)
		mask := (1 << s.globalDepth) - 1  // считаем маску для последних битов, количество каторых равно старому globalDepth
		targetIndex := index & byte(mask) // применяем маску что бы определить идентфикаторы диреткорий в старом списке директорий что бы понять на какой бакет должна указывать директория

		newDirList[i] = Directory{
			index:      byte(i),
			bucket:     s.dirList[targetIndex].bucket,
			localDepth: s.dirList[targetIndex].localDepth,
		}
	}

	s.dirList = newDirList
	s.globalDepth = newGlobalDepth
}

// Функция разделения бакета
func (s *Store) splitBucket(oldDir Directory, oldBucket *bkt.Bucket) error {
	s.log.Info("split bucket", zap.Int("bucket", oldBucket.GetBucketID()))

	newBkt, endOffset, err := bkt.CreateBucket(s.pathToDB, s.endOffset) // Создаем новый бакет
	if err != nil {
		return fmt.Errorf("split bucket: %w", err)
	}
	s.endOffset = endOffset

	records, err := oldBucket.GetBucketValues() // Получаем значения из бакета для дальнейшего их перераспределения
	if err != nil {
		return fmt.Errorf("split bucket: %w", err)
	}

	err = oldBucket.SetBucketIsEmpty() // Отчищаем бакет, у которого только что вытащили значения
	if err != nil {
		return fmt.Errorf("error in split - empty: %w", err)
	}

	// Перестановка указателей
	firstIndex := -1
	mask := (1 << (oldDir.localDepth + 1)) - 1 // определяем маску, что бы определить какие директории указывали на бакет в старой версии списка директорий
	for i := 0; i < len(s.dirList); i++ {
		if s.dirList[i].bucket == oldBucket { // если мы находим директорию, которая указывала на бакет, и это впервые то помечаем ее как первый тип (т к в итоге у нас всегда будет два типа окончание битов)
			if firstIndex == -1 {
				firstIndex = int(s.dirList[i].index) & mask
			} else { // в случае если это не первый раз смотрим первый тип это или второй и в завимисоти от этого прикрепляем указатель на новый бакет или нет. Если первый тип - старый бакет. Второй тип - новый.
				if ((s.dirList[i].index) & byte(mask)) != byte(firstIndex) {
					s.dirList[i].bucket = newBkt
				}
			}
			s.dirList[i].localDepth++ // В любом случае если мы находим нужную директорию нужно увеличить local depth т к по факту у бакета это уже новая версия
		}
	}

	for _, kv := range records { // Заново заполянем значения, которые до этого достали из переполненного бакета
		err := s.SetValue(kv.Key, kv.Val)
		if err != nil {
			return fmt.Errorf("error in split - set value: %w", err)
		}
	}

	s.log.Info("split bucket completed")

	return nil
}

// Функция получения значнеия по ключу
func (s *Store) GetValue(key string) (string, error) {
	index := getDirID(key, s.globalDepth) // высчитываем id дирекотрии где должна находиться запись

	if int(index) > len(s.dirList) { // проверяем на всякий что индекс валиден
		return "", fmt.Errorf("invalid index")
	}

	dir := s.dirList[int(index)] // получаем нужную директорию
	kv, err := dir.bucket.GetValue(key) // получаем значение 
	if err != nil {
		return "", fmt.Errorf("store get value: %w", err)
	}

	s.log.Info("Get data", zap.Int("directory", int(dir.index)), zap.Int("bucket", dir.bucket.GetBucketID()))

	return kv.Val, nil
}
