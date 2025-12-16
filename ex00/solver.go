package main

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
