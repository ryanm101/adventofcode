package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	println(fmt.Sprintf("Day1a: %d", AOC1a()))
	println(fmt.Sprintf("Day1b: %d", AOC1b()))
}

func AOC1a() int {
	infile := "day1_input.txt"
	file, err := os.Open(infile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	counter := 0
	for fileScanner.Scan() {
		cval := AOC1aGetInt(fileScanner.Text())
		counter += cval
	}
	return counter
}

func AOC1b() int {
	infile := "day1_input.txt"
	file, err := os.Open(infile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	counter := 0
	for fileScanner.Scan() {
		cval := AOC1bGetInt(fileScanner.Text())
		counter += cval
	}
	return counter
}

func AOC1aGetInt(s string) int {
	firstNum := regexp.MustCompile(`^.*?(\d).*?$`)
	lastNum := regexp.MustCompile(`^.*(\d).*?$`)
	strInt := strings.Join([]string{firstNum.FindAllStringSubmatch(s, -1)[0][1], lastNum.FindAllStringSubmatch(s, -1)[0][1]}, "")
	i, err := strconv.Atoi(strInt)
	if err != nil {
		panic(err)
	}
	return i
}

func AOC1bGetInt(s string) int {
	firstNum := regexp.MustCompile(`^.*?((?:\d|zero|one|two|three|four|five|six|seven|eight|nine)).*?$`)
	lastNum := regexp.MustCompile(`^.*((?:\d|zero|one|two|three|four|five|six|seven|eight|nine)).*?$`)

	strInt := strings.Join([]string{AOC1StrToNum(firstNum.FindAllStringSubmatch(s, -1)[0][1]),
		AOC1StrToNum(lastNum.FindAllStringSubmatch(s, -1)[0][1])}, "")

	i, err := strconv.Atoi(strInt)
	if err != nil {
		panic(err)
	}
	return i
}

func AOC1StrToNum(s string) string {
	aNum := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9, "zero": 0}
	i, err := strconv.Atoi(s)
	if err != nil {
		i = aNum[s]
	}
	return fmt.Sprint((i))
}
