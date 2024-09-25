package store

import (
	"crypto/sha256"
)

// Функция получения ID дирекотрии где должен быть элемент
// Опрелеояется по count последним битам хэша
func getDirID(key string, count int) byte {
	hashBytes := hash(key)

	lastByte := hashBytes[len(hashBytes)-1]	
	mask := byte((1 << count) - 1) // считаем маску для выборки нужны битов

	return lastByte & mask // применяем маску
}

// Функция подсчета хэша ключа
func hash(key string) []byte {
	hasher := sha256.New()
	hasher.Write([]byte(key))

	return hasher.Sum(nil)
}
