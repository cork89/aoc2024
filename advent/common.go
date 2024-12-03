package advent

import "strconv"

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
