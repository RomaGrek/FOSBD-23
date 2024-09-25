package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Тест для функции AddDocument в InvertedIndex
func TestInvertedIndexAddDocument(t *testing.T) {
	ii := NewInvertedIndex()
	ii.AddDocument(1, "кошка сидит на крыше")
	ii.AddDocument(2, "собака лает на кошку")
	ii.AddDocument(3, "крыша большая")

	require.Equal(t, []int{1}, ii.Index["кошка"], "Терм 'кошка' должен присутствовать в документе 1")
	require.Equal(t, []int{3}, ii.Index["крыша"], "Терм 'крыша' должен присутствовать в документе 3")
	require.Equal(t, []int{1, 2}, ii.Index["на"], "Терм 'на' должен присутствовать в документах 1 и 2")
}

// Тест для функции BooleanSearch в CompressedInvertedIndex
func TestCompressedInvertedIndexBooleanSearch(t *testing.T) {
	ii := NewInvertedIndex()
	ii.AddDocument(1, "кошка сидит на крыше")
	ii.AddDocument(2, "собака лает на кошку")
	ii.AddDocument(3, "крыша большая")

	cii := NewCompressedInvertedIndex()
	cii.BuildFromInvertedIndex(ii)

	// Проверка операции AND
	result := cii.BooleanSearch([]string{"кошка", "крыше"}, "AND")
	require.Equal(t, []int{1}, result, "Ожидаемый результат для 'кошка' AND 'крыша': [1]")

	// Проверка операции OR
	result = cii.BooleanSearch([]string{"кошка", "крыша"}, "OR")
	require.Equal(t, []int{1, 3}, result, "Ожидаемый результат для 'кошка' OR 'крыша': [1, 3]")

	// Проверка операции NOT
	result = cii.BooleanSearch([]string{"крыша", "собака"}, "NOT")
	require.Equal(t, []int{3}, result, "Ожидаемый результат для 'крыша' NOT 'собака': [3]")
}

// Тест для функции BooleanSearch в SearchEngine
func TestSearchEngineBooleanSearch(t *testing.T) {
	engine := NewSearchEngine()
	engine.AddDocument(1, "кошка сидит на крыше")
	engine.AddDocument(2, "собака лает на кошку")
	engine.AddDocument(3, "крыша большая")

	engine.BuildCompressedIndex()

	// Проверка операции AND
	result := engine.BooleanSearch([]string{"кошка", "крыше"}, "AND")
	require.Equal(t, []int{1}, result, "Ожидаемый результат для 'кошка' AND 'крыша': [1]")

	// Проверка операции OR
	result = engine.BooleanSearch([]string{"кошка", "крыша"}, "OR")
	require.Equal(t, []int{1, 3}, result, "Ожидаемый результат для 'кошка' OR 'крыша': [1, 3]")

	// Проверка операции NOT
	result = engine.BooleanSearch([]string{"крыша", "собака"}, "NOT")
	require.Equal(t, []int{3}, result, "Ожидаемый результат для 'крыша' NOT 'собака': [3]")

	// Проверка поиска с термом, который не существует
	result = engine.BooleanSearch([]string{"крыша", "кот"}, "AND")
	require.Equal(t, []int{}, result, "Ожидаемый результат для 'крыша' AND 'кот': []")
}
