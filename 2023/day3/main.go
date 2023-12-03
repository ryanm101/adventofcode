package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Coord struct {
	Row int
	Col int
}

type Part struct {
	StartIndex int
	EndIndex   int
	LineNumber int
}

func AOCReadFile(fname string) []string {
	file, err := os.Open(fname)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	lines := []string{}
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}
	return lines
}

func main() {
	input := AOCReadFile("input.txt")

	println(fmt.Sprintf("AOC3a: %d", AOC3a(input)))
}

func AOC3a(input []string) int {
	symbolLocs := SymbolScanner(input)
	validParts := []Part{}
	for i, line := range input {
		validParts = append(validParts, ValidParts(line, i, symbolLocs)...)
	}
	partNumbers := GetPartNumber(input, validParts)
	sum := 0
	for _, partNumber := range partNumbers {
		sum += partNumber
	}
	return sum
}

func SymbolScanner(s []string) []Coord {
	symbols := regexp.MustCompile(`[#$%&*+\-\/=@]`)
	symbolslocs := []Coord{}
	for i, line := range s {
		matches := symbols.FindAllIndex([]byte(line), -1)
		if matches == nil {
			continue
		}

		for _, match := range matches {
			symbolslocs = append(symbolslocs, Coord{i, match[0]})
		}
	}
	return symbolslocs
}

func GetValidPositions(itemCoords []Coord) []Coord {
	validPositions := []Coord{}
	// We dont have any dedup so inefficient
	for _, coord := range itemCoords {
		tl, tm, tr, ml, mr, bl, bm, br := Coord{coord.Row - 1, coord.Col - 1}, Coord{coord.Row - 1, coord.Col}, Coord{coord.Row - 1, coord.Col + 1},
			Coord{coord.Row, coord.Col - 1}, Coord{coord.Row, coord.Col + 1},
			Coord{coord.Row + 1, coord.Col - 1}, Coord{coord.Row + 1, coord.Col}, Coord{coord.Row + 1, coord.Col + 1}
		validPositions = append(validPositions, tl, tm, tr, ml, mr, bl, bm, br)
	}
	return validPositions
}

func ValidParts(line string, lineNumber int, symbolocs []Coord) []Part {
	numbers := regexp.MustCompile(`[0-9]+`)
	validParts := []Part{}

	// find all numbers
	matches := numbers.FindAllIndex([]byte(line), -1)
	//fmt.Printf("line %d: %v\n", lineNumber, matches)

	validPartsPositions := GetValidPositions(symbolocs)

	for _, match := range matches {
		//fmt.Printf("line %d: %v\n", lineNumber, match)
		found := false
		for i := match[0]; i < match[1]; i++ {
			digitCoord := Coord{lineNumber, i}
			for _, c := range validPartsPositions {
				if c == digitCoord {
					validParts = append(validParts, Part{StartIndex: match[0], EndIndex: match[1], LineNumber: lineNumber})
					found = true
					break
				}
				if found {
					break
				}
			}
		}
	}

	return validParts
}

func GetPartNumber(input []string, parts []Part) []int {
	partNumbers := []int{}
	for _, part := range parts {
		line := input[part.LineNumber]
		partNumber, _ := strconv.Atoi(line[part.StartIndex:part.EndIndex])
		partNumbers = append(partNumbers, partNumber)
	}
	return partNumbers
}
