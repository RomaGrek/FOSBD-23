package main

import (
    "testing"
)

// Бенчмарк для формирования обратного индекса из текста из 1 слова
func BenchmarkInvertedIndexFromText1(b *testing.B) {
    content := "кошка"
    ii := NewInvertedIndex()

    b.ResetTimer() // Сбрасываем таймер перед началом тестирования

    for i := 0; i < b.N; i++ {
        ii.AddDocument(i, content)
    }
}

// Бенчмарк для формирования обратного индекса из текста из 5 слов
func BenchmarkInvertedIndexFromText5(b *testing.B) {
    content := "кошка сидит на крыше собака"
    ii := NewInvertedIndex()

    b.ResetTimer() // Сбрасываем таймер перед началом тестирования

    for i := 0; i < b.N; i++ {
        ii.AddDocument(i, content)
    }
}

// Бенчмарк для формирования обратного индекса из текста из 10 слов
func BenchmarkInvertedIndexFromText10(b *testing.B) {
    content := "кошка сидит на крыше собака лает на кошку крыша большая"
    ii := NewInvertedIndex()

    b.ResetTimer() // Сбрасываем таймер перед началом тестирования

    for i := 0; i < b.N; i++ {
        ii.AddDocument(i, content)
    }
}

// Бенчмарк для формирования обратного индекса из текста из 20 слов
func BenchmarkInvertedIndexFromText20(b *testing.B) {
    content := "кошка сидит на крыше верхом и поет как ворон сильно но не понимает что такое вездесущие око которое смотрит на нее"
    ii := NewInvertedIndex()

    b.ResetTimer() // Сбрасываем таймер перед началом тестирования

    for i := 0; i < b.N; i++ {
        ii.AddDocument(i, content)
    }
}

// Бенчмарк для функции построения сжатого обратного индекса из 10 документов
func BenchmarkCompressedInvertedIndexBuildFromInvertedIndex10(b *testing.B) {
    ii := NewInvertedIndex()
    // Заполняем индекс 10 документами
    for i := 0; i < 10; i++ {
        ii.AddDocument(i, "кошка сидит на крыше собака лает на кошку крыша большая")
    }

    cii := NewCompressedInvertedIndex()

    b.ResetTimer() // Сбрасываем таймер перед началом тестирования

    for i := 0; i < b.N; i++ {
        cii.BuildFromInvertedIndex(ii)
    }
}

// Бенчмарк для функции построения сжатого обратного индекса из 100 документов
func BenchmarkCompressedInvertedIndexBuildFromInvertedIndex100(b *testing.B) {
    ii := NewInvertedIndex()
    // Заполняем индекс 100 документами
    for i := 0; i < 100; i++ {
        ii.AddDocument(i, "кошка сидит на крыше собака лает на кошку крыша большая")
    }

    cii := NewCompressedInvertedIndex()

    b.ResetTimer() // Сбрасываем таймер перед началом тестирования

    for i := 0; i < b.N; i++ {
        cii.BuildFromInvertedIndex(ii)
    }
}

// Бенчмарк для функции построения сжатого обратного индекса из 1000 документов
func BenchmarkCompressedInvertedIndexBuildFromInvertedIndex1000(b *testing.B) {
    ii := NewInvertedIndex()
    // Заполняем индекс 1000 документами
    for i := 0; i < 1000; i++ {
        ii.AddDocument(i, "кошка сидит на крыше собака лает на кошку крыша большая")
    }

    cii := NewCompressedInvertedIndex()

    b.ResetTimer() // Сбрасываем таймер перед началом тестирования

    for i := 0; i < b.N; i++ {
        cii.BuildFromInvertedIndex(ii)
    }
}

// Бенчмарк для функции построения сжатого обратного индекса из 10000 документов
func BenchmarkCompressedInvertedIndexBuildFromInvertedIndex10000(b *testing.B) {
    ii := NewInvertedIndex()
    // Заполняем индекс 10000 документами
    for i := 0; i < 10000; i++ {
        ii.AddDocument(i, "кошка сидит на крыше собака лает на кошку крыша большая")
    }

    cii := NewCompressedInvertedIndex()

    b.ResetTimer() // Сбрасываем таймер перед началом тестирования

    for i := 0; i < b.N; i++ {
        cii.BuildFromInvertedIndex(ii)
    }
}

// Бенчмарк для функции булева поиска среди 10 документов
func BenchmarkCompressedInvertedIndexBooleanSearch10(b *testing.B) {
    ii := NewInvertedIndex()
    // Заполняем индекс 10 документами
    for i := 0; i < 10; i++ {
        ii.AddDocument(i, "кошка сидит на крыше собака лает на кошку крыша большая")
    }

    cii := NewCompressedInvertedIndex()
    cii.BuildFromInvertedIndex(ii)

    terms := []string{"кошка", "крыша"}

    b.ResetTimer() // Сбрасываем таймер перед началом тестирования

    for i := 0; i < b.N; i++ {
        cii.BooleanSearch(terms, "AND")
    }
}

// Бенчмарк для функции булева поиска среди 100 документов
func BenchmarkCompressedInvertedIndexBooleanSearch100(b *testing.B) {
    ii := NewInvertedIndex()
    // Заполняем индекс 100 документами
    for i := 0; i < 100; i++ {
        ii.AddDocument(i, "кошка сидит на крыше собака лает на кошку крыша большая")
    }

    cii := NewCompressedInvertedIndex()
    cii.BuildFromInvertedIndex(ii)

    terms := []string{"кошка", "крыша"}

    b.ResetTimer() // Сбрасываем таймер перед началом тестирования

    for i := 0; i < b.N; i++ {
        cii.BooleanSearch(terms, "AND")
    }
}

