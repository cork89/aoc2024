package advent

import (
	_ "embed"
	"fmt"
	"slices"
	"sort"
	"strings"
)

//go:embed dayfive.txt
var dayfiveinput string

var rules map[string][]string = make(map[string][]string, 0)

func validate(subscore []string, depth int) bool {
	for i := 0; i < len(subscore)-1; i++ {
		for j := i + 1; j < len(subscore); j++ {
			rule := rules[subscore[j]]
			if slices.Contains(rule, subscore[i]) {
				temp := subscore[i]
				subscore[i] = subscore[j]
				subscore[j] = temp
				return validate(subscore, depth+1)
			}
		}
	}
	return true
}

func buildScores() string {
	dayfiveinput = strings.ReplaceAll(dayfiveinput, "\r\n", "\n")
	parts := strings.Split(dayfiveinput, "\n\n")

	allrules := strings.Split(parts[0], "\n")
	for _, v := range allrules {
		ruleline := strings.Split(v, "|")

		first := ruleline[0]
		second := ruleline[1]

		_, ok := rules[first]
		if !ok {
			rules[first] = make([]string, 0)
		}
		rules[first] = append(rules[first], second)
	}
	return parts[1]
}

func buildSubscores(update []string) map[string]int {
	subscores := make(map[string]int, 0)
	for _, first := range update {
		for _, second := range rules[first] {
			_, firstok := subscores[first]
			i, secondok := subscores[second]
			if !firstok {
				subscores[first] = 1
			} else {
				if i == 0 {
					subscores[first] += 1
				} else {
					subscores[first] += i
				}
			}
			if !secondok {
				subscores[second] = 0
			}
		}
	}
	return subscores
}

func correctlyOrderedPages(allupdates []string) (int, []string) {
	total := 0
	incorrectUpdates := make([]string, 0)
	for _, line := range allupdates {
		update := strings.Split(line, ",")
		isvalid := true
		subscores := buildSubscores(update)
		for i := len(update) - 1; i > 0; i-- {
			curr := update[i]
			prev := update[i-1]

			if subscores[prev] < subscores[curr] {
				isvalid = false
				break
			}
		}

		if isvalid {
			total += getNum(update[len(update)/2])
		} else {
			incorrectUpdates = append(incorrectUpdates, line)
		}
	}

	return total, incorrectUpdates
}

func incorrectlyOrderedPages(updates []string) int {
	total := 0

	for _, line := range updates {
		update := strings.Split(line, ",")
		subscores := buildSubscores(update)
		sort.Slice(update, func(i, j int) bool {
			a := subscores[update[i]]
			b := subscores[update[j]]
			return a > b
		})
		validate(update, 0)

		total += getNum(update[len(update)/2])
	}

	return total
}

func RunDayFive() {
	updates := buildScores()
	allupdates := strings.Split(updates, "\n")
	a1, incorrectupdates := correctlyOrderedPages(allupdates)
	fmt.Println("d5p1: ", a1)
	fmt.Println("d5p2: ", incorrectlyOrderedPages(incorrectupdates))
}
