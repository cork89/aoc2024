package advent

import (
	_ "embed"
	"fmt"
	"strconv"
)

//go:embed daynine.txt
var daynineinput string

var allvals string = "0123456789."

const dot string = "."

func setupdaynine() ([]string, string) {
	densefmt := make([]string, 0)
	var maxIdx string
	for i := range daynineinput {
		temp := int(daynineinput[i] - allvals[0])
		if i%2 == 0 {
			idx := strconv.Itoa(i / 2)
			for j := 0; j < temp; j++ {
				densefmt = append(densefmt, idx)
			}
			maxIdx = idx
		} else {
			for j := 0; j < temp; j++ {
				densefmt = append(densefmt, dot)
			}
		}
	}
	return densefmt, maxIdx
}

func removeGaps(densefmt []string) []string {
	i := 0
	j := len(densefmt) - 1

	for i < j {
		if densefmt[i] != dot {
			i++
			continue
		} else if densefmt[j] == dot {
			j--
			continue
		} else {
			densefmt[i] = densefmt[j]
			densefmt[j] = dot
			i++
			j--
		}
	}
	return densefmt
}

type OpenGap struct {
	idx    int
	spaces int
}

func buildOpenGaps(densefmt []string, j int) []OpenGap {
	var tempI int = -1
	var i int = 0
	var openGaps = make([]OpenGap, 0)
	for i < j {
		if densefmt[i] == dot {
			tempI = i
			for tempI < j && densefmt[tempI] == dot {
				tempI++
			}
			openGaps = append(openGaps, OpenGap{idx: i, spaces: tempI - i})
			i = tempI
		}
		i++
	}
	return openGaps
}

func buildCurrIdx(densefmt []string, j int, currIdx string) OpenGap {
	var tempJ int = j
	for j > 0 {
		if densefmt[j] != currIdx {
			j--
		} else {
			tempJ = j
			for densefmt[tempJ] == currIdx {
				tempJ--
			}
			return OpenGap{idx: tempJ + 1, spaces: j - tempJ}
		}
	}
	panic("buildCurrIdx fail")
}

func removeGaps2(densefmt []string, maxIdx string) []string {
	j := len(densefmt) - 1
	maxIdxNum := getNum(maxIdx)
	for maxIdxNum > 1 {
		openGaps := buildOpenGaps(densefmt, j)
		currIdx := buildCurrIdx(densefmt, j, maxIdx)
		j = currIdx.idx
		for _, openGap := range openGaps {
			if openGap.spaces >= currIdx.spaces {
				for k := 0; k < currIdx.spaces; k++ {
					densefmt[openGap.idx+k] = densefmt[j+k]
					densefmt[j+k] = dot
				}
			}
		}
		maxIdxNum--
		maxIdx = strconv.Itoa(maxIdxNum)
	}
	return densefmt
}

func getChecksum(densefmt []string) int {
	total := 0

	for i, ch := range densefmt {
		if ch == dot {
			continue
		}
		total += i * getNum(ch)
	}

	return total
}

func getChecksum2(densefmt []string) int {
	total := 0
	for i, ch := range densefmt {
		if ch == dot {
			continue
		}
		total += i * getNum(ch)
	}

	return total
}

func RunDayNine() {
	densefmtp1, _ := setupdaynine()
	densefmtp1 = removeGaps(densefmtp1)
	fmt.Println("d9p1: ", getChecksum(densefmtp1))
	densefmtp2, maxIdx := setupdaynine()
	densefmtp2 = removeGaps2(densefmtp2, maxIdx)
	fmt.Println("d9p2: ", getChecksum2(densefmtp2))
}
