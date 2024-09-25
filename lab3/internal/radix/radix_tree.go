package main

import (
    "fmt"
    "strings"
)

// RadixTreeNode представляет узел в Radix дереве
type RadixTreeNode struct {
    Prefix   string                     // Префикс узла
    Children map[string]*RadixTreeNode  // Дети узла, ключи — префиксы
    IsLeaf   bool                       // Флаг, указывающий, является ли узел листом (конечная точка слова)
}

// RadixTree представляет само дерево
type RadixTree struct {
    Root *RadixTreeNode // Корневой узел дерева
}

// NewRadixTree создает новое пустое Radix дерево
func NewRadixTree() *RadixTree {
    return &RadixTree{
        Root: &RadixTreeNode{
            Children: make(map[string]*RadixTreeNode), // Инициализируем пустой набор дочерних узлов
        },
    }
}

// Insert добавляет слово в Radix дерево
func (tree *RadixTree) Insert(word string) {
    currentNode := tree.Root // Начинаем с корневого узла
    remainingWord := word    // Остаток слова, который еще не вставлен

    for len(remainingWord) > 0 {
        found := false
        for key, child := range currentNode.Children {
            commonPrefix := commonPrefix(remainingWord, key) // Находим общий префикс между текущим словом и ключом узла

            if len(commonPrefix) > 0 {
                if commonPrefix == key {
                    // Полное совпадение, продолжаем с ребенком
                    currentNode = child
                    remainingWord = remainingWord[len(commonPrefix):] // Уменьшаем слово на длину совпадающего префикса
                    found = true
                    break
                } else {
                    // Частичное совпадение, разделяем узел
                    newChild := &RadixTreeNode{
                        Prefix:   key[len(commonPrefix):], // Оставшаяся часть ключа
                        Children: child.Children,          // Наследуем детей от оригинального узла
                        IsLeaf:   child.IsLeaf,            // Копируем состояние листа
                    }
                    // Создаем новый узел с общим префиксом
                    currentNode.Children[commonPrefix] = &RadixTreeNode{
                        Prefix:   commonPrefix,
                        Children: map[string]*RadixTreeNode{newChild.Prefix: newChild}, // Добавляем новый дочерний узел
                        IsLeaf:   false,
                    }
                    delete(currentNode.Children, key) // Удаляем старый ключ

                    if len(remainingWord) == len(commonPrefix) {
                        // Если остаток слова равен общему префиксу, узел становится листом
                        currentNode.Children[commonPrefix].IsLeaf = true
                    } else {
                        // Добавляем оставшуюся часть слова как новый узел
                        currentNode.Children[commonPrefix].Children[remainingWord[len(commonPrefix):]] = &RadixTreeNode{
                            Prefix:   remainingWord[len(commonPrefix):],
                            IsLeaf:   true,
                            Children: make(map[string]*RadixTreeNode),
                        }
                    }
                    return
                }
            }
        }

        if !found {
            // Если ни один префикс не найден, создаем новый узел
            currentNode.Children[remainingWord] = &RadixTreeNode{
                Prefix:   remainingWord,
                IsLeaf:   true,
                Children: make(map[string]*RadixTreeNode),
            }
            return
        }
    }
    currentNode.IsLeaf = true // Помечаем текущий узел как лист, если слово полностью вставлено
}

// Search выполняет точный поиск слова в дереве
func (tree *RadixTree) Search(word string) bool {
    node := tree.Root
    remainingWord := word

    for len(remainingWord) > 0 {
        found := false
        for key, child := range node.Children {
            if strings.HasPrefix(remainingWord, key) { // Проверяем, начинается ли оставшаяся часть слова с ключа
                remainingWord = remainingWord[len(key):] // Убираем часть ключа из слова
                node = child
                found = true
                break
            }
        }
        if !found {
            return false // Если ни один узел не соответствует, возвращаем false
        }
    }
    return node.IsLeaf // Возвращаем true, если узел является листом (окончанием слова)
}

// StartsWith проверяет, если слово начинается с заданного префикса
func (tree *RadixTree) StartsWith(prefix string) bool {
    node := tree.Root
    remainingPrefix := prefix

    for len(remainingPrefix) > 0 {
        found := false
        for key, child := range node.Children {
            if strings.HasPrefix(remainingPrefix, key) { // Проверяем, начинается ли префикс с ключа узла
                remainingPrefix = remainingPrefix[len(key):] // Убираем совпадающую часть
                node = child
                found = true
                break
            }
        }
        if !found {
            return false // Префикс не найден, возвращаем false
        }
    }
    return true // Префикс найден
}

// WordsWithPrefix возвращает все слова, начинающиеся с заданного префикса
func (tree *RadixTree) WordsWithPrefix(prefix string) []string {
    node := tree.Root
    remainingPrefix := prefix

    for len(remainingPrefix) > 0 {
        found := false
        for key, child := range node.Children {
            if strings.HasPrefix(remainingPrefix, key) {
                remainingPrefix = remainingPrefix[len(key):] // Переходим к следующей части префикса
                node = child
                found = true
                break
            }
        }
        if !found {
            return []string{} // Префикс не найден, возвращаем пустой список
        }
    }

    // Собираем все слова с данного узла
    return collectWords(node, prefix)
}

// collectWords собирает все слова с данного узла
func collectWords(node *RadixTreeNode, prefix string) []string {
    words := []string{}
    if node.IsLeaf {
        words = append(words, prefix) // Если узел — лист, добавляем его префикс в список слов
    }
    for key, child := range node.Children {
        words = append(words, collectWords(child, prefix+key)...) // Рекурсивно добавляем слова дочерних узлов
    }
    return words
}

// commonPrefix находит общий префикс двух строк
func commonPrefix(str1, str2 string) string {
    minLen := len(str1)
    if len(str2) < minLen {
        minLen = len(str2)
    }
    i := 0
    for i < minLen && str1[i] == str2[i] { // Ищем первый несоответствующий символ
        i++
    }
    return str1[:i] // Возвращаем общий префикс
}

// main демонстрирует использование RadixTree
func main() {
    tree := NewRadixTree()

    // Добавляем слова в дерево
    words := []string{"cat", "car", "cart", "dog", "door", "dot"}
    for _, word := range words {
        tree.Insert(word) // Вставляем каждое слово в дерево
    }

    // Демонстрация операций поиска
    fmt.Println("Поиск 'cat':", tree.Search("cat"))   // true - слово "cat" присутствует
    fmt.Println("Поиск 'can':", tree.Search("can"))   // false - слово "can" отсутствует
    fmt.Println("Поиск 'dog':", tree.Search("dog"))   // true - слово "dog" присутствует
    fmt.Println("Поиск 'door':", tree.Search("door")) // true - слово "door" присутствует
    fmt.Println("Префикс 'do':", tree.StartsWith("do")) // true - есть слова, начинающиеся с "do"
    fmt.Println("Префикс 'dor':", tree.StartsWith("dor")) // true - есть слова, начинающиеся с "dor"

    // Поиск слов по префиксу
    fmt.Println("Слова с префиксом 'do':", tree.WordsWithPrefix("do")) // ["dog", "door", "dot"]
    fmt.Println("Слова с префиксом 'ca':", tree.WordsWithPrefix("ca")) // ["cat", "car", "cart"]
}