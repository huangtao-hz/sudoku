package main

import (
	mapset "github.com/deckarep/golang-set/v2"
)

// Item 数独的单元格
type Item struct {
	Owner                      *Sudoku
	Pos, Row, Col, Grid, Value int
	Available                  mapset.Set[int]
}

// NewItem 单元格构造函数
func NewItem(Owner *Sudoku, Pos int, Value int) *Item {
	var Available mapset.Set[int]
	Row := Pos / 9
	Col := Pos % 9
	Grid := Row/3*3 + Col
	Available = mapset.NewSet(1, 2, 3, 4, 5, 6, 7, 8, 9)
	if Value > 0 {
		Available.Clear()
	}
	return &Item{Owner: Owner, Pos: Pos, Row: Row, Col: Col, Grid: Grid, Value: Value, Available: Available}
}

// Sudoku 数独类
type Sudoku struct {
	Items   []*Item
	Steps   []string
	Success bool
}

// NewSudoku 数独类构造函数
func NewSudoku(values ...int) *Sudoku {
	Items := make([]*Item, 81)
	Steps := make([]string, 0)
	sudoku := &Sudoku{Items: Items, Steps: Steps}
	for i := range 81 {
		Items[i] = NewItem(sudoku, i, 0)
	}
	for i, value := range values {
		if value > 0 && value <= 9 && i < 81 {
			sudoku.SetValue(i, value, "")
		}
	}
	return sudoku
}

// GetItem 获取值
func (s *Sudoku) GetItem(pos int) *Item {
	return s.Items[pos]
}

// SetValue 设置值
func (s *Sudoku) SetValue(pos int, value int, msg string) {
	item := s.GetItem(pos)
	item.Value = value
	item.Available.Clear()
	for _, i := range s.Items {
		if i.Value > 0 && (i.Col == item.Col || i.Row == item.Row || i.Grid == item.Grid) {
			i.Available.Remove(item.Value)
		}
	}
	if msg != "" {
		s.Steps = append(s.Steps, msg)
	}
}
