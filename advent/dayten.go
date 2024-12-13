package advent

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed dayten.txt
var dayteninput string

func setupdayten() [][]int {
	topomap := make([][]int, 0)
	rows := strings.Split(dayteninput, "\r\n")
	for _, row := range rows {
		cols := strings.Split(row, "")
		topocol := make([]int, 0)
		for _, col := range cols {
			topocol = append(topocol, getNum(col))
		}
		topomap = append(topomap, topocol)
	}
	return topomap
}

func getTrailheads(topomap [][]int) []Coord {
	coords := make([]Coord, 0)
	for i, row := range topomap {
		for j, col := range row {
			if col == 0 {
				coords = append(coords, Coord{x: j, y: i})
			}
		}
	}
	return coords
}

func getPath(topomap *[][]int, pos Position, maxelevations *[]Position) {
	var elevation int = (*topomap)[pos.coord.y][pos.coord.x]

	if elevation == 9 {
		*maxelevations = append(*maxelevations, pos)
	}

	//up
	upos := pos.Up()
	if upos.Inbounds() && (*topomap)[upos.coord.y][upos.coord.x] == elevation+1 {
		getPath(topomap, upos, maxelevations)
	}
	//left
	lpos := pos.Left()
	if lpos.Inbounds() && (*topomap)[lpos.coord.y][lpos.coord.x] == elevation+1 {
		getPath(topomap, lpos, maxelevations)
	}
	//right
	rpos := pos.Right()
	if rpos.Inbounds() && (*topomap)[rpos.coord.y][rpos.coord.x] == elevation+1 {
		getPath(topomap, rpos, maxelevations)
	}
	//down
	dpos := pos.Down()
	if dpos.Inbounds() && (*topomap)[dpos.coord.y][dpos.coord.x] == elevation+1 {
		getPath(topomap, dpos, maxelevations)
	}
}

func getUniqueElevations(maxelevations []Position) int {
	elevationmap := make(map[Coord]bool)
	for _, elevation := range maxelevations {
		elevationmap[elevation.coord] = true
	}

	return len(elevationmap)
}

func getTrailscores(topomap [][]int, trailheads []Coord) int {
	total := 0

	xmax := len(topomap[0])
	ymax := len(topomap)

	for _, trailhead := range trailheads {
		maxelevations := make([]Position, 0)
		pos := Position{coord: Coord{x: trailhead.x, y: trailhead.y}, xmax: xmax, ymax: ymax}
		getPath(&topomap, pos, &maxelevations)
		total += getUniqueElevations(maxelevations)
	}

	return total
}

func getTrailscores2(topomap [][]int, trailheads []Coord) int {
	total := 0

	xmax := len(topomap[0])
	ymax := len(topomap)

	for _, trailhead := range trailheads {
		maxelevations := make([]Position, 0)
		pos := Position{coord: Coord{x: trailhead.x, y: trailhead.y}, xmax: xmax, ymax: ymax}
		getPath(&topomap, pos, &maxelevations)
		total += len(maxelevations)
	}

	return total
}

// 238 too low
func RunDayTen() {
	topomap := setupdayten()
	trailheads := getTrailheads(topomap)
	trailscore := getTrailscores(topomap, trailheads)

	fmt.Println("d10p1: ", trailscore)
	fmt.Println("d10p2: ", getTrailscores2(topomap, trailheads))
}
