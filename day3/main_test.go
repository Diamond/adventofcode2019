package main

import "testing"

func TestThing(t *testing.T) {
	val := true
	if !val {
		t.Error("Failure")
	}
}
