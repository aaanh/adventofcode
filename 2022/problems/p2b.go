package problems

import (
  "aaanh/aoc/util"
  "fmt"
  "strings"
)

func P2b() {
  var fileScanner = util.FileReader("inputs/p2.txt")

  score_me := 0

  for fileScanner.Scan() {
    line := fileScanner.Text()
    line_arr := strings.Split(line, " ")

    switch line_arr[0] {
    case "A": // rock
      // TOOL + RESULT
      if line_arr[1] == "X" {
        score_me += 0 + 3
      } else if line_arr[1] == "Y" {
        score_me += 3 + 1
      } else if line_arr[1] == "Z" {
        score_me += 6 + 2
      }
    case "B": // paper
      if line_arr[1] == "X" {
        score_me += 0 + 1
      } else if line_arr[1] == "Y" {
        score_me += 3 + 2
      } else if line_arr[1] == "Z" {
        score_me += 6 + 3
      }
    case "C": // scissors
      if line_arr[1] == "X" {
        score_me += 0 + 2
      } else if line_arr[1] == "Y" {
        score_me += 3 + 3
      } else if line_arr[1] == "Z" {
        score_me += 6 + 1
      }
    }

  }

  fmt.Println(score_me)
}
