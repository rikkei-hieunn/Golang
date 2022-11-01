package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Tạo một file mới
	newFile, err := os.Create("./11_files/write_file_using_buffered/newFile.txt")
	if err != nil {
		panic(err)
	}
	_ = newFile.Close()

	// Mở file vừa tạo
	file, err := os.OpenFile("./11_files/write_file_using_buffered/newFile.txt",
		os.O_WRONLY|os.O_CREATE,
		0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	bufferedWriter := bufio.NewWriter(file)
	bytesSlice := []byte{97, 98, 99}
	bytesWritten, err := bufferedWriter.Write(bytesSlice)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Bytes written to buffer(not file) %d\n", bytesWritten)
	bytesAvailabale := bufferedWriter.Available()
	fmt.Printf("Bytes available in buffer %d\n", bytesAvailabale)

	// Ghi vào file
	_ = bufferedWriter.Flush()
	bufferedWriter.Reset(bufferedWriter)
}
