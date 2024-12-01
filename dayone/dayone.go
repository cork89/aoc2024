package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed dayone.txt
var input string

func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func GetNums() ([]int, []int) {
	test := strings.Split(strings.Trim(input, "\n"), "\n")
	var a []int = make([]int, 1000)
	var b []int = make([]int, 1000)
	for _, v := range test {
		temp := strings.Fields(v)
		first, err := strconv.Atoi(temp[0])
		if err != nil {
			panic(err)
		}
		second, err := strconv.Atoi(temp[1])
		if err != nil {
			panic(err)
		}
		a = append(a, first)
		b = append(b, second)
	}
	return a, b
}

func PartOne(a []int, b []int) int {

	sort.Sort(sort.IntSlice(a))
	sort.Sort(sort.IntSlice(b))

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

func main() {
	a, b := GetNums()
	fmt.Println(PartOne(a, b))
	fmt.Println(PartTwo(a, b))
}
