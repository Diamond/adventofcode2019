package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func makeGrid(width, height int) [][]int {
// 	grid := make([][]int, height)
// 	for y := 0; y < height; y++ {
// 		grid[y] = make([]int, width)
// 	}
// 	return grid
// }

// func printGrid(input [][]int) {
// 	for y := 0; y < len(input); y++ {
// 		for x := 0; x < len(input[y]); x++ {
// 			if input[y][x] == 0 {
// 				fmt.Print(".")
// 			} else if input[y][x] == 1 {
// 				fmt.Print("-")
// 			} else if input[y][x] == 2 {
// 				fmt.Print("|")
// 			} else if input[y][x] >= 3 {
// 				fmt.Print("+")
// 			} else if input[y][x] == -1 {
// 				fmt.Print("o")
// 			} else {
// 				fmt.Print("?")
// 			}
// 		}
// 		fmt.Print("\n")
// 	}
// }

// func drawCoord(input [][]int, coord string, startX int, startY int) ([][]int, int, int) {
// 	// Copy the input first
// 	gridCopy := makeGrid(len(input[0]), len(input))
// 	copy(gridCopy, input)
// 	nextX, nextY := 0, 0
// 	// Next, parse the coord to get the vector (direction and distance)
// 	direction, distance := convertCoord(coord)
// 	switch direction {
// 	case "R":
// 		for x := startX + 1; x < len(input[startY]) && x <= (startX+distance); x++ {
// 			gridCopy[startY][x] += 1
// 			nextX = x
// 			nextY = startY
// 		}
// 	case "L":
// 		for x := startX - 1; x >= 0 && x >= (startX-distance); x-- {
// 			gridCopy[startY][x] += 1
// 			nextX = x
// 			nextY = startY
// 		}
// 	case "U":
// 		for y := startY - 1; y >= 0 && y >= (startY-distance); y-- {
// 			gridCopy[y][startX] += 2
// 			nextX = startX
// 			nextY = y
// 		}
// 	case "D":
// 		for y := startY + 1; y < len(input) && y <= (startY+distance); y++ {
// 			gridCopy[y][startX] += 2
// 			nextX = startX
// 			nextY = y
// 		}
// 	}
// 	return gridCopy, nextX, nextY
// }

// func convertCoord(coord string) (string, int) {
// 	lhs := string(coord[0])
// 	rhs, err := strconv.Atoi(coord[1:])
// 	if err != nil {
// 		return "X", 0
// 	}
// 	return lhs, rhs
// }

// func startCoords(inputs []string) [][]int {
// 	grid := makeGrid(11, 10)
// 	nx := 1
// 	ny := 8
// 	grid[ny][nx] = -1
// 	for _, coord := range inputs {
// 		grid, nx, ny = drawCoord(grid, coord, nx, ny)
// 	}
// 	return grid
// }

// func addCoords(grid [][]int, inputs []string) [][]int {
// 	gridCopy := makeGrid(len(grid[0]), len(grid))
// 	copy(gridCopy, grid)
// 	nx := 1
// 	ny := 8
// 	gridCopy[ny][nx] = -1
// 	for _, coord := range inputs {
// 		gridCopy, nx, ny = drawCoord(gridCopy, coord, nx, ny)
// 	}
// 	return gridCopy
// }

// func flipGrid(grid [][]int) [][]int {
// 	gridCopy := makeGrid(len(grid[0]), len(grid))

// 	for y := 0; y < len(grid); y++ {
// 		ly := (len(grid) - 1) - y
// 		for x := 0; x < len(grid[y]); x++ {
// 			gridCopy[y][x] = grid[ly][x]
// 		}
// 	}
// 	return gridCopy
// }

// func findClosestIntersections(grid [][]int) int {
// 	sx, sy := 1, 1
// 	lx, ly := -1, -1
// 	for y := 1; y < len(grid); y++ {
// 		for x := 1; x < len(grid[y]); x++ {
// 			if grid[y][x] >= 3 {
// 				// Intersection found
// 				if lx == -1 && ly == -1 || ((lx + ly) > (x + y)) {
// 					// Lowest intersection
// 					lx, ly = x, y
// 				}
// 			}
// 		}
// 	}
// 	return (lx - sx) + (ly - sy)
// }

// func main() {
// 	coords := []string{
// 		"R8", "U5", "L5", "D3",
// 	}
// 	coords2 := []string{
// 		"U7", "R6", "D4", "L4",
// 	}
// 	result := startCoords(coords)
// 	result2 := addCoords(result, coords2)
// 	printGrid(result2)
// 	flipped := flipGrid(result2)
// 	distance := findIntersections(flipped)
// 	fmt.Println("Manhattan distance is:", distance)
// }
