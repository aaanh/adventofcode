package problems

import (
  "aaanh/aoc/util"
  "fmt"
  "strings"
)

func Intersection(a, b []string) (c []string) {
  m := make(map[string]bool)

  for _, item := range a {
    m[item] = true
  }

  for _, item := range b {
    if _, ok := m[item]; ok {
      c = append(c, item)
    }
  }
  return
}

func P3() {
  var fileScanner = util.FileReader("inputs/p3.txt")
  // initialize an array of letters from A to Z including lowercase
  var letter = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

  priority_map := map[string]int{}

  var tmp_grp []string

  var total int

  for i := 0; i < len(letter); i++ {
    priority_map[letter[i]] = i + 1
  }

  line_idx := 0

  for fileScanner.Scan() {
    line := fileScanner.Text()
    if line_idx == 0 {
      tmp_grp = append(tmp_grp, line)
    } else if (line_idx+1)%3 == 0 {
      tmp_grp = append(tmp_grp, line)
      total += priority_map[Intersection(Intersection(strings.Split(tmp_grp[0], ""), strings.Split(tmp_grp[1], "")), strings.Split(tmp_grp[2], ""))[0]]

      tmp_grp = nil
    } else {
      tmp_grp = append(tmp_grp, line)
    }
    line_idx++
  }

  fmt.Println(total)
}
