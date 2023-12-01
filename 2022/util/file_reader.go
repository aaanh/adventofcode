package util

import (
  "bufio"
  "fmt"
  "os"
)

func FileReader(filePath string) *bufio.Scanner {
  readFile, err := os.Open(filePath)

  if err != nil {
    fmt.Println(err)
  }

  fileScanner := bufio.NewScanner(readFile)
  fileScanner.Split(bufio.ScanLines)

  return fileScanner
}
