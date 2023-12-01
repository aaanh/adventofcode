package problems

import (
	"aaanh/aoc/util"
	"fmt"
	"strings"
)

func P2a() {
	var fileScanner = util.FileReader("inputs/p2.txt")

	score_me := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()
		line_arr := strings.Split(line, " ")

		switch line_arr[0] {
		case "A":
			// TOOL + RESULT
			if line_arr[1] == "X" {
				score_me += (1 + 3)
			} else if line_arr[1] == "Y" {
				score_me += (2 + 6)
			} else if line_arr[1] == "Z" {
				score_me += (3 + 0)
			}
		case "B":
			if line_arr[1] == "X" {
				score_me += (1 + 0)
			} else if line_arr[1] == "Y" {
				score_me += (2 + 3)
			} else if line_arr[1] == "Z" {
				score_me += (3 + 6)
			}
		case "C":
			if line_arr[1] == "X" {
				score_me += (1 + 6)
			} else if line_arr[1] == "Y" {
				score_me += (2 + 0)
			} else if line_arr[1] == "Z" {
				score_me += (3 + 3)
			}
		}

	}

	fmt.Println(score_me)
}
