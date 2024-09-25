package main

import (
	"math/rand"
	"testing"
	"time"
)

// Бенчмарк для функции Insert в RadixTree с 10 случайными словами
func BenchmarkRadixTreeInsert10(b *testing.B) {
	tree := NewRadixTree()
	words := generateRandomWords(10, 6) // Генерируем 10 случайных слов длиной 6 символов

	b.ResetTimer() // Сбрасываем таймер перед началом тестирования

	for i := 0; i < b.N; i++ {
		for _, word := range words {
			tree.Insert(word)
		}
	}
}

// Бенчмарк для функции Insert в RadixTree с 100 случайными словами
func BenchmarkRadixTreeInsert100(b *testing.B) {
	tree := NewRadixTree()
	words := generateRandomWords(100, 6) // Генерируем 100 случайных слов длиной 6 символов

	b.ResetTimer() // Сбрасываем таймер перед началом тестирования

	for i := 0; i < b.N; i++ {
		for _, word := range words {
			tree.Insert(word)
		}
	}
}

// Бенчмарк для функции Insert в RadixTree с 1000 случайными словами
func BenchmarkRadixTreeInsert1000(b *testing.B) {
	tree := NewRadixTree()
	words := generateRandomWords(1000, 6) // Генерируем 1000 случайных слов длиной 6 символов

	b.ResetTimer() // Сбрасываем таймер перед началом тестирования

	for i := 0; i < b.N; i++ {
		for _, word := range words {
			tree.Insert(word)
		}
	}
}

// Бенчмарк для функции Insert в RadixTree с 10000 случайными словами
func BenchmarkRadixTreeInsert10000(b *testing.B) {
	tree := NewRadixTree()
	words := generateRandomWords(10000, 6) // Генерируем 10000 случайных слов длиной 6 символов

	b.ResetTimer() // Сбрасываем таймер перед началом тестирования

	for i := 0; i < b.N; i++ {
		for _, word := range words {
			tree.Insert(word)
		}
	}
}

// Бенчмарк для функции Search в RadixTree с 10 случайными словами
func BenchmarkRadixTreeSearch10(b *testing.B) {
	tree := NewRadixTree()
	words := generateRandomWords(10, 6) // Генерируем 10 случайных слов длиной 6 символов
	for _, word := range words {
		tree.Insert(word)
	}

	b.ResetTimer() // Сбрасываем таймер перед началом тестирования

	for i := 0; i < b.N; i++ {
		for _, word := range words {
			tree.Search(word) // Поиск существующего слова
		}
		tree.Search("notpresent") // Поиск слова, которого нет в дереве
	}
}

// Бенчмарк для функции Search в RadixTree с 100 случайными словами
func BenchmarkRadixTreeSearch100(b *testing.B) {
	tree := NewRadixTree()
	words := generateRandomWords(100, 6) // Генерируем 100 случайных слов длиной 6 символов
	for _, word := range words {
		tree.Insert(word)
	}

	b.ResetTimer() // Сбрасываем таймер перед началом тестирования

	for i := 0; i < b.N; i++ {
		for _, word := range words {
			tree.Search(word) // Поиск существующего слова
		}
		tree.Search("notpresent") // Поиск слова, которого нет в дереве
	}
}

// Бенчмарк для функции Search в RadixTree с 1000 случайными словами
func BenchmarkRadixTreeSearch(b *testing.B) {
	tree := NewRadixTree()
	words := generateRandomWords(1000, 6) // Генерируем 1000 случайных слов длиной 6 символов
	for _, word := range words {
		tree.Insert(word)
	}

	b.ResetTimer() // Сбрасываем таймер перед началом тестирования

	for i := 0; i < b.N; i++ {
		for _, word := range words {
			tree.Search(word) // Поиск существующего слова
		}
		tree.Search("notpresent") // Поиск слова, которого нет в дереве
	}
}

// Бенчмарк для функции Search в RadixTree с 10000 случайными словами
func BenchmarkRadixTreeSearch10000(b *testing.B) {
	tree := NewRadixTree()
	words := generateRandomWords(10000, 6) // Генерируем 10000 случайных слов длиной 6 символов
	for _, word := range words {
		tree.Insert(word)
	}

	b.ResetTimer() // Сбрасываем таймер перед началом тестирования

	for i := 0; i < b.N; i++ {
		for _, word := range words {
			tree.Search(word) // Поиск существующего слова
		}
		tree.Search("notpresent") // Поиск слова, которого нет в дереве
	}
}

// Бенчмарк для функции StartsWith в RadixTree с 10 случайными словами
func BenchmarkRadixTreeStartsWith10(b *testing.B) {
	tree := NewRadixTree()
	words := generateRandomWords(10, 6) // Генерируем 10 случайных слов длиной 6 символов
	for _, word := range words {
		tree.Insert(word)
	}

	b.ResetTimer() // Сбрасываем таймер перед началом тестирования

	for i := 0; i < b.N; i++ {
		for _, word := range words {
			prefix := word[:2] // Используем первые два символа в качестве префикса
			tree.StartsWith(prefix)
		}
	}
}

