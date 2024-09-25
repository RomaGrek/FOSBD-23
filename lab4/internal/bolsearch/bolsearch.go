package main

import (
    "fmt"
    "sort"
    "strings"
)

// InvertedIndex представляет обратный индекс для хранения термов и списка документов
type InvertedIndex struct {
    Index map[string][]int // Карта термов и их связей с документами (постинг-листы)
}

// NewInvertedIndex создает новый пустой обратный индекс
func NewInvertedIndex() *InvertedIndex {
    return &InvertedIndex{Index: make(map[string][]int)} // Инициализация пустого индекса
}

// AddDocument добавляет документ в обратный индекс
func (ii *InvertedIndex) AddDocument(docID int, content string) {
    words := strings.Fields(content) // Разбиение текста документа на отдельные слова
    for _, word := range words {
        word = strings.ToLower(word) // Приведение всех слов к нижнему регистру для унификации
        if _, exists := ii.Index[word]; !exists {
            ii.Index[word] = []int{} // Создание пустого постинг-листа для нового терма
        }
        // Добавление docID в постинг-лист терма, если его там еще нет
        if len(ii.Index[word]) == 0 || ii.Index[word][len(ii.Index[word])-1] != docID {
            ii.Index[word] = append(ii.Index[word], docID)
        }
    }
}

// PrintIndex выводит содержимое обратного индекса
func (ii *InvertedIndex) PrintIndex() {
    fmt.Println("Построенный обратный индекс:")
    for term, postings := range ii.Index {
        fmt.Printf("Терм: '%s', Документы: %v\n", term, postings) // Печать терма и связанных с ним документов
    }
}

// CompressedInvertedIndex объединяет обратный индекс и дельта-кодирование для оптимизации хранения
type CompressedInvertedIndex struct {
    CompressedIndex map[string][]int // Карта термов на сжатые постинг-листы
}

// NewCompressedInvertedIndex создает новый пустой сжатый обратный индекс
func NewCompressedInvertedIndex() *CompressedInvertedIndex {
    return &CompressedInvertedIndex{CompressedIndex: make(map[string][]int)} // Инициализация пустого сжатого индекса
}

// BuildFromInvertedIndex строит сжатый обратный индекс из обычного обратного индекса
func (cii *CompressedInvertedIndex) BuildFromInvertedIndex(ii *InvertedIndex) {
    for term, postings := range ii.Index {
        cii.CompressedIndex[term] = pForDeltaEncode(postings) // Сжатие каждого постинг-листа
    }
}

// pForDeltaEncode выполняет сжатие постинг-листа с использованием дельта-кодирования
func pForDeltaEncode(postings []int) []int {
    if len(postings) == 0 {
        return postings // Если список пуст, возвращаем его без изменений
    }
    encoded := make([]int, len(postings))
    encoded[0] = postings[0] // Первый элемент остается без изменений
    for i := 1; i < len(postings); i++ {
        encoded[i] = postings[i] - postings[i-1] // Вычисление разности между соседними элементами
    }
    return encoded
}

// pForDeltaDecode выполняет декодирование постинг-листа, сжатого с использованием pForDelta
func pForDeltaDecode(encoded []int) []int {
    if len(encoded) == 0 {
        return encoded // Если список пуст, возвращаем его без изменений
    }
    decoded := make([]int, len(encoded))
    decoded[0] = encoded[0] // Первый элемент остается без изменений
    for i := 1; i < len(encoded); i++ {
        decoded[i] = decoded[i-1] + encoded[i] // Восстановление оригинального списка
    }
    return decoded
}

// BooleanSearch выполняет булев поиск по термам с использованием логического оператора
func (cii *CompressedInvertedIndex) BooleanSearch(terms []string, operator string) []int {
    if len(terms) == 0 {
        return []int{} // Если термы не заданы, возвращаем пустой список
    }

    // Декодирование первого постинг-листа для начального результата
    result := pForDeltaDecode(cii.CompressedIndex[terms[0]])

    for _, term := range terms[1:] {
        decodedTermList := pForDeltaDecode(cii.CompressedIndex[term]) // Декодирование постинг-листа текущего терма
        if operator == "AND" {
            result = intersect(result, decodedTermList) // Пересечение (AND)
        } else if operator == "OR" {
            result = union(result, decodedTermList) // Объединение (OR)
        } else if operator == "NOT" {
            result = difference(result, decodedTermList) // Разность (NOT)
        }
    }
    return result
}

