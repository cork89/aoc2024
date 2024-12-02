package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

//go:embed daytwo.txt
var input string
var scratch [8]int = [8]int{0, 0, 0, 0, 0, 0, 0, 0}
var debug [8]int = [8]int{0, 0, 0, 0, 0, 0, 0, 0}

func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func getNum(ch string) int {
	num, err := strconv.Atoi(ch)
	if err != nil {
		panic(err)
	}
	return num
}

func isSafe(lead int, follow int, ascending bool) (bool, int) {
	var diff int
	if ascending {
		diff = lead - follow
	} else {
		diff = follow - lead
	}
	if diff > 3 || diff < 1 {
		return false, diff
	}
	return true, diff
}

func GetReports() int {
	inputrows := strings.Split(strings.Trim(input, "\n"), "\n")
	var safeReports int = 0
	for _, v := range inputrows {
		row := strings.Fields(v)
		safe, _ := checkSafety(row)
		if safe {
			safeReports += 1
		}
	}
	return safeReports
}

func checkAscending(row []string) bool {
	inc := 0
	dec := 0
	for i := 1; i < len(row); i++ {
		if row[i] > row[i-1] {
			inc += 1
		} else {
			dec += 1
		}
	}
	return inc > dec
}

func checkSafety(row []string) (bool, int) {
	ascending := checkAscending(row)
	scratch[0] = getNum(row[0])
	var unsafeComparison int = -1
	var safe bool = true

	for i := 1; i < len(row); i++ {
		scratch[i] = getNum(row[i])
		safe, _ = isSafe(scratch[i], scratch[i-1], ascending)
		if !safe {
			unsafeComparison = i - 1
			break
		}
	}
	return safe, unsafeComparison
}

func GetReportsDamped() int {
	inputrows := strings.Split(strings.Trim(input, "\n"), "\n")
	var safeReports int = 0
	for _, v := range inputrows {
		row := strings.Fields(v)

		safe, unsafeComparison := checkSafety(row)

		if !safe {
			leadRow := slices.Clone(row)
			leadRow = append(leadRow[:unsafeComparison], leadRow[unsafeComparison+1:]...)
			safe, _ = checkSafety(leadRow)
		}

		if !safe {
			followRow := slices.Clone(row)
			followRow = append(followRow[:unsafeComparison+1], followRow[unsafeComparison+2:]...)
			safe, _ = checkSafety(followRow)
		}

		if safe {
			safeReports += 1
		}

	}
	return safeReports
}

func main() {
	fmt.Println(GetReports())
	fmt.Println(GetReportsDamped())
}
