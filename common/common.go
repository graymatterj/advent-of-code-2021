package common

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func RunFunctionAgaisntFile(functionToExecute func(*bufio.Scanner) string, filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	fmt.Println(functionToExecute(scanner))

	file.Close()
}
