package problems

import (
  "fmt"
  "strings"
  "io/ioutil"
  "strconv"
  "reflect"
)

func parseOperation(s string) []int {
  arrSplitted := strings.Split(s, " ")
  target, _ := strconv.Atoi(arrSplitted[1])
  origin, _ := strconv.Atoi(arrSplitted[3])
  destin, _ := strconv.Atoi(arrSplitted[5])

  return []int{target, origin, destin}
}

func executeOperation([]int) {
  
}

func P5() {
  // Read input file
  data, err := ioutil.ReadFile("inputs/p5.txt")
  if err != nil {
    panic(err)
  }

  fileContents := string(data)

  arrOfInputs := []string{
    "FDBZTJRN",
    "RSNJH",
    "CRNJGZFQ",
    "FVNGRTQ",
    "LTQF",
    "QCWZBRGN",
    "FCLSNHM",
    "DNQMTJ",
    "PGS"}

  letters := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
  encoding := map[string]int{}

  operations_str := fileContents[strings.Index(fileContents, "move"):]
  operations_tmp := strings.Split(operations_str, "\n")

  fmt.Println(reflect.TypeOf(operations_tmp)) // []string

  operations := [][]int{}

  for i := 0; i < len(letters); i++ {
    encoding[letters[i]] = i + 1
  }

  for j := 0; j < len(operations_tmp); j++ {
    operations = append(operations, parseOperation(operations_tmp[j]))
    
  }

  

  fmt.Println(operations)
  fmt.Println(arrOfInputs)

  // fmt.Println("Problem 5a")
}
