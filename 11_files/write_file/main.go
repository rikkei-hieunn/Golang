package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// Tạo một file mới
	newFile, err := os.Create("./11_files/write_file/newFile.txt")
	if err != nil {
		panic(err)
	}
	_ = newFile.Close()

	// Mở file vừa tạo
	file, err := os.OpenFile("./11_files/write_file/newFile.txt",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	bytesSlice := []byte("Le Quang Trung")
	byteWritten, err := file.Write(bytesSlice)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Bytes Written %d\n", byteWritten)

	secondBytesSlice := []byte("Le Quang Trung12345")
	err = ioutil.WriteFile("./11_files/write_file/newFile.txt", secondBytesSlice, 0644)
	if err != nil {
		panic(err)
	}
}