// intersect возвращает пересечение двух списков документов
func intersect(a, b []int) []int {
    result := []int{}
    i, j := 0, 0
    // Одновременный проход по обоим спискам для нахождения общих элементов
    for i < len(a) && j < len(b) {
        if a[i] == b[j] {
            result = append(result, a[i])
            i++
            j++
        } else if a[i] < b[j] {
            i++
        } else {
            j++
        }
    }
    return result
}

// union возвращает объединение двух списков документов
func union(a, b []int) []int {
    result := append(a, b...)
    sort.Ints(result) // Сортировка объединенного списка
    return unique(result) // Удаление дубликатов
}

// difference возвращает разницу между двумя списками документов (a - b)
func difference(a, b []int) []int {
    result := []int{}
    m := make(map[int]bool)
    for _, num := range b {
        m[num] = true // Добавляем все элементы второго списка в карту для быстрого поиска
    }
    for _, num := range a {
        if !m[num] { // Добавляем только те элементы, которых нет во втором списке
            result = append(result, num)
        }
    }
    return result
}

// unique возвращает список уникальных элементов
func unique(a []int) []int {
    result := []int{}
    m := make(map[int]bool)
    for _, num := range a {
        if !m[num] { // Добавляем элемент только если его еще нет в карте
            m[num] = true
            result = append(result, num)
        }
    }
    return result
}

// Полный интерфейс для работы с сжатым обратным индексом
type SearchEngine struct {
    ii  *InvertedIndex            // Обычный обратный индекс
    cii *CompressedInvertedIndex  // Сжатый обратный индекс
}

// NewSearchEngine создает новый поисковый движок
func NewSearchEngine() *SearchEngine {
    return &SearchEngine{
        ii:  NewInvertedIndex(),
        cii: NewCompressedInvertedIndex(),
    }
}

// AddDocument добавляет документ в поисковый движок
func (se *SearchEngine) AddDocument(docID int, content string) {
    se.ii.AddDocument(docID, content) // Добавление документа в обычный обратный индекс
}

// BuildCompressedIndex строит сжатый индекс
func (se *SearchEngine) BuildCompressedIndex() {
    se.cii.BuildFromInvertedIndex(se.ii) // Создание сжатого индекса на основе обычного
}

// PrintInvertedIndex выводит обычный обратный индекс
func (se *SearchEngine) PrintInvertedIndex() {
    se.ii.PrintIndex() // Печать обычного обратного индекса
}

// BooleanSearch выполняет булев поиск по сжатому индексу
func (se *SearchEngine) BooleanSearch(terms []string, operator string) []int {
    return se.cii.BooleanSearch(terms, operator) // Выполнение булевого поиска по сжатому индексу
}

func main() {
    // Создание нового поискового движка
    engine := NewSearchEngine()

    // Добавление документов в поисковый движок
    engine.AddDocument(1, "кошка сидит на крыше")
    engine.AddDocument(2, "собака лает на кошку")
    engine.AddDocument(3, "крыша большая")

    // Печать обычного обратного индекса
    engine.PrintInvertedIndex()

    // Построение сжатого индекса
    engine.BuildCompressedIndex()

    // Выполнение булевых поисков
    fmt.Println("\nПоиск 'кошка' AND 'крыша':")
    result := engine.BooleanSearch([]string{"кошка", "крыша"}, "AND")
    fmt.Println(result)

    fmt.Println("\nПоиск 'собака' OR 'крыша':")
    result = engine.BooleanSearch([]string{"собака", "крыша"}, "OR")
    fmt.Println(result)

    fmt.Println("\nПоиск 'крыша' NOT 'собака':")
    result = engine.BooleanSearch([]string{"крыша", "собака"}, "NOT")
    fmt.Println(result)

	fmt.Println("\nПоиск 'крыша' AND 'на':")
	result = engine.BooleanSearch([]string{"крыша", "на"}, "OR")
	fmt.Println(result)
}
