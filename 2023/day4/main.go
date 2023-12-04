package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

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

type ScratchCard struct {
	Id             int
	WinningNumbers []int
	Numbers        []int
	Wincount       int
	InstanceCount  int
}

func main() {
	input := AOCReadFile("input.txt")

	points, cards := AOC4a(input)
	println(fmt.Sprintf("AOC4a: %d", points))
	println(fmt.Sprintf("AOC4b: %d", cards))
}

func AOC4a(input []string) (int, int) {
	scratchCards := SplitData(input)

	var points int = 0
	var cards int = 0

	for i, s := range scratchCards {
		WinningNumbers := FindWinners(s)
		scratchCards[i].Wincount = len(WinningNumbers)
		if WinningNumbers != nil {
			if scratchCards[i].Wincount == 1 {
				points = points + 1
				continue
			}
			points += Pow2(scratchCards[i].Wincount)
		}
	}

	for _, s := range scratchCards {
		AddCopy(&scratchCards, s.Id)
	}

	for _, s := range scratchCards {
		cards += s.InstanceCount
	}

	return points, cards
}

func AddCopy(scs *[]ScratchCard, Id int) {
	cs := *scs
	gid := Id - 1
	for i := 0; i < cs[gid].InstanceCount; i++ {
		for j := 1; j <= cs[gid].Wincount; j++ {
			cs[gid+j].InstanceCount += 1
		}
	}
}

func SplitData(s []string) []ScratchCard {
	scratchCards := []ScratchCard{}
	space := regexp.MustCompile(`\s+`)

	for _, line := range s {
		scratchCard := ScratchCard{0, []int{}, []int{}, 0, 1}
		line = space.ReplaceAllString(line, " ")
		s1 := strings.Split(line, ":")
		numbers := strings.Split(s1[1], "|")

		scratchCard.Id, _ = strconv.Atoi(strings.Split(s1[0], " ")[1])
		scratchCard.WinningNumbers = StrToArrInt(numbers[0])
		scratchCard.Numbers = StrToArrInt(numbers[1])

		scratchCards = append(scratchCards, scratchCard)
	}
	return scratchCards
}

func StrToArrInt(s string) []int {
	arr := []int{}
	for _, v := range strings.Split(s, " ") {
		if v != "" {
			i, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			arr = append(arr, i)
		}
	}
	return arr
}

func FindWinners(s ScratchCard) []int {
	winners := []int{}
	for _, v := range s.Numbers {
		if slices.Contains(s.WinningNumbers, v) {
			winners = append(winners, v)
		}
	}
	return winners
}

func Pow2(n int) int {
	if n > 0 {
		return 1 << (n - 1)
	}
	return 0
}
