package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestAOC3a(t *testing.T) {
	input := []string{"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	}

	expected := 4361

	res := AOC3a(input)
	if res != expected {
		t.Errorf("AOC3a: %d != %d", res, expected)
	}
}

func TestSymbolScanner(t *testing.T) {
	input := []string{"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	}
	expected := []Coord{{1, 3}, {3, 6}, {4, 3}, {5, 5}, {8, 3}, {8, 5}}

	res := SymbolScanner(input)
	if !cmp.Equal(res, expected) {
		t.Errorf("SymbolScanner: %v != %v", res, expected)
	}
}

func TestValidParts(t *testing.T) {
	input := []string{"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	}
	symbolLocs := []Coord{{1, 3}, {3, 6}, {4, 3}, {5, 5}, {8, 3}, {8, 5}}

	expected := []Part{{StartIndex: 0, EndIndex: 3, LineNumber: 0},
		{StartIndex: 2, EndIndex: 4, LineNumber: 2},
		{StartIndex: 6, EndIndex: 9, LineNumber: 2},
		{StartIndex: 0, EndIndex: 3, LineNumber: 4},
		{StartIndex: 2, EndIndex: 5, LineNumber: 6},
		{StartIndex: 6, EndIndex: 9, LineNumber: 7},
		{StartIndex: 1, EndIndex: 4, LineNumber: 9},
		{StartIndex: 5, EndIndex: 8, LineNumber: 9}}

	res := []Part{}
	for i, line := range input {
		res = append(res, ValidParts(line, i, symbolLocs)...)
	}

	if !cmp.Equal(res, expected) {
		t.Errorf("SymbolScanner: %v != %v", res, expected)
	}
}

func TestGetValidPositions(t *testing.T) {
	symbolLocs := []Coord{{1, 3}, {3, 6}, {4, 3}, {5, 5}, {8, 3}, {8, 5}}

	// We dont have any dedup so inefficient
	expected := []Coord{
		{0, 2}, {0, 3}, {0, 4}, {1, 2}, {1, 4}, {2, 2}, {2, 3}, {2, 4}, // {1,3}
		{2, 5}, {2, 6}, {2, 7}, {3, 5}, {3, 7}, {4, 5}, {4, 6}, {4, 7}, // {3,6}
		{3, 2}, {3, 3}, {3, 4}, {4, 2}, {4, 4}, {5, 2}, {5, 3}, {5, 4}, // {4,3}
		{4, 4}, {4, 5}, {4, 6}, {5, 4}, {5, 6}, {6, 4}, {6, 5}, {6, 6}, // {5,5}
		{7, 2}, {7, 3}, {7, 4}, {8, 2}, {8, 4}, {9, 2}, {9, 3}, {9, 4}, // {8,3}
		{7, 4}, {7, 5}, {7, 6}, {8, 4}, {8, 6}, {9, 4}, {9, 5}, {9, 6}, // {8,5}
	}

	res := GetValidPositions(symbolLocs)

	if !cmp.Equal(res, expected) {
		t.Errorf("SymbolScanner: %v != %v", res, expected)
	}
}

func TestGetPartNumber(t *testing.T) {
	input := []string{"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	}
	Parts := []Part{{StartIndex: 0, EndIndex: 3, LineNumber: 0},
		{StartIndex: 2, EndIndex: 4, LineNumber: 2},
		{StartIndex: 6, EndIndex: 9, LineNumber: 2},
		{StartIndex: 0, EndIndex: 3, LineNumber: 4},
		{StartIndex: 2, EndIndex: 5, LineNumber: 6},
		{StartIndex: 6, EndIndex: 9, LineNumber: 7},
		{StartIndex: 1, EndIndex: 4, LineNumber: 9},
		{StartIndex: 5, EndIndex: 8, LineNumber: 9}}

	expected := []int{467, 35, 633, 617, 592, 755, 664, 598}

	res := GetPartNumber(input, Parts)

	if !cmp.Equal(res, expected) {
		t.Errorf("SymbolScanner: %v != %v", res, expected)
	}
}
