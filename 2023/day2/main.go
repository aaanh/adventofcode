package main

import (
	"os"
	"strconv"
	"strings"
)

type game struct {
	Id      int
	Configs []config
}

type config struct {
	Green int
	Blue  int
	Red   int
}

func parseConfig(config string) {

}

func parseLineToGame(line string) game {
	var g game

	id := strings.Split(string(line), "\n")[0]
	id = strings.Split(string(id), " ")[1]

	g.Id, _ = strconv.Atoi(id)

	configs := strings.Split(string(strings.Split(string(line), ":")[1]), ";")

	return g

}

func main() {
	const RED_MAX = 12
	const GREEN_MAX = 13
	const BLUE_MAX = 14

	input, _ := os.ReadFile("./day2.input")
	lines := strings.Split(string(input), "\n")

}
