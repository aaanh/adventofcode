package problems

import (
    "fmt"
    "io/ioutil"
    "strings"
    "strconv"
)

// parseAssignmentPair takes a string of the form "2-4,6-8"
// and returns the two assignments as integers in a slice
func parseAssignmentPair(s string) []int {
    parts := strings.Split(s, ",")
    a1 := strings.Split(parts[0], "-")
    a2 := strings.Split(parts[1], "-")
    a1Start, _ := strconv.Atoi(a1[0])
    a1End, _ := strconv.Atoi(a1[1])
    a2Start, _ := strconv.Atoi(a2[0])
    a2End, _ := strconv.Atoi(a2[1])
    return []int{a1Start, a1End, a2Start, a2End}
}

// countContainedAssignments takes a slice of assignment pairs
// and returns the number of pairs where one assignment fully
// contains the other
func countContainedAssignments(pairs []string) int {
    count := 0
    for _, pair := range pairs {
        assignments := parseAssignmentPair(pair)
        if assignments[0] <= assignments[2] && assignments[1] >= assignments[3] ||
            assignments[2] <= assignments[0] && assignments[3] >= assignments[1] {
            count++
        }
    }
    return count
}

// 1start, 1end, 2start, 2end
func countOverlappingAssignments(pairs []string) int {
    count := 0
    for _, pair := range pairs {
        assignments := parseAssignmentPair(pair)
        if (assignments[0] <= assignments[3] && assignments[1] >= assignments[2]) || (assignments[2] <= assignments[1] && assignments[3] >= assignments[0]) {
            count++
        }
    }
    return count
}

func P4() {
    // Read the file contents
    data, err := ioutil.ReadFile("inputs/p4.txt")
    if err != nil {
        panic(err)
    }
    // Convert the file contents to a string
    fileContents := string(data)
    // Split the file contents on newlines to get a slice of assignment pairs
    assignmentPairs := strings.Split(fileContents, "\n")
    // Count the number of pairs where one assignment fully contains the other
    count_4a := countContainedAssignments(assignmentPairs)
    // Count the number of pairs where one assignment overlaps the other
    count_4b := countOverlappingAssignments(assignmentPairs)

    fmt.Println("Problem 4a:")
    fmt.Println(count_4a)
    fmt.Println("Problem 4b:")
    fmt.Println(count_4b)
}