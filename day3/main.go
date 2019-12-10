package main

import (
	"fmt"
	"math"
	"strconv"
)

func findGridSizePerCoords(coords []string) (int, int, int, int) {
	left, right, top, bottom, cursorX, cursorY := 0, 0, 0, 0, 0, 0
	for _, coord := range coords {
		direction, distance := convertCoord(coord)
		switch direction {
		case "R":
			temp := cursorX + distance
			if temp > right {
				right = temp
			}
			cursorX += distance
		case "L":
			temp := cursorX - distance
			if temp < left {
				left = temp
			}
			cursorX -= distance
		case "U":
			temp := cursorY - distance
			if temp < top {
				top = temp
			}
			cursorY -= distance
		case "D":
			temp := cursorY + distance
			if temp > bottom {
				bottom = temp
			}
			cursorY += distance
		}
	}
	dx := right - left + 1
	dy := bottom - top + 1
	sx, sy := int(math.Abs(float64(left))), int(math.Abs(float64(top)))
	if left < 0 {
		sx--
	}
	return dx, dy, sx, sy
}

func convertCoord(coord string) (string, int) {
	lhs := string(coord[0])
	rhs, err := strconv.Atoi(coord[1:])
	if err != nil {
		return "X", 0
	}
	return lhs, rhs
}

func makeGrid(width, height int) [][]int {
	grid := make([][]int, height)
	for y := 0; y < height; y++ {
		grid[y] = make([]int, width)
	}
	return grid
}

func printGrid(input [][]int) {
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			if input[y][x] == 0 {
				fmt.Print(".")
			} else if input[y][x] == 1 {
				fmt.Print("#")
			} else if input[y][x] == 2 {
				fmt.Print("|")
			} else if input[y][x] >= 3 {
				fmt.Print("+")
			} else if input[y][x] == -1 {
				fmt.Print("o")
			} else {
				fmt.Print("?")
			}
		}
		fmt.Print("\n")
	}
}

func startCoords(grid [][]int, inputs []string, sx, sy int) [][]int {
	nx := sx
	ny := sy
	grid[ny][nx] = -1
	for _, coord := range inputs {
		grid, nx, ny = drawCoord(grid, coord, nx, ny)
	}
	return grid
}

func drawCoord(input [][]int, coord string, startX int, startY int) ([][]int, int, int) {
	// Copy the input first
	gridCopy := makeGrid(len(input[0]), len(input))
	copy(gridCopy, input)
	nextX, nextY := 0, 0
	// Next, parse the coord to get the vector (direction and distance)
	direction, distance := convertCoord(coord)
	fmt.Printf("Starting at %d, %d drawing %s for distance %d\n", startX, startY, direction, distance)
	switch direction {
	case "R":
		for x := startX + 1; x < len(input[startY]) && x <= (startX+distance); x++ {
			gridCopy[startY][x] = 1
			nextX = x
			nextY = startY
		}
	case "L":
		for x := startX - 1; x >= 0 && x >= (startX-distance); x-- {
			gridCopy[startY][x] = 1
			nextX = x
			nextY = startY
		}
	case "U":
		for y := startY - 1; y >= 0 && y >= (startY-distance); y-- {
			gridCopy[y][startX] = 1
			nextX = startX
			nextY = y
		}
	case "D":
		for y := startY + 1; y < len(input) && y <= (startY+distance); y++ {
			gridCopy[y][startX] = 1
			nextX = startX
			nextY = y
		}
	}
	return gridCopy, nextX, nextY
}

func main() {
	coords := []string{
		"R8", "U5", "L5", "D3",
	}
	w, h, sx, sy := findGridSizePerCoords(coords)
	fmt.Printf("Grid should be %d X %d\n", w, h)
	fmt.Printf("And start at %d, %d\n", sx, sy)

	fmt.Println("\n\nMaking the grid")
	grid := makeGrid(w, h)
	printGrid(grid)

	fmt.Println("\n\nTime to draw the coords")
	drawnGrid := startCoords(grid, coords, sx, sy)
	printGrid(drawnGrid)

}
