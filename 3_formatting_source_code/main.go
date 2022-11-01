package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
}

// Trong Golang ho tro mot tool laf gofmt(Golang Formatter) giups format mai code cho dep hon
// Vi du: gofmt main.go (show ra toan bo code trong file main.go sau khi duoc format code)
// Vi du: gofmt -w main.go (tu dong ghi lai file main.go sau khi duoc format code)
// Vi du: gofmt -w -l main.go (tu dong ghi lai file main.go va show ra danh sach file duoc format code)
// Vi du: go fmt hoac gofmt -w . (format toan bo source code trong folder hien tai)
