package parser

import (
	"bytes"
	"fmt"
	"math/bits"
)

// 1 element = 128 + 1237 B
// 1 key = 1 B (len) + 127 B (data)
// 1 value = 2 B (len) + 1235 B (data)
const (
	lenKey   = 128
	lenValue = 1237
)

const (
	keyType = iota
	valType
)

// Сериализует пару ключ-значение,
// где ключ занимает 128 байт
// значение занимает 1237 байт.
// Возвращает слайс байт размером 1365 байт
func MarshalKV(key, val string) ([]byte, error) {
	keyBytes, err := serializeString(key, keyType)	// Сериализация ключа
	if err != nil {
		return nil, err
	}
	valBytes, err := serializeString(val, valType) // Сериализация значения
	if err != nil {
		return nil, err
	}

	keyBytes = padSlice(keyBytes, lenKey)
	valBytes = padSlice(valBytes, lenValue)

	result := make([]byte, lenKey+lenValue)
	copy(result[0:], keyBytes)
	copy(result[lenKey:], valBytes)

	return result, nil
}

// Функция дополнения нулями в случае если значение меньше целевого размера
func padSlice(data []byte, targetLen int) []byte {
	if len(data) < targetLen {
		padding := make([]byte, targetLen-len(data))
		data = append(data, padding...)
	}
	return data
}

// Сериализация числа, используется для сериализации длины строки
func serializeUint(value uint64) ([]byte, error) {
	bf := bytes.NewBuffer(make([]byte, 0))
	bitsLen := bits.Len64(value)
	bytesLen, remainder := bitsLen/7, bitsLen%7
	if remainder > 0 {
		bytesLen++
	}

	for i := 0; i < bytesLen; i++ {
		curByte := byte((value>>(7*i))&0x7f | 0x80)
		if i == (bytesLen - 1) {
			curByte &= 0x7f
		}
		bf.WriteByte(curByte)
	}

	return bf.Bytes(), nil
}

// Сериализация строки - состоит из сериализации числа (длины строки) и сериализации самой строки
func serializeString(value string, typeVal int) ([]byte, error) {
	var bf *bytes.Buffer
	if typeVal == keyType {
		bf = bytes.NewBuffer(make([]byte, 0, 128))
	} else {
		bf = bytes.NewBuffer(make([]byte, 0, 1237))
	}

	bLen, err := serializeUint(uint64(len(value)))
	if err != nil {
		return nil, fmt.Errorf("error in SerializeString: %w", err)
	}
	bf.Write(bLen)
	bf.Write([]byte(value))

	return bf.Bytes(), nil
}

// Парсинг записи ключ-значение
// На вход подается слайс байт, длинной в 1365 байт
func UnmarshalKV(dataKV []byte) (string, string, error) {
	bufKey := bytes.NewBuffer(dataKV[:128])
	bufValue := bytes.NewBuffer(dataKV[128:])
	key, err := deserializeString(bufKey)
	if err != nil {
		return "", "", err
	}

	val, err := deserializeString(bufValue)
	if err != nil {
		return "", "", err
	}

	return key, val, nil

}

// Функция дессериализации числа (длина строки)
func deserializeUint(bf *bytes.Buffer) (uint64, error) {
	res := uint64(0)
	i := 0
	for {
		curByte, err := bf.ReadByte()
		if err != nil {
			return 0, fmt.Errorf("error read byte")
		}
		res |= uint64(curByte&0x7F) << (7 * i)
		if (curByte & 0x80) == 0 {
			break
		}
		i++
	}
	return res, nil
}

// Функция для десериализации строки
func deserializeString(bf *bytes.Buffer) (string, error) {
	countBytes, err := deserializeUint(bf) // десериализация числа (длины строки)
	if err != nil {
		return "", fmt.Errorf("error in DeserializeString: %w", err)
	}
	strBytes := make([]byte, countBytes)
	_, err = bf.Read(strBytes)
	if err != nil {
		return "", fmt.Errorf("error in _deserializeString: %w", err)
	}
	str := string(strBytes)
	return str, nil
}
