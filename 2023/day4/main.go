package main

import (
	"fmt"
	"os"
	"strconv"

	// "strconv"
	"strings"
)

type Card struct {
	Id       int
	Winnings []int
	Draws    []int
}

func parseLineToCard(line string) Card {
	var c Card
	separated := strings.Split(line, ":")
	card_header := strings.Split(separated[0], " ")
	var winnings []string
	for _, winning := range strings.Split(separated[1], "|") {
		winning = strings.Replace(winning, " ", "", -1)
		winnings = append(winnings, winning)
		fmt.Println("Winnings:", winnings)
	}

	c.Id, _ = strconv.Atoi(card_header[1])
	return c
}

func main() {
	input, _ := os.ReadFile("./input")
	lines := strings.Split(string(input), "\n")
	for _, line := range lines {
		fmt.Println(parseLineToCard(line))
	}
}
