package main

import (
	"math"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

// Тестируем вставку точек в KD-дерево
func TestInsert(t *testing.T) {
	points := []Point{
		{[]float64{2, 3}},
		{[]float64{5, 4}},
		{[]float64{9, 6}},
		{[]float64{4, 7}},
		{[]float64{8, 1}},
		{[]float64{7, 2}},
	}

	// Создаем корневой узел
	root := NewNode(points[0], 0)

	// Вставляем остальные точки в дерево
	for i := 1; i < len(points); i++ {
		root.Insert(points[i], 0)
	}

	// Проверяем наличие всех точек
	require.True(t, treeContains(root, points[1]), "Точка %v не была найдена в KD-дереве после вставки", points[1])
	require.True(t, treeContains(root, points[5]), "Точка %v не была найдена в KD-дереве после вставки", points[5])
}

// Вспомогательная функция для проверки наличия точки в KD-дереве
func treeContains(root *Node, target Point) bool {
	if root == nil {
		return false
	}
	if reflect.DeepEqual(root.Point, target) {
		return true
	}
	return treeContains(root.Left, target) || treeContains(root.Right, target)
}

// Тестируем поиск ближайшей точки в KD-дереве
func TestFindNearest(t *testing.T) {
	points := []Point{
		{[]float64{2, 3}},
		{[]float64{5, 4}},
		{[]float64{9, 6}},
		{[]float64{4, 7}},
		{[]float64{8, 1}},
		{[]float64{7, 2}},
	}

	// Создаем корневой узел
	root := NewNode(points[0], 0)

	// Вставляем остальные точки в дерево
	for i := 1; i < len(points); i++ {
		root.Insert(points[i], 0)
	}

	target := Point{[]float64{9, 2}}
	expected := Point{[]float64{8, 1}}

	nearest, _ := root.FindNearest(target, 0)

	require.NotNil(t, nearest, "Ближайший узел не найден")
	require.Equal(t, expected, nearest.Point, "Ожидаемая ближайшая точка %v, полученная %v", expected.Coordinates, nearest.Point.Coordinates)
}

// Тестируем диапазонный поиск в KD-дереве
func TestRangeSearch(t *testing.T) {
	points := []Point{
		{[]float64{2, 3}},
		{[]float64{5, 4}},
		{[]float64{9, 6}},
		{[]float64{4, 7}},
		{[]float64{8, 1}},
		{[]float64{7, 2}},
	}

	// Создаем корневой узел
	root := NewNode(points[0], 0)

	// Вставляем остальные точки в дерево
	for i := 1; i < len(points); i++ {
		root.Insert(points[i], 0)
	}

	targetMin := Point{Coordinates: []float64{3, 2}}
	targetMax := Point{Coordinates: []float64{10, 8}}
	expectedResults := []Point{
		{[]float64{5, 4}},
		{[]float64{9, 6}},
		{[]float64{4, 7}},
		{[]float64{7, 2}},
	}

	results := []Point{}
	root.RangeSearch(targetMin, targetMax, 0, &results)

	require.Len(t, results, len(expectedResults), "Ожидаемое количество точек %v, полученное количество %v", len(expectedResults), len(results))

	for _, expected := range expectedResults {
		require.True(t, contains(results, expected), "Ожидаемая точка %v не найдена в результатах диапазонного поиска", expected.Coordinates)
	}
}

// Вспомогательная функция для проверки наличия точки в результатах поиска
func contains(points []Point, target Point) bool {
	for _, p := range points {
		if reflect.DeepEqual(p, target) {
			return true
		}
	}
	return false
}

// Тестируем расчет расстояния между двумя точками
func TestDistance(t *testing.T) {
	pointA := Point{[]float64{1, 2}}
	pointB := Point{[]float64{4, 6}}
	expected := math.Sqrt(25) // 5.0

	result := distance(pointA, pointB)

	require.Equal(t, expected, result, "Ожидаемое расстояние %v, полученное %v", expected, result)
}
