package main

import (
    "fmt"
    "math"
)

// Point представляет точку в многомерном пространстве
type Point struct {
    Coordinates []float64
}

// Node представляет узел KD-дерева
type Node struct {
    Point  Point
    Left   *Node
    Right  *Node
    Axis   int
}

// NewNode создает новый узел KD-дерева
func NewNode(point Point, axis int) *Node {
    return &Node{Point: point, Axis: axis} // Инициализация узла с заданной точкой и осью
}

// Insert добавляет точку в KD-дерево
func (node *Node) Insert(point Point, depth int) {
    axis := depth % len(point.Coordinates) // Вычисляем ось разбиения в зависимости от текущей глубины
    if point.Coordinates[axis] < node.Point.Coordinates[axis] { // Сравниваем координату точки с координатой узла по текущей оси
        if node.Left == nil {
            node.Left = NewNode(point, axis) // Создаем новый левый узел, если он отсутствует
        } else {
            node.Left.Insert(point, depth+1) // Рекурсивно вставляем точку в левое поддерево
        }
    } else {
        if node.Right == nil {
            node.Right = NewNode(point, axis) // Создаем новый правый узел, если он отсутствует
        } else {
            node.Right.Insert(point, depth+1) // Рекурсивно вставляем точку в правое поддерево
        }
    }
}

// FindNearest ищет ближайшую точку к заданной
func (node *Node) FindNearest(target Point, depth int) (*Node, float64) {
    if node == nil {
        return nil, math.Inf(1) // Если узел пустой, возвращаем бесконечность как расстояние
    }

    axis := depth % len(target.Coordinates) // Определяем ось разбиения

    var next, other *Node
    if target.Coordinates[axis] < node.Point.Coordinates[axis] {
        next, other = node.Left, node.Right // Определяем, в какое поддерево двигаться дальше
    } else {
        next, other = node.Right, node.Left
    }

    bestNode, _ := next.FindNearest(target, depth+1) // Ищем ближайшую точку в выбранном поддереве

    best, bestDist := closerDistance(
        node,
        bestNode,
        target,
    )

    // Проверка другой ветки, если текущий лучший результат не является наилучшим
    if math.Abs(node.Point.Coordinates[axis]-target.Coordinates[axis]) < bestDist {
        bestNode, _ := other.FindNearest(target, depth+1)
        
        best, bestDist = closerDistance(
            best,
            bestNode,
            target,
        )
    }

    return best, bestDist
}

// RangeSearch выполняет поиск всех точек в пределах заданного диапазона
func (node *Node) RangeSearch(targetMin, targetMax Point, depth int, results *[]Point) {
    if node == nil {
        return
    }

    // Проверяем, находится ли точка в диапазоне
    isInRange := true
    for i := 0; i < len(targetMin.Coordinates); i++ {
        if node.Point.Coordinates[i] < targetMin.Coordinates[i] || node.Point.Coordinates[i] > targetMax.Coordinates[i] {
            isInRange = false // Если хотя бы одна координата вне диапазона, то точка не подходит
            break
        }
    }

    if isInRange {
        *results = append(*results, node.Point) // Добавляем точку в результаты поиска, если она в диапазоне
    }

    axis := depth % len(targetMin.Coordinates)

    // Проверяем левую ветвь, если есть вероятность пересечения диапазона
    if targetMin.Coordinates[axis] <= node.Point.Coordinates[axis] {
        node.Left.RangeSearch(targetMin, targetMax, depth+1, results)
    }

    // Проверяем правую ветвь, если есть вероятность пересечения диапазона
    if targetMax.Coordinates[axis] >= node.Point.Coordinates[axis] {
        node.Right.RangeSearch(targetMin, targetMax, depth+1, results)
    }
}

// closerDistance возвращает ближайший узел и расстояние между узлами
func closerDistance(node *Node, bestNode *Node, target Point) (*Node, float64) {
    if bestNode == nil {
        return node, distance(node.Point, target) // Если лучший узел не найден, возвращаем текущий узел и его расстояние
    }
    d1 := distance(node.Point, target) // Расстояние от текущего узла до цели
    d2 := distance(bestNode.Point, target) // Расстояние от лучшего узла до цели
    if d1 < d2 {
        return node, d1 // Возвращаем текущий узел, если он ближе
    }
    return bestNode, d2 // Возвращаем лучший узел, если он ближе
}

// distance вычисляет евклидово расстояние между двумя точками
func distance(a, b Point) float64 {
    var sum float64
    for i := range a.Coordinates { // Проходим по всем координатам и считаем сумму квадратов разностей
        diff := a.Coordinates[i] - b.Coordinates[i]
        sum += diff * diff
    }
    return math.Sqrt(sum) // Возвращаем квадратный корень из суммы квадратов разностей
}

// main демонстрирует использование KD-дерева и диапазонного поиска
func main() {
    // Пример точек для KD-дерева
    points := []Point{
        {[]float64{2, 3}},
        {[]float64{5, 4}},
        {[]float64{9, 6}},
        {[]float64{4, 7}},
        {[]float64{8, 1}},
        {[]float64{7, 2}},
    }

    // Создаем корневой узел KD-дерева
    root := NewNode(points[0], 0)

    // Вставляем остальные точки в дерево
    for i := 1; i < len(points); i++ {
        root.Insert(points[i], 0) // Вставка каждой точки с начальной глубиной 0
    }

    // Целевая точка для поиска ближайшего соседа
    target := Point{[]float64{9, 2}}

    // Поиск ближайшего соседа
    nearest, dist := root.FindNearest(target, 0)

    // Печать результатов
    fmt.Printf("Ближайшая точка к %v: %v на расстоянии %v\n", target.Coordinates, nearest.Point.Coordinates, dist)

    // Диапазонный поиск: ищем точки в прямоугольнике [3, 2] - [9, 5]
    targetMin := Point{Coordinates: []float64{3, 2}}
    targetMax := Point{Coordinates: []float64{10, 8}}
    results := []Point{}
    root.RangeSearch(targetMin, targetMax, 0, &results)

    fmt.Printf("Точки в диапазоне [%v, %v]:\n", targetMin.Coordinates, targetMax.Coordinates)
    for _, point := range results {
        fmt.Println(point.Coordinates) // Печать точек, найденных в диапазоне
    }
}
