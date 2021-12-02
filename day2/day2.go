package main

import (
    "advent-of-code/common"
	"bufio"
    "fmt"
	"strconv"
	"strings"
)

const OUTPUT_STRING = "Horizontal position: %d; Depth: %d; Multiplication: %d"
const UP, DOWN, FORWARD = "up","down","forward"

func calculatePosition(scanner *bufio.Scanner) string {
	var horizontalPosition, depth int
	for scanner.Scan() {
		inputString := scanner.Text()
		fields := strings.Fields(inputString)
		direction := fields[0]
		quantity, _ := strconv.Atoi(fields[1])
		if direction == UP {
			depth -= quantity
		} else if direction == DOWN {
			depth += quantity
		} else if direction == FORWARD {
			horizontalPosition += quantity
		}
	}
    return fmt.Sprintf(OUTPUT_STRING, horizontalPosition, depth, horizontalPosition * depth )
}

func calculatePositionWithAim(scanner *bufio.Scanner) string {
	var horizontalPosition, depth, aim int
	for scanner.Scan() {
		inputString := scanner.Text()
		fields := strings.Fields(inputString)
		direction := fields[0]
		quantity, _ := strconv.Atoi(fields[1])
		if direction == UP {
			aim -= quantity
		} else if direction == DOWN {
			aim += quantity
		} else if direction == FORWARD {
			horizontalPosition += quantity
			depth += aim * quantity
		}
	}
    return fmt.Sprintf(OUTPUT_STRING, horizontalPosition, depth, horizontalPosition * depth )
}

func main() {
	inputFile := "input-directions.txt"
	common.RunFunctionAgaisntFile(calculatePosition, inputFile)
	common.RunFunctionAgaisntFile(calculatePositionWithAim, inputFile)
}