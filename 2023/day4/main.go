package main

import (
	"fmt"
	"os"
	// "strconv"
	"strings"
)

type Card struct {
	Id       int
	Winnings []int
	Draws    []int
}

func parseLineToCard(line string) Card {
	spli
}

func main() {
	input, _ := os.ReadFile("./input")
	lines := strings.Split(string(input), "\n")
	for _, line := range lines {
		fmt.Println(line)
	}
}