// Бенчмарк для функции StartsWith в RadixTree с 100 случайными словами
func BenchmarkRadixTreeStartsWith100(b *testing.B) {
	tree := NewRadixTree()
	words := generateRandomWords(100, 6) // Генерируем 100 случайных слов длиной 6 символов
	for _, word := range words {
		tree.Insert(word)
	}

	b.ResetTimer() // Сбрасываем таймер перед началом тестирования

	for i := 0; i < b.N; i++ {
		for _, word := range words {
			prefix := word[:2] // Используем первые два символа в качестве префикса
			tree.StartsWith(prefix)
		}
	}
}

// Бенчмарк для функции StartsWith в RadixTree с 1000 случайными словами
func BenchmarkRadixTreeStartsWith1000(b *testing.B) {
	tree := NewRadixTree()
	words := generateRandomWords(1000, 6) // Генерируем 1000 случайных слов длиной 6 символов
	for _, word := range words {
		tree.Insert(word)
	}

	b.ResetTimer() // Сбрасываем таймер перед началом тестирования

	for i := 0; i < b.N; i++ {
		for _, word := range words {
			prefix := word[:2] // Используем первые два символа в качестве префикса
			tree.StartsWith(prefix)
		}
	}
}

// Бенчмарк для функции StartsWith в RadixTree с 10000 случайными словами
func BenchmarkRadixTreeStartsWith10000(b *testing.B) {
	tree := NewRadixTree()
	words := generateRandomWords(10000, 6) // Генерируем 10000 случайных слов длиной 6 символов
	for _, word := range words {
		tree.Insert(word)
	}

	b.ResetTimer() // Сбрасываем таймер перед началом тестирования

	for i := 0; i < b.N; i++ {
		for _, word := range words {
			prefix := word[:2] // Используем первые два символа в качестве префикса
			tree.StartsWith(prefix)
		}
	}
}

// Бенчмарк для функции WordsWithPrefix в RadixTree с 10 случайными словами
func BenchmarkRadixTreeWordsWithPrefix10(b *testing.B) {
	tree := NewRadixTree()
	words := generateRandomWords(10, 6) // Генерируем 10 случайных слов длиной 6 символов
	for _, word := range words {
		tree.Insert(word)
	}

	b.ResetTimer() // Сбрасываем таймер перед началом тестирования

	for i := 0; i < b.N; i++ {
		for _, word := range words {
			prefix := word[:2] // Используем первые два символа в качестве префикса
			tree.WordsWithPrefix(prefix)
		}
	}
}

// Бенчмарк для функции WordsWithPrefix в RadixTree с 100 случайными словами
func BenchmarkRadixTreeWordsWithPrefix100(b *testing.B) {
	tree := NewRadixTree()
	words := generateRandomWords(100, 6) // Генерируем 100 случайных слов длиной 6 символов
	for _, word := range words {
		tree.Insert(word)
	}

	b.ResetTimer() // Сбрасываем таймер перед началом тестирования

	for i := 0; i < b.N; i++ {
		for _, word := range words {
			prefix := word[:2] // Используем первые два символа в качестве префикса
			tree.WordsWithPrefix(prefix)
		}
	}
}

// Бенчмарк для функции WordsWithPrefix в RadixTree с 1000 случайными словами
func BenchmarkRadixTreeWordsWithPrefix1000(b *testing.B) {
	tree := NewRadixTree()
	words := generateRandomWords(1000, 6) // Генерируем 1000 случайных слов длиной 6 символов
	for _, word := range words {
		tree.Insert(word)
	}

	b.ResetTimer() // Сбрасываем таймер перед началом тестирования

	for i := 0; i < b.N; i++ {
		for _, word := range words {
			prefix := word[:2] // Используем первые два символа в качестве префикса
			tree.WordsWithPrefix(prefix)
		}
	}
}

// Бенчмарк для функции WordsWithPrefix в RadixTree с 10000 случайными словами
func BenchmarkRadixTreeWordsWithPrefix10000(b *testing.B) {
	tree := NewRadixTree()
	words := generateRandomWords(10000, 6) // Генерируем 10000 случайных слов длиной 6 символов
	for _, word := range words {
		tree.Insert(word)
	}

	b.ResetTimer() // Сбрасываем таймер перед началом тестирования

	for i := 0; i < b.N; i++ {
		for _, word := range words {
			prefix := word[:2] // Используем первые два символа в качестве префикса
			tree.WordsWithPrefix(prefix)
		}
	}
}

// Генерация случайного слова заданной длины
func generateRandomWord(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyz"
	rand.Seed(time.Now().UnixNano())
	word := make([]byte, length)
	for i := range word {
		word[i] = charset[rand.Intn(len(charset))]
	}
	return string(word)
}

// Генерация массива случайных слов
func generateRandomWords(numWords, length int) []string {
	words := make([]string, numWords)
	for i := 0; i < numWords; i++ {
		words[i] = generateRandomWord(length)
	}
	return words
}
