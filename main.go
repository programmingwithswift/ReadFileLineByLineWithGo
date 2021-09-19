package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("fileToReadLargeLine.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	bufferSize := 1024 * 1024
	scannerBuffer := make([]byte, bufferSize)
	scanner.Buffer(scannerBuffer, bufferSize)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if scanner.Err() == bufio.ErrTooLong {
		log.Fatal(scanner.Err())
	}
}
