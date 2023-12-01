package main

import "testing"

func TestAOC1aGetInt(t *testing.T) {
	indata := []string{"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet"}
	expected_output := []int{12, 38, 15, 77}

	for i, v := range indata {
		if AOC1aGetInt(v) != expected_output[i] {
			t.Errorf("AOC1(%s) != %d", v, expected_output[i])
		}
	}

}

func TestAOC1GetIntSTR(t *testing.T) {
	indata := []string{"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
	}
	expected_output := []int{29, 83, 13, 24, 42, 14, 76}

	for i, v := range indata {
		if AOC1bGetInt(v) != expected_output[i] {
			t.Errorf("AOC1(%s) != %d", v, expected_output[i])
		}
	}

}

func TestAOC1StrToNum(t *testing.T) {
	indata := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "zero", "1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
	expected_output := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}

	for i, v := range indata {
		if AOC1StrToNum(v) != expected_output[i] {
			t.Errorf("AOC1(%s) != %s", v, expected_output[i])
		}
	}
}
