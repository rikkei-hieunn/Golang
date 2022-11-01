// Trong Golang tat ca source code deu duoc de trong package
// Package duoc thuc thi truoc tien la package main
package main

// Trong Golang cac package co the call den nhau de su dung cac function cua nhau
// Khi muon goi den mot package khac, ta su dung tu khoa `import`
// Vi du, import package fmt de su dung function in ra man hinh console Println
import "fmt"

// Trong package main luon co mot function la function main
// function main la noi khi chuong trinh duoc chay se nhay vao dau tien
func main() {
	// Goi function Println trong package fmt de in ra man hinh chuoi Hello World!
	fmt.Println("Hello World!")
}
