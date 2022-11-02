package main

import (
	"fmt"
	"strings"
)

// Dưới đây là những function string hay được sử dụng

func main() {
	message1 := "Hieu dz vcl"
	message2 := "<3"
	// use len() function to count length
	stringLength := len(message1)
	fmt.Println(stringLength)
	//Join Two Strings Together
	// concatenation using + operator
	result := message1 + " " + message2
	fmt.Println(result)
	// dưới đây là một số function được sử dụng để làm việc với chuỗi

	//	Compare()	compares two strings
	// create three strings
	string1 := "Programiz"
	string2 := "Programiz Pro"
	string3 := "Programiz"

	// compare strings
	fmt.Println(strings.Compare(string1, string2)) // -1
	fmt.Println(strings.Compare(string2, string3)) // 1
	fmt.Println(strings.Compare(string1, string3)) // 0

	//	Contains()	checks if a substring is present inside a string  //check xem chuỗi kĩ tự có trong chuỗi cần tìm không?
	text := "Go Programming"
	substring1 := "Go"
	substring2 := "Golang"

	// check if Go is present in Go Programming
	resultContains := strings.Contains(text, substring1)
	fmt.Println(resultContains)

	// check if Golang is present in Go Programming
	resultContains = strings.Contains(text, substring2)
	fmt.Println(resultContains)

	//	Replaces()	replaces a substring with another substring //thay đổi kí tự trong chuỗi
	textReplaces := "car"
	fmt.Println("Old String:", textReplaces) //car

	// replace r with t
	replacedText := strings.Replace(textReplaces, "r", "t", 1) //cat
	fmt.Println("New String:", replacedText)

	//	ToLower()	converts a string to lowercase // chuyển từ viết hoa thành viết thường
	//	ToUpper()	converts a string to uppercase // chuyển viết thường thành viết hoa
	text1 := "go is fun to learn"
	// convert to uppercase
	text1 = strings.ToUpper(text1)
	fmt.Println(text1)

	text2 := "I LOVE GOLANG"
	// convert to lowercase
	text2 = strings.ToLower(text2)
	fmt.Println(text2)

	//	Split()		splits a string into multiple substrings  // chia chuổi thành nhiều kí tự cách nhau
	var message = "I Love Golang"
	// split string from space " "
	splittedString := strings.Split(message, " ")
	fmt.Println(splittedString)
}
