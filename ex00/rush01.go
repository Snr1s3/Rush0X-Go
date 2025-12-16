package main

import (
	"fmt"
	"os"
	"strconv"
)

var Views [][]int
var Map [][]int
var Rows int = 4
var Cols int = 4

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run rush01.go <cluestring>")
		return
	}
	clues := os.Args[1]
	if len(clues) != 16 {
		fmt.Println("Clue string must be 16 digits")
		return
	}
	input := make([]string, 16)
	for i := 0; i < 16; i++ {
		input[i] = string(clues[i])
	}
	Allocate(input)
	PrintMap()
	if !Solve(0, 0) {
		fmt.Println("No solution found for the given clues.")
	}

	fmt.Println()
	PrintMap()
}
func Allocate(input []string) {
	fmt.Println("Allocating")
	Map = make([][]int, Rows)
	for i := range Map {
		Map[i] = make([]int, Cols)
		for j := range Map[i] {
			Map[i][j] = 0
		}
	}
	Views = make([][]int, 4)
	for i := range Views {
		Views[i] = make([]int, Cols)
	}
	count := 0
	for i := 0; i < Rows; i++ {
		for j := 0; j < Cols; j++ {
			val, err := strconv.Atoi(input[count])
			if err != nil {
				fmt.Printf("Invalid integer: %s\n", input[count])
				return
			}
			Views[i][j] = val
			count++
		}
	}
}
func PrintMapLine(row int) {
	for c := 0; c < Cols; c++ {
		fmt.Print(" ")
		fmt.Print(Map[row][c])
	}
}
func Solve(row int, col int) bool {
	var nextR, nextC int

	if row == Rows {
		return true
	}
	if col == Cols-1 {
		nextR = row + 1
		nextC = 0
	} else {
		nextR = row
		nextC = col + 1
	}
	for num := 1; num <= 4; num++ {
		if ValidNum(row, col, num) {
			Map[row][col] = num
			if CheckVisibility(row, col) && Solve(nextR, nextC) {
				return true
			}
			Map[row][col] = 0
		}
	}
	return false
}
func ValidNum(row int, col int, num int) bool {
	for c := 0; c < Cols; c++ {
		if Map[row][c] == num {
			return false
		}
	}
	for r := 0; r < Rows; r++ {
		if Map[r][col] == num {
			return false
		}
	}
	return true
}
func CheckVisibility(row int, col int) bool {
	if col == Cols-1 && !RowVisibility(row) {
		return false
	}

	if row == Rows-1 && !ColVisibility(col) {
		return false
	}

	return true
}
func RowVisibility(row int) bool {
	var MaxLeft, MaxRight, VisiLeft, VisiRight int = 0, 0, 0, 0
	for c := 0; c < Cols; c++ {
		if MaxLeft < Map[row][c] {
			MaxLeft = Map[row][c]
			VisiLeft++
		}
	}
	for c := Cols - 1; c >= 0; c-- {
		if MaxRight < Map[row][c] {
			MaxRight = Map[row][c]
			VisiRight++
		}
	}
	return (Views[2][row] == 0 || Views[2][row] == VisiLeft) && (Views[3][row] == 0 || Views[3][row] == VisiRight)
}

func ColVisibility(col int) bool {
	var MaxTop, MaxBot, VisiTop, VisiBot int = 0, 0, 0, 0
	for c := 0; c < Cols; c++ {
		if MaxTop < Map[c][col] {
			MaxTop = Map[c][col]
			VisiTop++
		}
	}
	for c := Cols - 1; c >= 0; c-- {
		if MaxBot < Map[c][col] {
			MaxBot = Map[c][col]
			VisiBot++
		}
	}
	return (Views[0][col] == 0 || Views[0][col] == VisiTop) && (Views[1][col] == 0 || Views[1][col] == VisiBot)
}
func PrintMap() {
	fmt.Print("  ")
	for c := 0; c < Cols; c++ {
		fmt.Print(Views[0][c], " ")
	}
	fmt.Println()
	for r := 0; r < Rows; r++ {
		fmt.Print(Views[2][r], " ")
		for c := 0; c < Cols; c++ {
			fmt.Print(Map[r][c], " ")
		}
		fmt.Print(Views[3][r])
		fmt.Println()
	}
	fmt.Print("  ")
	for c := 0; c < Cols; c++ {
		fmt.Print(Views[1][c], " ")
	}
	fmt.Println()
}
