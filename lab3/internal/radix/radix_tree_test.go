package main

import (
    "testing"

    "github.com/stretchr/testify/require"
)

// Тест для функции Insert и Search
func TestRadixTreeInsertAndSearch(t *testing.T) {
    tree := NewRadixTree()

    // Вставляем слова в дерево
    words := []string{"cat", "car", "cart", "dog", "door", "dot"}
    for _, word := range words {
        tree.Insert(word)
    }

    // Проверяем точное совпадение
    require.True(t, tree.Search("cat"), "Должен найти слово 'cat'")
    require.True(t, tree.Search("car"), "Должен найти слово 'car'")
    require.True(t, tree.Search("cart"), "Должен найти слово 'cart'")
    require.True(t, tree.Search("dog"), "Должен найти слово 'dog'")
    require.True(t, tree.Search("door"), "Должен найти слово 'door'")
    require.True(t, tree.Search("dot"), "Должен найти слово 'dot'")
    require.False(t, tree.Search("can"), "Не должен найти слово 'can'")
    require.False(t, tree.Search("do"), "Не должен найти слово 'do'")
}

// Тест для функции StartsWith
func TestRadixTreeStartsWith(t *testing.T) {
    tree := NewRadixTree()

    // Вставляем слова в дерево
    words := []string{"cat", "car", "cart", "dog", "door", "dot"}
    for _, word := range words {
        tree.Insert(word)
    }

    // Проверяем префиксы
    require.True(t, tree.StartsWith("ca"), "Должен найти слова, начинающиеся на 'ca'")
    require.True(t, tree.StartsWith("car"), "Должен найти слова, начинающиеся на 'car'")
    require.True(t, tree.StartsWith("do"), "Должен найти слова, начинающиеся на 'do'")
    require.False(t, tree.StartsWith("dor"), "Должен найти слова, начинающиеся на 'dor'")
    require.False(t, tree.StartsWith("dorm"), "Не должен найти слова, начинающиеся на 'dorm'")
    require.False(t, tree.StartsWith("caz"), "Не должен найти слова, начинающиеся на 'caz'")
    require.False(t, tree.StartsWith("dan"), "Не должен найти слова, начинающиеся на 'dan'")
}

// Тест для функции WordsWithPrefix
func TestRadixTreeWordsWithPrefix(t *testing.T) {
    tree := NewRadixTree()

    // Вставляем слова в дерево
    words := []string{"cat", "car", "cart", "dog", "door", "dot"}
    for _, word := range words {
        tree.Insert(word)
    }

    // Проверяем слова по префиксу
    require.ElementsMatch(t, tree.WordsWithPrefix("ca"), []string{"cat", "car", "cart"}, "Должен найти все слова с префиксом 'ca'")
    require.ElementsMatch(t, tree.WordsWithPrefix("do"), []string{"dog", "door", "dot"}, "Должен найти все слова с префиксом 'do'")
    require.ElementsMatch(t, tree.WordsWithPrefix("car"), []string{"car", "cart"}, "Должен найти все слова с префиксом 'car'")
    require.ElementsMatch(t, tree.WordsWithPrefix("dog"), []string{"dog"}, "Должен найти все слова с префиксом 'dog'")
    require.Empty(t, tree.WordsWithPrefix("can"), "Не должен найти слова с префиксом 'can'")
    require.Empty(t, tree.WordsWithPrefix("dar"), "Не должен найти слова с префиксом 'dar'")
}