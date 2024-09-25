package main

import (
    "testing"
)

// Бенчмарк для функции Insert c 10 точками
func BenchmarkInsert10(b *testing.B) {
    points := generateRandomPoints(10) // Генерируем 10 случайных точек для вставки
    root := NewNode(points[0], 0)

    b.ResetTimer() // Сбрасываем таймер перед началом бенчмарка

    for i := 0; i < b.N; i++ {
        // Вставляем все точки в дерево
        for j := 1; j < len(points); j++ {
            root.Insert(points[j], 0)
        }
    }
}

// Бенчмарк для функции Insert с 100 точками
func BenchmarkInsert100(b *testing.B) {
    points := generateRandomPoints(100) // Генерируем 100 случайных точек для вставки
    root := NewNode(points[0], 0)

    b.ResetTimer() // Сбрасываем таймер перед началом бенчмарка

    for i := 0; i < b.N; i++ {
        // Вставляем все точки в дерево
        for j := 1; j < len(points); j++ {
            root.Insert(points[j], 0)
        }
    }
}

// Бенчмарк для функции Insert с 1000 точками
func BenchmarkInsert1000(b *testing.B) {
    points := generateRandomPoints(1000) // Генерируем 1000 случайных точек для вставки
    root := NewNode(points[0], 0)

    b.ResetTimer() // Сбрасываем таймер перед началом бенчмарка

    for i := 0; i < b.N; i++ {
        // Вставляем все точки в дерево
        for j := 1; j < len(points); j++ {
            root.Insert(points[j], 0)
        }
    }
}

// Бенчмарк для функции Insert с 10000 точками
func BenchmarkInsert10000(b *testing.B) {
    points := generateRandomPoints(10000) // Генерируем 10000 случайных точек для вставки
    root := NewNode(points[0], 0)

    b.ResetTimer() // Сбрасываем таймер перед началом бенчмарка

    for i := 0; i < b.N; i++ {
        // Вставляем все точки в дерево
        for j := 1; j < len(points); j++ {
            root.Insert(points[j], 0)
        }
    }
}

// Бенчмарк для функции FindNearest с 10 точками
func BenchmarkFindNearest10(b *testing.B) {
    points := generateRandomPoints(10)
    root := NewNode(points[0], 0)
    for i := 1; i < len(points); i++ {
        root.Insert(points[i], 0)
    }

    target := Point{Coordinates: []float64{50, 50}}

    b.ResetTimer()

    for i := 0; i < b.N; i++ {
        root.FindNearest(target, 0)
    }
}

// Бенчмарк для функции FindNearest с 100 точками
func BenchmarkFindNearest100(b *testing.B) {
    points := generateRandomPoints(100)
    root := NewNode(points[0], 0)
    for i := 1; i < len(points); i++ {
        root.Insert(points[i], 0)
    }

    target := Point{Coordinates: []float64{50, 50}}

    b.ResetTimer()

    for i := 0; i < b.N; i++ {
        root.FindNearest(target, 0)
    }
}

// Бенчмарк для функции FindNearest с 1000 точками
func BenchmarkFindNearest1000(b *testing.B) {
    points := generateRandomPoints(1000)
    root := NewNode(points[0], 0)
    for i := 1; i < len(points); i++ {
        root.Insert(points[i], 0)
    }

    target := Point{Coordinates: []float64{50, 50}}

    b.ResetTimer()

    for i := 0; i < b.N; i++ {
        root.FindNearest(target, 0)
    }
}

// Бенчмарк для функции FindNearest с 10000 точками
func BenchmarkFindNearest10000(b *testing.B) {
    points := generateRandomPoints(10000)
    root := NewNode(points[0], 0)
    for i := 1; i < len(points); i++ {
        root.Insert(points[i], 0)
    }

    target := Point{Coordinates: []float64{50, 50}}

    b.ResetTimer()

    for i := 0; i < b.N; i++ {
        root.FindNearest(target, 0)
    }
}

// Бенчмарк для функции RangeSearch с 10 точками
func BenchmarkRangeSearch10(b *testing.B) {
    points := generateRandomPoints(10)
    root := NewNode(points[0], 0)
    for i := 1; i < len(points); i++ {
        root.Insert(points[i], 0)
    }

    targetMin := Point{Coordinates: []float64{30, 30}}
    targetMax := Point{Coordinates: []float64{70, 70}}

    b.ResetTimer()

    for i := 0; i < b.N; i++ {
        results := []Point{}
        root.RangeSearch(targetMin, targetMax, 0, &results)
    }
}

// Бенчмарк для функции RangeSearch с 100 точками
func BenchmarkRangeSearch100(b *testing.B) {
    points := generateRandomPoints(100)
    root := NewNode(points[0], 0)
    for i := 1; i < len(points); i++ {
        root.Insert(points[i], 0)
    }

    targetMin := Point{Coordinates: []float64{30, 30}}
    targetMax := Point{Coordinates: []float64{70, 70}}

    b.ResetTimer()

    for i := 0; i < b.N; i++ {
        results := []Point{}
        root.RangeSearch(targetMin, targetMax, 0, &results)
    }
}

// Бенчмарк для функции RangeSearch с 1000 точками
func BenchmarkRangeSearch1000(b *testing.B) {
    points := generateRandomPoints(1000)
    root := NewNode(points[0], 0)
    for i := 1; i < len(points); i++ {
        root.Insert(points[i], 0)
    }

    targetMin := Point{Coordinates: []float64{30, 30}}
    targetMax := Point{Coordinates: []float64{70, 70}}

    b.ResetTimer()

    for i := 0; i < b.N; i++ {
        results := []Point{}
        root.RangeSearch(targetMin, targetMax, 0, &results)
    }
}

// Бенчмарк для функции RangeSearch с 10000 точками
func BenchmarkRangeSearch10000(b *testing.B) {
    points := generateRandomPoints(10000)
    root := NewNode(points[0], 0)
    for i := 1; i < len(points); i++ {
        root.Insert(points[i], 0)
    }

    targetMin := Point{Coordinates: []float64{30, 30}}
    targetMax := Point{Coordinates: []float64{70, 70}}

    b.ResetTimer()

    for i := 0; i < b.N; i++ {
        results := []Point{}
        root.RangeSearch(targetMin, targetMax, 0, &results)
    }
}

// Вспомогательная функция для генерации случайных точек
func generateRandomPoints(num int) []Point {
    points := make([]Point, num)
    for i := 0; i < num; i++ {
        points[i] = Point{Coordinates: []float64{float64(i), float64(i * i % 100)}}
    }
    return points
}