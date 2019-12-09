package main

import "testing"

// TestFuelRequired tests the fuel function against the sample inputs
func TestFuelRequired(t *testing.T) {
	tables := []struct {
		mass int
		fuel int
	}{
		{12, 2},
		{14, 2},
		{1969, 654},
		{100756, 33583},
	}
	for _, table := range tables {
		f := FindFuelRequired(table.mass)
		if f != table.fuel {
			t.Errorf("Fuel %d did not match expected value %d", f, table.fuel)
		}
	}
}

func TestGetFuelTotal(t *testing.T) {
	inputs := []int{12, 14, 1969, 100756}
	total := GetFuelTotal(inputs)
	expected := 34241
	if total != expected {
		t.Errorf("Total Fuel for masses did not equal %d, instead got %d", expected, total)
	}
}

func TestGetPerItemFuelForFuel(t *testing.T) {
	tables := []struct {
		mass []int
		fuel int
	}{
		{[]int{14}, 2},
		{[]int{1969}, 966},
		{[]int{100756}, 50346},
	}
	for _, table := range tables {
		f := GetPerItemFuelForFuel(table.mass)
		if f != table.fuel {
			t.Errorf("Fuel %d did not match expected value %d", f, table.fuel)
		}
	}
}
