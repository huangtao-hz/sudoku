package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// ReadStdio 从标准输入读取数据
func ReadStdio(prompt string) (buf string) {
	fmt.Println(prompt)
	input := bufio.NewReader(os.Stdin)
	for {
		line, err := input.ReadString('\n')
		if err != nil {
			return
		}
		buf = buf + line
	}
}

func main() {
	lines := ReadStdio("请按行输入待解的数独，每行9个数字，共9行，空白以0代替，如001002003：")
	numbers := make([]int, 0)
	for _, ch := range lines {
		if val, err := strconv.Atoi(string(ch)); err == nil {
			numbers = append(numbers, val)
		}
	}
	sudoku := NewSudoku(numbers...)
	sudoku.Resolve()
	sudoku.Print()
}
