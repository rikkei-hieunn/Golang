package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// Mở file
	file, err := os.Open("./11_files/read_file_using_buffered/newFile.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	bytesSlice := make([]byte, 10)

	numberBytesRead, err := io.ReadFull(file, bytesSlice)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Number bytes read: %d\n", numberBytesRead)
	fmt.Printf("Data read: %s\n", bytesSlice)

	// Đọc toàn bộ file
	data, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Data read all: %s\n", data)
}
