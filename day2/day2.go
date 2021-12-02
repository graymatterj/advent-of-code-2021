package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func calculatePosition(scanner *bufio.Scanner) (horizontalPosition, depth, multiplication int) {
	for scanner.Scan() {
		inputString := scanner.Text()
		fields := strings.Fields(inputString)
		direction := fields[0]
		quantity, _ := strconv.Atoi(fields[1])
		if direction == "up" {
			depth -= quantity
		} else if direction == "down" {
			depth += quantity
		} else if direction == "forward" {
			horizontalPosition += quantity
		}
	}
	multiplication = horizontalPosition * depth
	return
}

func main() {
	file, err := os.Open("input-directions.txt")
    if err != nil {
        log.Fatal(err)
    }
	
	scanner := bufio.NewScanner(file)
	fmt.Println(calculatePosition(scanner))

	file.Close()
}