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

func countIncreasingWindowedDepths(scanner *bufio.Scanner) int {
	var depthCache[] int;
	currentSum:=0
	increaseDepthCount:=0
	for scanner.Scan() {
		currentLine:=scanner.Text()
		currentNumber, _:=strconv.Atoi(currentLine)
		depthCache = append(depthCache, currentNumber)
		if len(depthCache) > 3 {
			lastSum := currentSum
			currentSum = currentSum - depthCache[0] + currentNumber
			depthCache = depthCache[1:];
			if currentSum > lastSum {
				increaseDepthCount++
			}
		} else {
			currentSum = currentSum + currentNumber
		}
    }
	return increaseDepthCount
}

func main() {
	file, err := os.Open("input-depths.txt")
    if err != nil {
        log.Fatal(err)
    }
	
	scanner := bufio.NewScanner(file)
	// fmt.Println(countIncreasingDepths(scanner))
	fmt.Println(countIncreasingWindowedDepths(scanner))

	file.Close()
}
