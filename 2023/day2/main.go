package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const RED_MAX = 12
const GREEN_MAX = 13
const BLUE_MAX = 14

type Game struct {
	Id      int
	Configs []Config
}

type Config struct {
	Green int
	Blue  int
	Red   int
}

func isValidGame(g Game) bool {
	return true
}

func parseConfig(config string) {

}

func parseLineToGame(line string) Game {
	var g Game

	id := strings.Split(string(line), "\n")[0]
	id = strings.Split(string(id), " ")[1]

	g.Id, _ = strconv.Atoi(id)

	configs := strings.Split(string(strings.Split(string(line), ":")[1]), ";")

	re := regexp.MustCompile(`\d+|red|green|blue`)

	for _, config := range configs {
		fmt.Println(config)
		for _, set := range strings.Split(config, ",") {
			matches := re.FindAllString(set, -1)
			_ = matches
		}
	}

	fmt.Println()

	return g
}

func main() {

	input, _ := os.ReadFile("./day2.input")
	lines := strings.Split(string(input), "\n")
	for _, line := range lines {
		parseLineToGame(line)
	}

}
