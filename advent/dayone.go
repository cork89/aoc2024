package advent

import (
	_ "embed"
	"fmt"
	"sort"
	"strings"
)

//go:embed dayone.txt
var dayoneinput string

func GetNums() ([]int, []int) {
	test := strings.Split(strings.Trim(dayoneinput, "\n"), "\n")
	var a []int = make([]int, 1000)
	var b []int = make([]int, 1000)
	for _, v := range test {
		temp := strings.Fields(v)
		first := getNum(temp[0])
		second := getNum(temp[1])
		a = append(a, first)
		b = append(b, second)
	}
	return a, b
}

func PartOne(a []int, b []int) int {

	sort.Ints(a)
	sort.Ints(b)

	diffsum := 0

	for i := range len(a) {
		diffsum += Abs(b[i] - a[i])
	}

	return diffsum
}

var bmap map[int]int = make(map[int]int)

func PartTwo(a []int, b []int) int {
	for _, bval := range b {
		_, ok := bmap[bval]
		if !ok {
			bmap[bval] = 1
		} else {
			bmap[bval] += 1
		}
	}
	simscore := 0
	for _, aval := range a {
		bcount, ok := bmap[aval]
		if ok {
			simscore += aval * bcount
		}
	}

	return simscore
}

func RunDayOne() {
	a, b := GetNums()
	fmt.Println("d1p1: ", PartOne(a, b))
	fmt.Println("d1p2: ", PartTwo(a, b))
}
