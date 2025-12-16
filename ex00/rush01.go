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
	if !Allocate(input) {
		return
	}
	fmt.Println()
	fmt.Println("Map to solve:")
	PrintMap()
	if !Solve(0, 0) {
		fmt.Println("No solution found for the given clues.")
	}
	fmt.Println()
	fmt.Println("Map solved:")
	PrintMap()
}
func Allocate(input []string) bool {
	fmt.Println("Allocating Arrays in Memory")
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
			if err != nil || val > 4 || val < 1 {
				fmt.Printf("Invalid integer: %s\n", input[count])
				return false
			}
			Views[i][j] = val
			count++
		}
	}
	return true
}

func PrintMap() {
	PrintTopBotViews(0)
	PrintTopBotLine()
	PrintMiddle()
	PrintTopBotViews(1)
}
func PrintMiddle() {
	for r := 0; r < Rows; r++ {
		fmt.Printf(" %d |", Views[2][r])
		for c := 0; c < Cols; c++ {
			if Map[r][c] == 0 {
				fmt.Print("   |")
			} else {
				fmt.Printf(" %d |", Map[r][c])
			}
		}
		fmt.Printf(" %d\n", Views[3][r])
		PrintTopBotLine()
	}
}
func PrintTopBotLine() {
	fmt.Print("   +")
	for c := 0; c < Cols; c++ {
		fmt.Print("---+")
	}
	fmt.Println()
}
func PrintTopBotViews(view int) {
	fmt.Print("     ")
	for c := 0; c < Cols; c++ {
		fmt.Printf("%d   ", Views[view][c])
	}
	fmt.Println()
}
