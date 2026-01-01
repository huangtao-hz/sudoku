package main

import (
	"fmt"

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
	Grid := Row/3*3 + Col/3
	if Value == 0 {
		Available = mapset.NewSet(1, 2, 3, 4, 5, 6, 7, 8, 9)
	}
	return &Item{Owner: Owner, Pos: Pos, Row: Row, Col: Col, Grid: Grid, Value: Value, Available: Available}
}

// Print 打印单元格
func (i *Item) Print() {
	fmt.Printf("第%d行，第%d列，值：%d，可能值：%v", i.Row, i.Col, i.Value, i.Available)
}
