package main

import (
	"fmt"
	"os"
)

func main() {
	// Tạo một file mới
	var newFile *os.File
	var err error
	newFile, err = os.Create("./11_files/basic_file/newFile.txt")
	if err != nil {
		panic(err)
	}
	_ = newFile.Close()

	// Mở một file
	file, err := os.OpenFile("./11_files/basic_file/newFile.txt", os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		// Kiểm tra lỗi có phải lỗi file ko tồn tại hay ko?
		if os.IsNotExist(err) {
			fmt.Println("File does not esist!")
		}
		panic(err)
	}
	_ = file.Close()

	// Lấy toàn bộ thông tin của file
	var fileInfo os.FileInfo
	fileInfo, err = os.Stat("./11_files/basic_file/newFile.txt")
	fmt.Println("File name:", fileInfo.Name())
	fmt.Println("Size in bytes:", fileInfo.Size())
	fmt.Println("Last modified:", fileInfo.ModTime())
	fmt.Println("Is directory?:", fileInfo.IsDir())
	fmt.Println("Permission:", fileInfo.Mode())

	// Đổi tên file
	oldPath := "./11_files/basic_file/newFile.txt"
	newPath := "./11_files/basic_file/newNewFile.txt"
	err = os.Rename(oldPath, newPath)
	if err != nil {
		panic(err)
	}

	// Xóa file
	err = os.Remove(newPath)
	if err != nil {
		panic(err)
	}
}
