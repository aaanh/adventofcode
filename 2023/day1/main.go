package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
	// "unicode"
)

func part1() {
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
		"1":     "1",
		"2":     "2",
		"3":     "3",
		"4":     "4",
		"5":     "5",
		"6":     "6",
		"7":     "7",
		"8":     "8",
		"9":     "9",
	}

	// Regular expression to identify numbers and words representing numbers
	re := regexp.MustCompile(`\d|one|two|three|four|five|six|seven|eight|nine`)
	var result strings.Builder

	for i := 0; i < len(line); i++ {
		matches := re.FindAllString(line[i:], -1)
		fmt.Println(matches)
		for _, match := range matches {
			if val, exists := wordToNumber[match]; exists {
				result.WriteString(val)
			}
		}
	}

	return result.String()
}

func part2() {

	input, _ := os.ReadFile("./day1.input")
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

func main() {
	part1()
	part2()
}