// Бенчмарк для функции булева поиска среди 1000 документов
func BenchmarkCompressedInvertedIndexBooleanSearch1000(b *testing.B) {
    ii := NewInvertedIndex()
    // Заполняем индекс 1000 документами
    for i := 0; i < 1000; i++ {
        ii.AddDocument(i, "кошка сидит на крыше собака лает на кошку крыша большая")
    }

    cii := NewCompressedInvertedIndex()
    cii.BuildFromInvertedIndex(ii)

    terms := []string{"кошка", "крыша"}

    b.ResetTimer() // Сбрасываем таймер перед началом тестирования

    for i := 0; i < b.N; i++ {
        cii.BooleanSearch(terms, "AND")
    }
}

// Бенчмарк для функции булева поиска среди 10000 документов
func BenchmarkCompressedInvertedIndexBooleanSearch10000(b *testing.B) {
    ii := NewInvertedIndex()
    // Заполняем индекс 10000 документами
    for i := 0; i < 10000; i++ {
        ii.AddDocument(i, "кошка сидит на крыше собака лает на кошку крыша большая")
    }

    cii := NewCompressedInvertedIndex()
    cii.BuildFromInvertedIndex(ii)

    terms := []string{"кошка", "крыша"}

    b.ResetTimer() // Сбрасываем таймер перед началом тестирования

    for i := 0; i < b.N; i++ {
        cii.BooleanSearch(terms, "AND")
    }
}

// Бенчмарк для функции pForDeltaEncode 10 элементов
func BenchmarkPForDeltaEncode10(b *testing.B) {
    // Список документов для кодирования
    postings := generateSequentialInts(10) // Генерируем список из 10 элементов

    b.ResetTimer() // Сбрасываем таймер перед началом тестирования

    for i := 0; i < b.N; i++ {
        _ = pForDeltaEncode(postings)
    }
}

// Бенчмарк для функции pForDeltaEncode 100 элементов
func BenchmarkPForDeltaEncode100(b *testing.B) {
    // Список документов для кодирования
    postings := generateSequentialInts(100) // Генерируем список из 100 элементов

    b.ResetTimer() // Сбрасываем таймер перед началом тестирования

    for i := 0; i < b.N; i++ {
        _ = pForDeltaEncode(postings)
    }
}

// Бенчмарк для функции pForDeltaEncode 1000 элементов
func BenchmarkPForDeltaEncode1000(b *testing.B) {
    // Список документов для кодирования
    postings := generateSequentialInts(1000) // Генерируем список из 1000 элементов

    b.ResetTimer() // Сбрасываем таймер перед началом тестирования

    for i := 0; i < b.N; i++ {
        _ = pForDeltaEncode(postings)
    }
}

// Бенчмарк для функции pForDeltaEncode 10000 элементов
func BenchmarkPForDeltaEncode10000(b *testing.B) {
    // Список документов для кодирования
    postings := generateSequentialInts(10000) // Генерируем список из 10000 элементов

    b.ResetTimer() // Сбрасываем таймер перед началом тестирования

    for i := 0; i < b.N; i++ {
        _ = pForDeltaEncode(postings)
    }
}

// Бенчмарк для функции pForDeltaDecode 10 элементов
func BenchmarkPForDeltaDecode10(b *testing.B) {
    // Список документов для кодирования
    postings := generateSequentialInts(10) // Генерируем список из 10 элементов
    encoded := pForDeltaEncode(postings)

    b.ResetTimer() // Сбрасываем таймер перед началом тестирования

    for i := 0; i < b.N; i++ {
        _ = pForDeltaDecode(encoded)
    }
}

// Бенчмарк для функции pForDeltaDecode 100 элементов
func BenchmarkPForDeltaDecode100(b *testing.B) {
    // Список документов для кодирования
    postings := generateSequentialInts(100) // Генерируем список из 100 элементов
    encoded := pForDeltaEncode(postings)

    b.ResetTimer() // Сбрасываем таймер перед началом тестирования

    for i := 0; i < b.N; i++ {
        _ = pForDeltaDecode(encoded)
    }
}

// Бенчмарк для функции pForDeltaDecode 1000 элементов
func BenchmarkPForDeltaDecode1000(b *testing.B) {
    // Список документов для кодирования
    postings := generateSequentialInts(1000) // Генерируем список из 1000 элементов
    encoded := pForDeltaEncode(postings)

    b.ResetTimer() // Сбрасываем таймер перед началом тестирования

    for i := 0; i < b.N; i++ {
        _ = pForDeltaDecode(encoded)
    }
}

// Бенчмарк для функции pForDeltaDecode 10000 элементов
func BenchmarkPForDeltaDecode10000(b *testing.B) {
    // Список документов для кодирования
    postings := generateSequentialInts(10000) // Генерируем список из 10000 элементов
    encoded := pForDeltaEncode(postings)

    b.ResetTimer() // Сбрасываем таймер перед началом тестирования

    for i := 0; i < b.N; i++ {
        _ = pForDeltaDecode(encoded)
    }
}

// Вспомогательная функция для генерации последовательных чисел
func generateSequentialInts(n int) []int {
    seq := make([]int, n)
    for i := 0; i < n; i++ {
        seq[i] = i + 1
    }
    return seq
}