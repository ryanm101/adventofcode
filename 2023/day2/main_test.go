package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestAOC2a(t *testing.T) {
	config := map[string]int{"red": 12, "green": 13, "blue": 14}

	input := []string{"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green"}

	expected := 8

	res := AOC2a(input, config)
	if res != expected {
		t.Errorf("AOC2a: %d != %d", res, expected)
	}
}

func TestAOC2b(t *testing.T) {
	config := map[string]int{"red": 12, "green": 13, "blue": 14}

	input := []string{"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green"}

	expected := 2286

	res := AOC2b(input, config)
	if res != expected {
		t.Errorf("AOC2b: %d != %d", res, expected)
	}
}

func TestAOC2GetGame(t *testing.T) {
	config := map[string]int{"red": 12, "green": 13, "blue": 14}

	input := []string{"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green"}

	expectedGames := []Game{
		{Id: 1,
			Sets:     []map[string]int{{"red": 4, "blue": 3}, {"red": 1, "green": 2, "blue": 6}, {"green": 2}},
			Highest:  map[string]int{"red": 4, "blue": 6, "green": 2},
			Possible: true},
		{Id: 2,
			Sets:     []map[string]int{{"blue": 1, "green": 2}, {"green": 3, "blue": 4, "red": 1}, {"green": 1, "blue": 1}},
			Highest:  map[string]int{"blue": 4, "green": 3, "red": 1},
			Possible: true},
		{Id: 3,
			Sets:     []map[string]int{{"green": 8, "blue": 6, "red": 20}, {"blue": 5, "red": 4, "green": 13}, {"green": 5, "red": 1}},
			Highest:  map[string]int{"green": 13, "blue": 6, "red": 20},
			Possible: false},
		{Id: 4,
			Sets:     []map[string]int{{"green": 1, "red": 3, "blue": 6}, {"green": 3, "red": 6}, {"green": 3, "blue": 15, "red": 14}},
			Highest:  map[string]int{"green": 3, "red": 14, "blue": 15},
			Possible: false},
		{Id: 5,
			Sets:     []map[string]int{{"red": 6, "blue": 1, "green": 3}, {"blue": 2, "red": 1, "green": 2}},
			Highest:  map[string]int{"red": 6, "blue": 2, "green": 3},
			Possible: true},
	}

	for i, v := range input {
		actual := AOC2GetGame(v, config)
		if !cmp.Equal(actual, expectedGames[i]) {
			t.Errorf("AOC2GetGame(%s): %v != %v\n\tAcutal Id: %d - ExpectedId: %d\n\tActualSets: %v - ExpectedSets: %v\n\tActualHighest: %v - ExpectedHighest: %v\n\tActualPossible: %v - ExpectedPossible: %v",
				v, actual, expectedGames[i], actual.Id, expectedGames[i].Id, actual.Sets, expectedGames[i].Sets, actual.Highest, expectedGames[i].Highest, actual.Possible, expectedGames[i].Possible)
		}
	}
}

func TestAOC2GetSet(t *testing.T) {
	input := []string{"3 blue, 4 red", "1 red, 2 green, 6 blue", "2 green"}
	expected := []map[string]int{{"blue": 3, "red": 4}, {"red": 1, "green": 2, "blue": 6}, {"green": 2}}

	for i, v := range input {
		if !cmp.Equal(AOC2GetSet(v), expected[i]) {
			t.Errorf("AOC2GetSet(%s) != %v", v, expected[i])
		}
	}
}
