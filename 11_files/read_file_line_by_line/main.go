package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open(".\\11_files\\read_file_line_by_line\\newFile.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Printf("Line: %s\n", scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
