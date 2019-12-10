package main

import "fmt"

// HandleOpcode handles opcode
func HandleOpcode(input []int, pos int) ([]int, int) {
	opcode := make([]int, len(input))
	copy(opcode, input)
	op := opcode[pos]
	if op == 99 || pos == -1 {
		return opcode, -1
	}
	if op == 1 {
		// Add opcode
		lhsPos := opcode[pos+1]
		rhsPos := opcode[pos+2]
		lhsVal := opcode[lhsPos]
		rhsVal := opcode[rhsPos]
		valPos := opcode[pos+3]
		opcode[valPos] = lhsVal + rhsVal
		return HandleOpcode(opcode, pos+4)
	}
	if op == 2 {
		// Multiply opcode
		lhsPos := opcode[pos+1]
		rhsPos := opcode[pos+2]
		lhsVal := opcode[lhsPos]
		rhsVal := opcode[rhsPos]
		valPos := opcode[pos+3]
		opcode[valPos] = lhsVal * rhsVal
		return HandleOpcode(opcode, pos+4)
	}
	return opcode, -1
}

// FixInput corrects the pre-fire state of the application
func FixInput(input []int, noun int, verb int) []int {
	opcode := make([]int, len(input))
	copy(opcode, input)
	opcode[1] = noun
	opcode[2] = verb
	return opcode
}

// ScoreInput scores the noun and verbs against the score algorithm (100 * n + v)
func ScoreInput(noun int, verb int) int {
	return 100*noun + verb
}

// CalcInputs figures out what combination of inputs returns the expected score
func CalcInputs(startOpcode []int, expected int) int {
	for lhs := 0; lhs < 100; lhs++ {
		for rhs := 0; rhs < 100; rhs++ {
			newOpcode := FixInput(startOpcode, lhs, rhs)
			resultOpcode, resultCode := HandleOpcode(newOpcode, 0)
			if resultOpcode[0] == expected && resultCode == -1 {
				score := ScoreInput(lhs, rhs)
				return score
			}
		}
	}
	return -1
}

func part1(start []int) {
	opcode := make([]int, len(start))
	copy(opcode, start)
	realOpcode := FixInput(opcode, 12, 2)
	finalOpcode, resultCode := HandleOpcode(realOpcode, 0)
	if resultCode == -1 {
		fmt.Println("Final result at position 0:", finalOpcode[0])
	} else {
		fmt.Println("Received result code:", resultCode)
	}
}

func part2(start []int) {
	fmt.Println("Part 2")
	fmt.Println("--------------------------------------")
	finalScore := CalcInputs(start, 19690720)
	fmt.Println("Final Score:", finalScore)
}

func main() {
	realOpcode := []int{1, 0, 0, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 1, 13, 19, 1, 10, 19, 23, 2, 9, 23, 27, 1, 6, 27, 31, 1, 10, 31, 35, 1, 35, 10, 39, 1, 9, 39, 43, 1, 6, 43, 47, 1, 10, 47, 51, 1, 6, 51, 55, 2, 13, 55, 59, 1, 6, 59, 63, 1, 10, 63, 67, 2, 67, 9, 71, 1, 71, 5, 75, 1, 13, 75, 79, 2, 79, 13, 83, 1, 83, 9, 87, 2, 10, 87, 91, 2, 91, 6, 95, 2, 13, 95, 99, 1, 10, 99, 103, 2, 9, 103, 107, 1, 107, 5, 111, 2, 9, 111, 115, 1, 5, 115, 119, 1, 9, 119, 123, 2, 123, 6, 127, 1, 5, 127, 131, 1, 10, 131, 135, 1, 135, 6, 139, 1, 139, 5, 143, 1, 143, 9, 147, 1, 5, 147, 151, 1, 151, 13, 155, 1, 5, 155, 159, 1, 2, 159, 163, 1, 163, 6, 0, 99, 2, 0, 14, 0}
	part1(realOpcode)
	part2(realOpcode)
}
