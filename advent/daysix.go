package advent

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed daysix.txt
var daysixinput string

type Coord struct {
	x int
	y int
}

type Position struct {
	coord Coord
	xmax  int
	ymax  int
	dir   Direction
}

func (p Position) Copy() Position {
	return Position{
		coord: Coord{x: p.coord.x, y: p.coord.y},
		xmax:  p.xmax,
		ymax:  p.ymax,
		dir:   p.dir,
	}
}

type Direction string

const (
	up    Direction = "up"
	left  Direction = "left"
	right Direction = "right"
	down  Direction = "down"
)

var debug bool = false

var uniquepositions map[Coord]int

var possibleloops []Coord = make([]Coord, 0)

var guard string = "^#"

func inbounds(pos Position) bool {
	if pos.dir == up && pos.coord.y <= 0 {
		return false
	} else if pos.dir == right && pos.coord.x >= pos.xmax-1 {
		return false
	} else if pos.dir == left && pos.coord.x <= 0 {
		return false
	} else if pos.dir == down && pos.coord.y >= pos.ymax-1 {
		return false
	}
	return true
}

func calcpos(pos Position, patrol []string) Position {
	// possibleloops = append(possibleloops, pos.coord)
	if pos.dir == up {
		if patrol[pos.coord.y-1][pos.coord.x] != guard[1] {
			pos.coord.y = pos.coord.y - 1
		} else {
			pos.dir = right
		}
	} else if pos.dir == right {
		if patrol[pos.coord.y][pos.coord.x+1] != guard[1] {
			pos.coord.x = pos.coord.x + 1
		} else {
			pos.dir = down
		}
	} else if pos.dir == down {
		if patrol[pos.coord.y+1][pos.coord.x] != guard[1] {
			pos.coord.y = pos.coord.y + 1
		} else {
			pos.dir = left
		}
	} else if pos.dir == left {
		if patrol[pos.coord.y][pos.coord.x-1] != guard[1] {
			pos.coord.x = pos.coord.x - 1
		} else {
			pos.dir = up
		}
	}
	return pos
}

func daysixsetup() (Position, []string, []Coord) {
	patrol := strings.Split(daysixinput, "\n")
	pos := Position{xmax: len(patrol[0]) - 1, ymax: len(patrol), dir: up}
	barriers := make([]Coord, 0)
	uniquepositions = make(map[Coord]int, pos.xmax*pos.ymax)
	for i := 0; i < pos.ymax; i++ {
		for j := 0; j < pos.xmax; j++ {
			coord := Coord{x: j, y: i}
			if patrol[i][j] == guard[0] {
				pos.coord = coord
			}
			uniquepositions[coord] = 0
			if patrol[i][j] == guard[1] {
				barriers = append(barriers, coord)
			}
		}
	}
	uniquepositions[pos.coord] = 1
	return pos, patrol, barriers
}

func printpatrol(patrol []string, pos Position, prevpos Position) {
	if debug {
		for i, row := range patrol {
			for j, col := range row {
				if i == prevpos.coord.y && j == prevpos.coord.x {
					fmt.Print("%")
				} else if i == pos.coord.y && j == pos.coord.x {
					fmt.Print("&")
				} else {
					fmt.Print(string(col))
				}
			}
			fmt.Print("\n")
		}
	}
}

func getUniquePositions(pos Position, patrol []string, depth int) (int, bool) {

	total := 1
	var prevpos Position
	var looped bool = false
	for {
		prevpos = pos
		pos = calcpos(pos, patrol)

		inbounds := inbounds(pos)

		if uniquepositions[pos.coord] == 0 {
			uniquepositions[pos.coord] = 1
			total += 1
		} else {
			uniquepositions[pos.coord] += 1

			if uniquepositions[pos.coord] > depth {
				return total, true
			}
		}

		if !inbounds {
			break
		}
	}
	printpatrol(patrol, pos, prevpos)

	return total, looped
}

func buildpatrol(patrol []string, coord Coord) []string {
	var modifiedpatrol []string = make([]string, 0)

	for i, row := range patrol {
		if coord.y == i {
			newrow := fmt.Sprintf("%s%s%s", row[:coord.x], "#", row[coord.x+1:])
			modifiedpatrol = append(modifiedpatrol, newrow)
		} else {
			modifiedpatrol = append(modifiedpatrol, row)
		}
	}
	return modifiedpatrol
}

func getLoops(pos Position, patrol []string, barriers []Coord) int {
	total := 0

	var ycoords map[int][]Coord = make(map[int][]Coord, 0)
	var xcoords map[int][]Coord = make(map[int][]Coord, 0)

	for _, coord := range barriers {
		yslc, ok := ycoords[coord.y]

		if !ok {
			yslc = make([]Coord, 0)
			yslc = append(yslc, coord)
			ycoords[coord.y] = yslc
		}

		xslc, ok := xcoords[coord.x]

		if !ok {
			xslc = make([]Coord, 0)
			xslc = append(xslc, coord)
			xcoords[coord.x] = xslc
		}
	}

	for i := 0; i < pos.ymax; i++ {
		for j := 0; j < pos.xmax; j++ {
			coord := Coord{x: j, y: i}
			if coord == pos.coord || patrol[i][j] == guard[0] {
				continue
			}
			_, yok := ycoords[i]
			_, xok := xcoords[j]
			if yok {
				_, xminusok := xcoords[j-1]
				_, xplusok := xcoords[j+1]
				if xminusok || xplusok {
					possibleloops = append(possibleloops, coord)
				}
			} else if xok {
				_, yminusok := ycoords[j-1]
				_, yplusok := ycoords[j+1]
				if yminusok || yplusok {
					possibleloops = append(possibleloops, coord)
				}
			}
		}
	}
	depth := 500
	for _, coord := range possibleloops {
		for k := range uniquepositions {
			uniquepositions[k] = 0
		}
		modifiedpatrol := buildpatrol(patrol, coord)
		_, looped := getUniquePositions(pos, modifiedpatrol, depth)
		if looped {
			total += 1
		}
	}

	return total
}

func RunDaySix() {
	pos, patrol, barriers := daysixsetup()
	uniquepositions, _ := getUniquePositions(pos, patrol, 500)
	fmt.Println("d6p1: ", uniquepositions)
	fmt.Println("d6p2: ", getLoops(pos, patrol, barriers))
}
