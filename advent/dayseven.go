package advent

import (
	_ "embed"
	"fmt"
	"slices"
	"strings"
)

//go:embed dayseven.txt
var dayseveninput string

type Problem struct {
	Result      int
	Operands    []int
	Resultlen   int
	Operandlens int
}

func setupProblems() []Problem {
	subproblems := strings.Split(dayseveninput, "\r\n")

	var problems = make([]Problem, 0, len(subproblems))

	for _, subproblem := range subproblems {
		parts := strings.Split(subproblem, ": ")

		operands := strings.Split(parts[1], " ")
		operandlens := 0
		var operandsslc = make([]int, 0, len(operands))
		for _, num := range operands {
			operandlens += len(num)
			operandsslc = append(operandsslc, getNum(num))
		}
		problem := Problem{Result: getNum(parts[0]),
			Operands:    operandsslc,
			Resultlen:   len(parts[0]),
			Operandlens: operandlens}
		problems = append(problems, problem)
	}
	return problems
}

var operators = "*+|"

func operate(problem Problem, depth int, currTotal int, rslts *[]int, operand byte) {
	if depth == len(problem.Operands) {
		*rslts = append(*rslts, currTotal)
		return
	}
	if operand == operators[0] {
		currTotal *= problem.Operands[depth]

	} else if operand == operators[1] {
		currTotal += problem.Operands[depth]
	}
	if currTotal <= problem.Result {
		operate(problem, depth+1, currTotal, rslts, operators[0])
		operate(problem, depth+1, currTotal, rslts, operators[1])
	}
}

func checkvalidresults(problems []Problem) int {
	total := 0

	for _, problem := range problems {
		if problem.Resultlen > problem.Operandlens {
			continue
		}

		var rslts = make([]int, 0)
		operate(problem, 1, problem.Operands[0], &rslts, operators[0])
		operate(problem, 1, problem.Operands[0], &rslts, operators[1])
		if slices.Contains(rslts, problem.Result) {
			total += problem.Result
		}
	}

	return total
}

func operate2(problem Problem, depth int, currTotal int, rslts *[]int, operator byte) {
	if operator == operators[0] {
		currTotal *= problem.Operands[depth]
	} else if operator == operators[1] {
		currTotal += problem.Operands[depth]
	} else if operator == operators[2] {
		if (depth + 2) < len(problem.Operands) {
			currTotal1 := getNum(fmt.Sprintf("%d%d", currTotal, problem.Operands[depth]+problem.Operands[depth+1]))
			currTotal2 := getNum(fmt.Sprintf("%d%d", currTotal, problem.Operands[depth]*problem.Operands[depth+1]))
			if currTotal1 <= problem.Result {
				operate2(problem, depth+2, currTotal1, rslts, operators[0])
				operate2(problem, depth+2, currTotal1, rslts, operators[1])
				operate2(problem, depth+2, currTotal1, rslts, operators[2])
			} else if currTotal2 <= problem.Result {
				operate2(problem, depth+2, currTotal2, rslts, operators[0])
				operate2(problem, depth+2, currTotal2, rslts, operators[1])
				operate2(problem, depth+2, currTotal2, rslts, operators[2])
			}
		}
		currTotal = getNum(fmt.Sprintf("%d%d", currTotal, problem.Operands[depth]))
	}
	if (depth + 1) >= len(problem.Operands) {
		*rslts = append(*rslts, currTotal)
		return
	}
	if currTotal <= problem.Result {
		operate2(problem, depth+1, currTotal, rslts, operators[0])
		operate2(problem, depth+1, currTotal, rslts, operators[1])
		operate2(problem, depth+1, currTotal, rslts, operators[2])
	}
}

func checkvalidresults2(problems []Problem) int {
	total := 0

	for _, problem := range problems {
		if problem.Resultlen > problem.Operandlens {
			continue
		}

		var rslts = make([]int, 0)
		operate2(problem, 1, problem.Operands[0], &rslts, operators[0])
		operate2(problem, 1, problem.Operands[0], &rslts, operators[1])
		operate2(problem, 1, problem.Operands[0], &rslts, operators[2])
		if slices.Contains(rslts, problem.Result) {
			total += problem.Result
		}
	}

	return total
}

func RunDaySeven() {
	problems := setupProblems()
	fmt.Println("d7p1: ", checkvalidresults(problems))
	fmt.Println("d7p2: ", checkvalidresults2(problems))
}
