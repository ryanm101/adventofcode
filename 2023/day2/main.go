package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	Id       int
	Sets     []map[string]int
	Highest  map[string]int
	Possible bool
}

func main() {
	config := map[string]int{"red": 12, "green": 13, "blue": 14}
	data := AOC2ReadFile("input.txt")

	println(fmt.Sprintf("Day2a: %d", AOC2a(data, config)))
	println(fmt.Sprintf("Day2b: %d", AOC2b(data, config)))
}

func AOC2ReadFile(fname string) []string {
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

func AOC2a(data []string, config map[string]int) int {
	counter := 0
	for _, line := range data {
		game := AOC2GetGame(line, config)
		if game.Possible {
			counter = counter + game.Id
		}
	}
	return counter
}

func AOC2b(data []string, config map[string]int) int {
	counter := 0
	for _, line := range data {
		game := AOC2GetGame(line, config)
		a := 0
		for _, v := range game.Highest {
			if a == 0 {
				a = v
			} else {
				a = a * v
			}
		}
		counter = counter + a
	}
	return counter
}

func AOC2GetGame(s string, config map[string]int) Game {
	game := Game{}
	game.Highest = map[string]int{"red": 0, "green": 0, "blue": 0}

	game.Id, _ = strconv.Atoi(strings.Split(strings.Split(s, ":")[0], " ")[1])
	s = strings.Replace(s, "; ", ";", -1)
	s = strings.Replace(s, ": ", ":", -1)
	sets := strings.Split(strings.Split(s, ":")[1], ";")
	game.Highest = make(map[string]int)

	for _, v := range sets {
		set := AOC2GetSet(v)
		game.Sets = append(game.Sets, set)
		for k, v := range set {
			if game.Highest[k] < v {
				game.Highest[k] = v
			}
		}
	}

	for k, v := range game.Highest {
		game.Possible = true
		if config[k] < v {
			game.Possible = false
			break
		}
		//println(fmt.Sprintf("ID: %d - GH Key %s: GH Value %d - CV %d, Poss: %v", game.Id, k, v, config[k], game.Possible))
	}

	return game
}

func AOC2GetSet(s string) map[string]int {
	set := make(map[string]int)
	s = strings.Replace(s, ", ", ",", -1)
	for _, v := range strings.Split(s, ",") {
		set[strings.Split(v, " ")[1]], _ = strconv.Atoi(strings.Split(v, " ")[0])
	}
	return set
}
