package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// Trong Golang hỗ trợ một kiểu dữ liệu để lưu trữ chuỗi văn bản gọi là string

func main() {
	// Khai bóa string
	var firstString string = "hello"
	fmt.Println("The value of firstString:", firstString)

	// Cú pháp khai báo ngắn string
	secondString := "Trung Le Quang"
	fmt.Println("The value of secondString:", secondString)

	// Để khai báo một chuỗi string, phần giá trị được để trong dấu "" hoặc ``
	// Khi giá trị được đặt trong dấu "" thì các ký tự định dạng văn bản như \n, \t sẽ có hiệu lực
	thirdString := "Trung Le \n Quang"
	fmt.Println("The value of thirdString:", thirdString)
	// Khi giá trị đặt trong dấu `` khi này giá trị được coi là raw string, các ký tự định dạng không có hiệu lực
	fourthString := `Trung Le \n Quang`
	fmt.Println("The value of fourthString:", fourthString)

	// Về mặt bản chất string chính là một Slice bytes, nên trong string hỗ trợ index
	// Giá trị lấy ra tại một index cụ thể là giá trị Decimal của ký tự đó trong bảng mã ASCII
	// Ví trị lấy giá trị tại index = 0 tức là ký tự 'T', theo như bảng mã ASCII thì ký tự 'T' có Decimal = 84
	fmt.Println("Element at index 0 of fourthString:", fourthString[0])

	// String trong Golang là bất biến(immutable) nên không thể nào thay đổi giá trị tại một index cụ thể trong string
	// fourthString[2] = 'Y' -> return error

	// Trong Golang không có kiểu dữ liệu Character nên để biểu thị được giá trị của các ký tự ta sử dụng kiểu dữ liệu Byte và Rune
	// Các ký tự Rune hoặc Byte được đặt trong dấu nháy đơn ''
	var firstByte byte = 'A'
	var firstRune rune = ''
	fmt.Println("The value of firstByte:", firstByte)
	fmt.Println("The value of firstRune:", firstRune)

	fifthString := "Lê Quáng Trúng"
	sixthString := "Le Quang Trung"
	for i := 0; i < len(fifthString); i++ {
		fmt.Printf("%c", fifthString[i])
	}
	fmt.Printf("\n")
	fmt.Println("Lengh of string fifthString", "=", len(fifthString))
	fmt.Println("Lengh of string sixthString", "=", len(sixthString))
	// Nhận thấy length của 2 string fifthString và sixthString là hoàn toàn khác nhau, do các ký tự UTF-8 bình thường
	// sẽ tốn 1 byte số ký tự trong string sixthString là 14 nên length của sixthString là 14 bytes
	// Đối với string fifthString bao gồm các ký tự không phải là UTF-8 như ê, á, ú nên số lượng bytes của các ký tự này
	// không phải là 1
	fmt.Printf("Number of bytes ê %d\n", len("ê"))
	fmt.Printf("Number of bytes á %d\n", len("á"))
	fmt.Printf("Number of bytes ú %d\n", len("ú"))
	// Về mặt bản chất string là một slice bytes nên nếu dùng for duyệt từng ký tự trong string chỉ đúng trong trường
	// hợp string đó bao gồm toàn bộ là các ký tự UTF-8
	// Trong trường hợp string bao gồm cả các ký tự unicode như fifthString phải áp dụng với Rune
	for i := 0; i < len(fifthString); {
		sampleRune, size := utf8.DecodeRuneInString(fifthString[i:])
		fmt.Printf("%c", sampleRune)
		i += size
	}
	fmt.Printf("\n")

	// Do string là slice nên hoàn toàn có thể sử dụng index để lấy ra một string con trong string cha
	fmt.Println("The value from index 0 to 5 of sixthString:", sixthString[0:5])
	// Đối với những string có chưa unicode để có thể lấy substring theo index thì cần convert từ slice bytes sang slice Rune
	fifthStringRune := []rune(fifthString)
	fmt.Println("The value from index 0 to 5 of fifthString:", string(fifthStringRune[0:5]))

	// Kiểm tra một substring có chứa trong một string khác hay không
	isContain := strings.Contains("Hello Golang", "Golang")
	fmt.Println("isContain:", isContain)

	// Kiểm tra xem string có chứ 1 trong các ký tự lex hay không
	isContainAny := strings.ContainsAny("Le Quang Trung", "lex")
	fmt.Println("isContainAny:", isContainAny)

	// Kiểm tra xem string có chưa Rune ê hay không
	isContainRune := strings.ContainsRune("Lê Quang Trung", 'ê')
	fmt.Println("isContainRune:", isContainRune)

	// Đếm xem trong string có bao nhiêu chữ u
	uNumber := strings.Count("Le Quang Trung", "u")
	fmt.Println("uNumber:", uNumber)

	// So sánh 2 string
	str1 := "Golang"
	str2 := "golang"
	isEqual := str1 == str2
	fmt.Println("isEqual:", isEqual)

	// So sánh 2 string bỏ qua chữ viết hoa viết thường
	isEqualForce := strings.EqualFold(str1, str2)
	fmt.Println("isEqualForce:", isEqualForce)

	// Toàn bộ các functions trong package strings: https://pkg.go.dev/strings
}
