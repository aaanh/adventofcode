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

func parseLineToGame(line string) Game {
	var g Game

	id := strings.Split(string(line), "\n")[0]
	id = strings.Split(string(id), " ")[1]
	id = strings.Split(string(id), ":")[0]
	g.Id, _ = strconv.Atoi(id)

	configs := strings.Split(string(strings.Split(string(line), ":")[1]), ";")

	// fmt.Println(configs)

	re_red := regexp.MustCompile(`(\d+.red)`)
	re_green := regexp.MustCompile(`(\d+.green)`)
	re_blue := regexp.MustCompile(`(\d+.blue)`)

	for _, config := range configs {
		var red string
		var green string
		var blue string
		for _, set := range strings.Split(config, ",") {
			// fmt.Println(set)
			red_match := re_red.FindAllString(set, -1)
			green_match := re_green.FindAllString(set, -1)
			blue_match := re_blue.FindAllString(set, -1)

			// fmt.Println(red_match, green_match, blue_match)

			if len(red_match) > 0 {
				red = strings.Split(red_match[0], " ")[0]
			}

			if len(green_match) > 0 {
				green = strings.Split(green_match[0], " ")[0]
			}

			if len(blue_match) > 0 {
				blue = strings.Split(blue_match[0], " ")[0]
			}

		}
		var c Config
		c.Red, _ = strconv.Atoi(red)
		c.Green, _ = strconv.Atoi(green)
		c.Blue, _ = strconv.Atoi(blue)
		g.Configs = append(g.Configs, c)
		// fmt.Println()
	}

	return g
}

func main() {

	input, _ := os.ReadFile("./day2.input")
	lines := strings.Split(string(input), "\n")
	id_sum := 0
	var games []Game

	for _, line := range lines {
		games = append(games, parseLineToGame(line))
	}

	for _, game := range games {
		isValid := true
		// fmt.Println(game.Id)
		for _, config := range game.Configs {
			// fmt.Println("red", config.Red, "green", config.Green, "blue", config.Blue)

			if !(config.Red <= RED_MAX && config.Green <= GREEN_MAX && config.Blue <= BLUE_MAX) {
				isValid = false
				break
			}
		}
		if isValid {
			id_sum += game.Id
		}
		// fmt.Println()
	}

	sum2 := 0
	for _, game := range games {
		minRed := 0
		minGreen := 0
		minBlue := 0
		power := 1
		for _, config := range game.Configs {
			if minRed <= config.Red {
				minRed = config.Red
			}

			if minGreen <= config.Green {
				minGreen = config.Green
			}

			if minBlue <= config.Blue {
				minBlue = config.Blue
			}
		}
		if minGreen == 0 {
			minGreen = 1
		}
		if minBlue == 0 {
			minBlue = 1
		}
		if minRed == 0 {
			minRed = 1
		}
		power = minRed * minBlue * minGreen
		fmt.Println("power:", power)
		sum2 += power
	}

	fmt.Println(id_sum)
	fmt.Println(sum2)
}
