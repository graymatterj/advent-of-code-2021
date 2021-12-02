package main

import (
	"advent-of-code/common"
	"bufio"
	"fmt"
	"strconv"
)

const OUTPUT_STRING = "Number of increasing depths: %d"

func countIncreasingDepths(scanner *bufio.Scanner) string {
	previousLine := 0
	increaseDepthCount := -1 // Avoid counting the first depth as greater than 0
	for scanner.Scan() {
		currentNumber, _ := strconv.Atoi(scanner.Text())
		if currentNumber > previousLine {
			increaseDepthCount++
		}
		previousLine = currentNumber
	}
	return fmt.Sprintf(OUTPUT_STRING, increaseDepthCount)
}

func countIncreasingWindowedDepths(scanner *bufio.Scanner) string {
	var depthCache []int
	currentSum := 0
	increaseDepthCount := 0
	for scanner.Scan() {
		currentLine := scanner.Text()
		currentNumber, _ := strconv.Atoi(currentLine)
		depthCache = append(depthCache, currentNumber)
		if len(depthCache) > 3 {
			lastSum := currentSum
			currentSum = currentSum - depthCache[0] + currentNumber
			depthCache = depthCache[1:]
			if currentSum > lastSum {
				increaseDepthCount++
			}
		} else {
			currentSum = currentSum + currentNumber
		}
	}
	return fmt.Sprintf(OUTPUT_STRING, increaseDepthCount)
}

func main() {
	inputFile := "input-depths.txt"
	common.RunFunctionAgaisntFile(countIncreasingDepths, inputFile)
	common.RunFunctionAgaisntFile(countIncreasingWindowedDepths, inputFile)
}
