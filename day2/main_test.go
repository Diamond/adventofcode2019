package main

import (
	"fmt"
	"testing"
)

func compareOpcode(in []int, out []int) bool {
	for i := 0; i < len(in); i++ {
		if in[i] != out[i] {
			fmt.Println("Failure: in ->", in[i], "out ->", out[i])
			return false
		}
	}
	return true
}

func TestThing(t *testing.T) {
	table := []struct {
		inputOpcode  []int
		outputOpcode []int
	}{
		{inputOpcode: []int{1, 0, 0, 0, 99}, outputOpcode: []int{2, 0, 0, 0, 99}},
		{inputOpcode: []int{2, 3, 0, 3, 99}, outputOpcode: []int{2, 3, 0, 6, 99}},
		{inputOpcode: []int{2, 4, 4, 5, 99, 0}, outputOpcode: []int{2, 4, 4, 5, 99, 9801}},
		{inputOpcode: []int{1, 1, 1, 4, 99, 5, 6, 0, 99}, outputOpcode: []int{30, 1, 1, 4, 2, 5, 6, 0, 99}},
	}
	for line, input := range table {
		newOpcode, resultCode := HandleOpcode(input.inputOpcode, 0)
		if !compareOpcode(newOpcode, input.outputOpcode) || resultCode != -1 {
			t.Error("Error, expected opcode to match at line", line)
		}
	}
}

func TestFixInput(t *testing.T) {
	input := []int{99, 37, 21, 62}
	expected := []int{99, 12, 2, 62}
	if !compareOpcode(FixInput(input, 12, 2), expected) {
		t.Error("FixInput did not correct values as expected")
	}
}

func TestScoreInput(t *testing.T) {
	input := []int{12, 2}
	expected := 1202
	if ScoreInput(input[0], input[1]) != expected {
		t.Error("Expected value did not match score")
	}
}
