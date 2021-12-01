package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
)

func countIncreasingDepths(scanner *bufio.Scanner) int {
    previousLine := 0
	increaseDepthCount := -1 // Avoid counting the first depth as greater than 0
	for scanner.Scan() {
		currentNumber, _ := strconv.Atoi(scanner.Text())
		if currentNumber > previousLine {
			increaseDepthCount++
		}
		previousLine = currentNumber
    }

	return increaseDepthCount
}

func main() {
	file, err := os.Open("input-depths.txt")
    if err != nil {
        log.Fatal(err)
    }
	
	scanner := bufio.NewScanner(file)
	fmt.Println(countIncreasingDepths(scanner))

	file.Close()
}
