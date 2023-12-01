package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	// "unicode"
)

// numberWords maps spelled-out numbers to their integer values.
var numberWords = map[string]int{
	"zero": 0, "one": 1, "two": 2, "three": 3, "four": 4,
	"five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9,
}

func parseLine(line string) string {
	wordToNumber := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	// Regular expression to identify numbers and words representing numbers
	re := regexp.MustCompile(`\d+|one|two|three|four|five|six|seven|eight|nine`)
	matches := re.FindAllString(line, -1)
	fmt.Println(matches)

	var result strings.Builder
	for _, match := range matches {
		if val, exists := wordToNumber[match]; exists {
			// Append the numeric value of the word
			result.WriteString(val)
		} else {
			// Append the digit as it is
			result.WriteString(match)
		}
	}

	return result.String()
}

func day1_2() {
	input, _ := os.ReadFile("./example.input")
	lines := strings.Split(string(input), "\n")
	sum := 0

	for idx, line := range lines {
		parsed := parseLine(line)
		runes := []rune(parsed)
		first := string(runes[0])
		last := string(runes[len(runes)-1])
		num, _ := strconv.Atoi(first + last)
		fmt.Println(idx+1, ">", num)
		sum += num
	}

	fmt.Println(sum)
}
