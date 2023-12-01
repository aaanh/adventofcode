package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func day1_1() {
	input, _ := os.ReadFile("./day1.input")
	// fmt.Print(string(input))
	lines := strings.Split(string(input), "\n")
	sum := 0
	for _, line := range lines {
		tmpStr := ""

		for _, char := range line {
			if unicode.IsDigit(char) {
				tmpStr += string(char)
			}
		}

		// get first and last char and assign to tmpStr
		tmpStr = string(tmpStr[0]) + string(tmpStr[len(tmpStr)-1])

		converted, _ := strconv.Atoi(tmpStr)
		sum += converted
	}
	fmt.Println(sum)
}
