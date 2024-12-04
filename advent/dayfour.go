package advent

import (
	_ "embed"
	"fmt"
	"strings"
	"sync"
	"sync/atomic"
)

//go:embed dayfour.test
var dayfourinput string

func checkHorizontalDir(lines []string, row int, col int, forward bool) bool {
	if forward {
		if len(lines[row])-col < 4 {
			return false
		}
		if lines[row][col:col+4] == "XMAS" {
			return true
		}
	} else {
		if col < 3 {
			return false
		}
		if (lines[row][col-3 : col+1]) == "SAMX" {
			return true
		}
	}
	return false
}

func checkHorizontal(lines []string, row int, col int) int {
	horizontalTotal := 0
	if checkHorizontalDir(lines, row, col, true) {
		horizontalTotal += 1
	}
	if checkHorizontalDir(lines, row, col, false) {
		horizontalTotal += 1
	}
	return horizontalTotal
}

const xmas string = "XMAS"

func checkVerticalDir(lines []string, row int, col int, down bool) bool {
	if down {
		if len(lines)-row < 4 {
			return false
		}
		if lines[row][col] == xmas[0] {
			for i := 1; i < 4; i++ {
				if lines[row+i][col] != xmas[i] {
					return false
				}
			}
			return true
		}
	} else {
		if row < 3 {
			return false
		}
		if lines[row][col] == xmas[0] {
			for i := 1; i < 4; i++ {
				if lines[row-i][col] != xmas[i] {
					return false
				}
			}
			return true
		}
	}
	return false
}

func checkVertical(lines []string, row int, col int) int {
	verticalTotal := 0
	if checkVerticalDir(lines, row, col, true) {
		verticalTotal += 1
	}
	if checkVerticalDir(lines, row, col, false) {
		verticalTotal += 1
	}
	return verticalTotal
}

func checkDiagonalDir(lines []string, row int, col int, forward bool, down bool) bool {
	if down && forward {
		if len(lines)-row < 4 || len(lines[row])-col < 4 {
			return false
		}
		if lines[row][col] == xmas[0] {
			for i := 1; i < 4; i++ {
				if lines[row+i][col+i] != xmas[i] {
					return false
				}
			}
			return true
		}
	} else if down && !forward {
		if len(lines)-row < 4 || col < 3 {
			return false
		}
		if lines[row][col] == xmas[0] {
			for i := 1; i < 4; i++ {
				if lines[row+i][col-i] != xmas[i] {
					return false
				}
			}
			return true
		}
	} else if !down && forward {
		if row < 3 || len(lines[row])-col < 4 {
			return false
		}
		if lines[row][col] == xmas[0] {
			for i := 1; i < 4; i++ {
				if lines[row-i][col+i] != xmas[i] {
					return false
				}
			}
			return true
		}
	} else {
		if row < 3 || col < 3 {
			return false
		}
		if lines[row][col] == xmas[0] {
			for i := 1; i < 4; i++ {
				if lines[row-i][col-i] != xmas[i] {
					return false
				}
			}
			return true
		}
	}
	return false
}

func checkDiagonal(lines []string, row int, col int) int {
	diagonalTotal := 0
	if checkDiagonalDir(lines, row, col, true, true) {
		diagonalTotal += 1
	}
	if checkDiagonalDir(lines, row, col, true, false) {
		diagonalTotal += 1
	}
	if checkDiagonalDir(lines, row, col, false, true) {
		diagonalTotal += 1
	}
	if checkDiagonalDir(lines, row, col, false, false) {
		diagonalTotal += 1
	}
	return diagonalTotal
}

func getXmas(lines []string) int {
	total := 0
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			if lines[i][j] != xmas[0] {
				continue
			}
			total += checkHorizontal(lines, i, j)
			total += checkVertical(lines, i, j)
			total += checkDiagonal(lines, i, j)
		}
	}
	return total
}

var m byte = xmas[1]
var s byte = xmas[3]

func getXmas2(lines []string) int {
	total := 0
	for i := 1; i < len(lines)-1; i++ {
		for j := 1; j < len(lines[i])-2; j++ {
			if lines[i][j] == xmas[2] {
				subcount := 0
				if (lines[i-1][j-1] == m && lines[i+1][j+1] == s) || (lines[i-1][j-1] == s && lines[i+1][j+1] == m) {
					subcount += 1
				}
				if (lines[i-1][j+1] == m && lines[i+1][j-1] == s) || (lines[i-1][j+1] == s && lines[i+1][j-1] == m) {
					subcount += 1
				}
				if subcount == 2 {
					total += 1
				}
			}
		}
	}
	return total
}

type counter struct {
	total int32
}

func (c *counter) run(wg *sync.WaitGroup, lines []string, i int) {
	defer wg.Done()
	var subCount int = 0
	for j := 0; j < len(lines[i]); j++ {
		if lines[i][j] != xmas[0] {
			continue
		}
		subCount += checkHorizontal(lines, i, j)
		subCount += checkVertical(lines, i, j)
		subCount += checkDiagonal(lines, i, j)
	}
	_ = atomic.AddInt32(&c.total, int32(subCount))
}

func getXmasConcurrent(lines []string) int32 {
	cnt := counter{}
	wg := sync.WaitGroup{}
	wg.Add(len(lines))

	for i := 0; i < len(lines); i++ {
		cnt.run(&wg, lines, i)
	}
	wg.Wait()

	return cnt.total
}

func RunDayFour() {
	lines := strings.Split(dayfourinput, "\n")
	xmas1 := getXmas(lines)
	getXmasConcurrent(lines)

	fmt.Println("d4p1: ", xmas1)
	fmt.Println("d4p2: ", getXmas2(lines))
}
