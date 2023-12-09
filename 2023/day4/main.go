package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"

	// "strconv"
	"strings"
)

type Card struct {
	Id       int
	Winnings []int
	Rolls    []int
	Matches  int
	Points   int
}

func countMatches(arr1, arr2 []int) int {
	sort.Ints(arr1)
	sort.Ints(arr2)

	ptr1, ptr2 := 0, 0
	matches := 0

	for ptr1 < len(arr1) && ptr2 < len(arr2) {
		if arr1[ptr1] == arr2[ptr2] {
			matches++
			ptr1++
			ptr2++
		} else if arr1[ptr1] < arr2[ptr2] {
			ptr1++
		} else {
			ptr2++
		}
	}

	return matches
}

func parseLineToCard(line string) Card {
	var c Card
	separated := strings.Split(line, ":")
	card_header := strings.Split(separated[0], " ")
	var winnings []int
	var rolls []int
	for _, winning := range strings.Fields(strings.Split(strings.Split(line, ":")[1], "|")[0]) {
		winning = strings.ReplaceAll(winning, " ", "")
		winning_int, _ := strconv.Atoi(winning)
		winnings = append(winnings, winning_int)
	}

	for _, roll := range strings.Fields(strings.Split(strings.Split(line, ":")[1], "|")[1]) {
		roll = strings.ReplaceAll(roll, " ", "")
		roll_int, _ := strconv.Atoi(roll)
		rolls = append(rolls, roll_int)
	}

	c.Winnings = winnings
	c.Rolls = rolls

	c.Matches = countMatches(c.Winnings, c.Rolls)

	if c.Matches == 0 {
		c.Points = 0
	} else if c.Matches == 1 {
		c.Points = 1
	} else {
		c.Points = int(math.Pow(2, float64(c.Matches-1)))
	}

	c.Id, _ = strconv.Atoi(card_header[1])

	// fmt.Println(c)
	return c
}

func main() {
	input, _ := os.ReadFile("./input")
	lines := strings.Split(string(input), "\n")
	var cards []Card

	for _, line := range lines {
		cards = append(cards, parseLineToCard(line))
	}

	part1 := 0
	for _, card := range cards {
		part1 += card.Points
	}

	cardCopies := make(map[int]int)

	for i := 0; i < len(cards); i++ {
		card := cards[i]
		numOfMatches := card.Matches

		// Win copies of the next cards based on the matching numbers.
		for j := i + 1; j <= i+numOfMatches && j < len(cards); j++ {
			nextCard := &cards[j]
			cardCopies[nextCard.Id] += 1
			nextCard.Winnings = append(nextCard.Winnings, card.Id)
			nextCard.Rolls = append(nextCard.Rolls, i+1)
		}
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", len(cards))
}
