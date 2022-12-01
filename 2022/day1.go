package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	f, err := os.Open("./day1_1.txt")
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	defer f.Close()

	elves := []int{}
	scanner := bufio.NewScanner(f)
	z := 0
	for scanner.Scan() {
		if scanner.Text() == "" {
			elves = append(elves, z)
			z = 0
			continue
		}
		i, err := strconv.ParseInt(scanner.Text(), 10, 0)
		if err != nil {
			fmt.Printf("Error: %v", err)
		}
		z += int(i)
	}
	sort.Slice(elves, func(i, j int) bool {
		return elves[i] < elves[j]
	})
	fmt.Println(elves[len(elves)-1])
	result := 0
	for _, v := range elves[len(elves)-3:] {
		result += v
	}
	fmt.Println(result)
}
