package main

import (
	"fmt"

	mapset "github.com/deckarep/golang-set/v2"
)

// Sudoku 数独类
type Sudoku struct {
	Items []*Item  // 单元格
	Steps []string // 解题步骤
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
			sudoku.SetValue(i, value, 9)
		}
	}
	return sudoku
}

// Clone 复制数独类
func (s *Sudoku) Clone() *Sudoku {
	values := make([]int, 81)
	for i, item := range s.Items {
		values[i] = item.Value
	}
	new := NewSudoku(values...)
	new.Steps = make([]string, len(s.Steps))
	copy(new.Steps, s.Steps)
	return new
}

// SetValue 设置值
func (s *Sudoku) SetValue(pos int, value int, method int) {
	var Methods = []string{"行排除", "列排除", "方格排除", "可选值唯一", "随机"}
	item := s.Items[pos]
	item.Value = value
	item.Available = nil
	for _, i := range s.Items {
		if i.Value == 0 && (i.Col == item.Col || i.Row == item.Row || i.Grid == item.Grid) {
			i.Available.Remove(value)
		}
	}
	if method < 5 {
		s.Steps = append(s.Steps, fmt.Sprintf("第%d行，第%d列，设为：%d，方法：%s", item.Row+1, item.Col+1, value, Methods[method]))
	}
}

// Remove 删除切片中的指定元素
func Remove[T comparable](s []T, i T) []T {
	l := len(s)
	for j := range s {
		if s[j] == i {
			s[j], s[l-1] = s[l-1], s[j]
			return s[:l-1]
		}
	}
	return s
}

// Print 打印结果
func (s *Sudoku) Print() {
	fmt.Println("行/列 1 2 3 4 5 6 7 8 9")
	for i, item := range s.Items {
		if i%9 == 0 {
			fmt.Printf("  %d   ", i/9+1)
		}
		if item.Value > 0 {
			fmt.Print(item.Value)
		} else {
			fmt.Print(" ")
		}
		if i%9 < 8 {
			fmt.Print(" ")
		} else {
			fmt.Println()
		}
	}
	if len(s.Steps) > 0 {
		fmt.Println("解题步骤：")
	}

	for _, line := range s.Steps {
		fmt.Println(line)
	}
}

// getBlankItems 获取未填充的单元格
func (s *Sudoku) getBlankItems() []*Item {
	items := make([]*Item, 0)
	for _, i := range s.Items {
		if i.Value == 0 {
			items = append(items, i)
		}
	}
	return items
}

// Resolve 解决数独
func (s *Sudoku) Resolve() bool {
	items := s.getBlankItems()
	var found bool
	for len(items) > 0 {
		found = false
		for _, item := range items {
			avl := item.Available.ToSlice()
			if len(avl) == 1 {
				s.SetValue(item.Pos, avl[0], 3)
				items = Remove(items, item)
				found = true
				break
			}
			rowa, cola, grida := item.Available.Clone(), item.Available.Clone(), item.Available.Clone()
			for _, it := range items {
				if it.Pos == item.Pos {
					continue
				}
				if it.Row == item.Row {
					rowa = rowa.Difference(it.Available)
				}
				if it.Col == item.Col {
					cola = cola.Difference(it.Available)
				}
				if it.Grid == item.Grid {
					grida = grida.Difference(it.Available)
				}
			}
			for method, b := range []mapset.Set[int]{rowa, cola, grida} {
				avl := b.ToSlice()
				if len(avl) == 1 {
					s.SetValue(item.Pos, avl[0], method)
					items = Remove(items, item)
					found = true
					break
				}
			}
			if found {
				break
			}
		}
		if !found {
			item := items[0]
			for _, value := range item.Available.ToSlice() {
				new := s.Clone()
				new.SetValue(item.Pos, value, 4)
				if success := new.Resolve(); success {
					for i := range 81 {
						s.Items[i].Value = new.Items[i].Value
					}
					s.Steps = new.Steps
					return true
				}
			}
			return false
		}
	}
	return true
}
