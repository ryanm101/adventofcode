package main

import (
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var rawTestData = []string{
	"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
	"Card 2: 13 32 20 16 61 | 61 30 68 82  17 32 24 19",
	"Card 3:  1 21  53 59 44 | 69 82 63 72 16 21 14  1",
	"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
	"Card 5: 87 83 26 28 32 | 88 30 70 12  93 22 82 36",
	"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
}

var structTestData = []ScratchCard{
	{Id: 1, WinningNumbers: []int{41, 48, 83, 86, 17}, Numbers: []int{83, 86, 6, 31, 17, 9, 48, 53}, Wincount: 4, InstanceCount: 1},
	{Id: 2, WinningNumbers: []int{13, 32, 20, 16, 61}, Numbers: []int{61, 30, 68, 82, 17, 32, 24, 19}, Wincount: 2, InstanceCount: 1},
	{Id: 3, WinningNumbers: []int{1, 21, 53, 59, 44}, Numbers: []int{69, 82, 63, 72, 16, 21, 14, 1}, Wincount: 2, InstanceCount: 1},
	{Id: 4, WinningNumbers: []int{41, 92, 73, 84, 69}, Numbers: []int{59, 84, 76, 51, 58, 5, 54, 83}, Wincount: 1, InstanceCount: 1},
	{Id: 5, WinningNumbers: []int{87, 83, 26, 28, 32}, Numbers: []int{88, 30, 70, 12, 93, 22, 82, 36}, Wincount: 0, InstanceCount: 1},
	{Id: 6, WinningNumbers: []int{31, 18, 13, 56, 72}, Numbers: []int{74, 77, 10, 23, 35, 67, 36, 11}, Wincount: 0, InstanceCount: 1},
}

var copyCounts = []int{1, 2, 4, 8, 14, 1}

func TestAOC4a(t *testing.T) {
	input := rawTestData
	expectedPoints := 13
	expectedCards := 30

	resPoints, resCards := AOC4a(input)
	if resPoints != expectedPoints {
		t.Errorf("AOC4a: %d != %d", resPoints, expectedPoints)
	}
	if resCards != expectedCards {
		t.Errorf("AOC4a: %d != %d", resCards, expectedCards)
	}
}

func TestAddCopy(t *testing.T) {
	input := []ScratchCard{
		{Id: 1, WinningNumbers: []int{}, Numbers: []int{}, Wincount: 4, InstanceCount: 1},
		{Id: 2, WinningNumbers: []int{}, Numbers: []int{}, Wincount: 2, InstanceCount: 1},
		{Id: 3, WinningNumbers: []int{}, Numbers: []int{}, Wincount: 2, InstanceCount: 1},
		{Id: 4, WinningNumbers: []int{}, Numbers: []int{}, Wincount: 1, InstanceCount: 1},
		{Id: 5, WinningNumbers: []int{}, Numbers: []int{}, Wincount: 0, InstanceCount: 1},
		{Id: 6, WinningNumbers: []int{}, Numbers: []int{}, Wincount: 0, InstanceCount: 1},
	}
	expected := []ScratchCard{
		{Id: 1, WinningNumbers: []int{}, Numbers: []int{}, Wincount: 4, InstanceCount: 1},
		{Id: 2, WinningNumbers: []int{}, Numbers: []int{}, Wincount: 2, InstanceCount: 2},
		{Id: 3, WinningNumbers: []int{}, Numbers: []int{}, Wincount: 2, InstanceCount: 4},
		{Id: 4, WinningNumbers: []int{}, Numbers: []int{}, Wincount: 1, InstanceCount: 8},
		{Id: 5, WinningNumbers: []int{}, Numbers: []int{}, Wincount: 0, InstanceCount: 14},
		{Id: 6, WinningNumbers: []int{}, Numbers: []int{}, Wincount: 0, InstanceCount: 1},
	}

	for i := range expected {
		expected[i].InstanceCount = copyCounts[i]
	}

	for _, v := range input {
		AddCopy(&input, v.Id)
	}
	if !cmp.Equal(input, expected) {
		t.Errorf("GetCopies: %v != %v", input, expected)
	}
}

func TestSplitData(t *testing.T) {
	input := rawTestData
	expected := structTestData
	for i := range expected {
		expected[i].Wincount = 0
	}

	res := SplitData(input)
	if !cmp.Equal(res, expected) {
		t.Errorf("SplitData: %v != %v", res, expected)
	}
}

func TestStrToArrInt(t *testing.T) {
	input := "1 2 3 4 5"
	expected := []int{1, 2, 3, 4, 5}

	res := StrToArrInt(input)
	if !cmp.Equal(res, expected) {
		t.Errorf("StrToArrInt: %v != %v", res, expected)
	}
}

func TestFindWinners(t *testing.T) {
	input := structTestData
	expected := [][]int{{48, 83, 86, 17}, {32, 61}, {1, 21}, {84}, {}, {}}

	for i, v := range input {
		res := FindWinners(v)
		sort.Ints(res)
		sort.Ints(expected[i])
		if !cmp.Equal(res, expected[i]) {
			t.Errorf("FindWinners: %v != %v", res, expected[i])
		}
	}
}

func TestPow2(t *testing.T) {
	input := 4
	expected := 8

	res := Pow2(input)
	if res != expected {
		t.Errorf("pow2: %d != %d", res, expected)
	}
}
