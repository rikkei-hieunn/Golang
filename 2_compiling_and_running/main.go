package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
}

// go run: thuc hien compile va chay chuong trinh ngay lap tuc
// Vi du: go run main.go

// go build: thuc hien compile cac file trong folder hien tai
// Truoc khi thuc hien compile, tao module: go mod init <ten_module>
// Vi du: go build -o main
// GOOS va GOARCH Thuc hien build voi config he dieu hanh va CPU
// Vi du: GOOS=windows GOARCH=amd64 go build -o winapp.exe (thuc hien build voi he dieu hanh window va CPU amd64)
// Vi du: GOOS=linux GOARCH=amd64 go build -o linuxapp (thuc hien build voi he dieu hanh linux va CPU amd64)
// Vi du: GOOS=darwin GOARCH=amd64 go build -o macapp (thuc hien build voi he dieu hanh mac va CPU amd64)

// go install: cung giong nhu go build, go install cung compile source code nhung file sau khi compile se duoc di chuyen
// den thu muc GOPATH/bin
