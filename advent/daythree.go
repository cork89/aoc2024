package advent

import (
	_ "embed"
	"fmt"
	"regexp"
	"strings"
)

//go:embed daythree.txt
var daythreeinput string

var validMuls = regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)

func GetMuls(input string) int {
	matches := validMuls.FindAllString(input, -1)
	total := 0
	for _, v := range matches {
		parts := strings.Split(v, ",")
		op1 := getNum(parts[0][4:])
		op2 := getNum(strings.TrimSuffix(parts[1], ")"))
		total += op1 * op2
	}
	return total
}

func GetPreciseMuls(input string) int {
	doStatements := strings.Split(input, "do()")
	total := 0
	for i := 0; i < len(doStatements); i++ {
		dontParts := strings.Split(doStatements[i], "don't()")
		if len(dontParts) == 0 {
			total += GetMuls(doStatements[i])
		} else {
			total += GetMuls(dontParts[0])
		}
	}

	return total
}

func RunDayThree() {
	fmt.Println("d3p1: ", GetMuls(daythreeinput))
	fmt.Println("d3p2: ", GetPreciseMuls(daythreeinput))
}
