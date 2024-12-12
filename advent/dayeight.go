package advent

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed dayeight.test
var dayeightinput string

// var antennas = "abcdefghiklmnopqrstuvwxyzABCDEFGHIKLMNOPQRSTUVWXYZ0123456789"
var noantenna = "."

func setupdayeight() map[byte][]Coord {
	rows := strings.Split(dayeightinput, "\r\n")
	antennamap := make(map[byte][]Coord, 0)
	for i := range rows {
		for j := range rows[i] {
			if rows[i][j] == noantenna[0] {
				continue
			}
			_, ok := antennamap[rows[i][j]]
			if !ok {
				antennamap[rows[i][j]] = make([]Coord, 0)
			}
			antennamap[rows[i][j]] = append(antennamap[rows[i][j]], Coord{x: j, y: i})
		}
	}
	return antennamap
}

func getAntinodes(antennas map[byte][]Coord) int {
	total := 0

	for k := range antennas {
		coords := antennas[k]
		if len(coords) == 1 {
			continue
		} else if len(coords) == 2 {
			var xdist int
			var ydist int
			if coords[0].x < coords[1].x {
				xdist = coords[1].x - coords[0].x
			} else {
				xdist = coords[0].x - coords[1].x
			}
			fmt.Println(xdist, ydist)
		} else {

		}
	}

	return total
}

func RunDayEight() {
	antennas := setupdayeight()
	fmt.Println("d8p1:", getAntinodes(antennas))
}
