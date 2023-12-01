package problems

import (
  "bufio"
  "fmt"
  "os"
  "sort"
  "strconv"
)

func P1() {
  readFile, err := os.Open("inputs/p1.txt")

  var sums []int
  var max_three int
  tmp := 0

  if err != nil {
    fmt.Println(err)
  }

  fileScanner := bufio.NewScanner(readFile)
  fileScanner.Split(bufio.ScanLines)

  for fileScanner.Scan() {
    line := fileScanner.Text()
    if line != "" {
      num, err := strconv.Atoi(line)
      if err != nil {
        fmt.Println("Error conversion")
      } else {
        tmp = tmp + num
      }
    } else {
      sums = append(sums, tmp)
      tmp = 0
    }
    fmt.Println(fileScanner.Text())
  }

  sort.Ints(sums)

  for i := len(sums) - 1; i > len(sums)-4; i-- {
    max_three += sums[i]
    fmt.Println(max_three)
  }

  fmt.Println(max_three)

  readFile.Close()
}
