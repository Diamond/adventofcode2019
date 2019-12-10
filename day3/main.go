package main

import (
	"fmt"
	"strconv"
)

type point struct {
	x int
	y int
}

func convertCoord(coord string) (string, int) {
	lhs := string(coord[0])
	rhs, err := strconv.Atoi(coord[1:])
	if err != nil {
		return "X", 0
	}
	return lhs, rhs
}

func walkPath(direction string, distance int, start *point) ([]point, point) {
	start := point{x: 0, y: 0}
	coords := make([]point, distance)
	switch direction {
	case "R":
		for i := 0; i < distance; i++ {
			coords[i] = point{x: start.x + i, y: start.y}
		}
	case "L":
		for i := 0; i < distance; i++ {
			coords[i] = point{x: start.x - i, y: start.y}
		}
	case "U":
		for i := 0; i < distance; i++ {
			coords[i] = point{x: start.x, y: start.y - i}
		}
	case "D":
		for i := 0; i < distance; i++ {
			coords[i] = point{x: start.x, y: start.y + i}
		}
	}
	return nil, coords
}

func main() {
	coords := []string{"R8", "U5", "L5", "D3"}
	for _, coord := range coords {
		direction, distance := convertCoord(coord)
		fmt.Println(walkPath())
	}
}
