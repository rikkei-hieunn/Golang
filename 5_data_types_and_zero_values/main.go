package main

import "fmt"

// Trong Golang bao gồm các loại kiểu dữ liệu lớn như sau
// 1. Basic type: Numbers, strings and booleans
// 2. Aggregate type: Array and structs
// 3. Reference type: Pointers, slices, maps, functions and channels
// 4. Interface type

// Trong phần này mình chỉ tập trung vào phần Basic type

func main() {
	// Kiểu dữ liệu số
	// int8: kiểu số nguyên 8 bit từ -128 đến 127
	var number int8 = 100
	fmt.Println("Number:", number)

	// int16: kiểu số nguyên 16 bit từ -32768 đến 32767
	var secondNumber int16 = 32767
	fmt.Println("Second Number:", secondNumber)

	// int32: kiểu số nguyên 32 bit từ -2147483648 đến 2147483647
	var thirdNumber int32 = 2147483647
	fmt.Println("Third Number:", thirdNumber)

	// int32: kiểu số nguyên 64 bit từ -9223372036854775808 đến 9223372036854775807
	var fourthNumber int64 = -9223372036854775808
	fmt.Println("Fourth Number:", fourthNumber)

	// uint8: kiểu số nguyên dương 8 bit từ 0 đến 255
	var fifthNumber uint8 = 0
	fmt.Println("Fifth Number:", fifthNumber)

	// uint16: kiểu số nguyên dương 16 bit từ 0 đến 65535
	var sixthNumber uint16 = 65535
	fmt.Println("Sixth Number:", sixthNumber)

	// uint32: kiểu số nguyên dương 32 bit từ 0 đến 4294967295
	var seventhNumber uint32 = 4294967295
	fmt.Println("Seventh Number:", seventhNumber)

	// uint64: kiểu số nguyên dương 64 bit từ 0 đến 18446744073709551615
	var eighthNumber uint64 = 18446744073709551615
	fmt.Println("Eighth Number:", eighthNumber)

	// float32: kiểu số thập phân 32 bit
	var ninthNumber float32 = 123456.123456
	fmt.Println("Ninth Number:", ninthNumber)

	// float64: kiểu số thập phân 64 bit
	var tenthNumber float64 = 123456789.123456
	fmt.Println("Tenth Number:", tenthNumber)

	// complex64: kiểu số với hai phần thật(real) và phần ảo(imaginary) 64 bit
	var eleventhNumber complex64 = complex(100, 2)
	fmt.Println("Eleventh Number:", eleventhNumber)

	// complex128: kiểu số với hai phần thật(real) và phần ảo(imaginary) 128 bit
	var twelfthNumber complex128 = complex(2000, 2)
	fmt.Println("Twelfth Number:", twelfthNumber)

	// rune: tương đương với int32 nó là số đại diện cho mã Unicode
	var demoRune rune = 10
	var demoRuneUnicode = ''
	fmt.Println("Demo Rune:", demoRune)
	fmt.Println("Demo Rune Unicode:", demoRuneUnicode)

	// byte: tương đương với uint8
	var demobyte byte = 127
	fmt.Println("Demo byte:", demobyte)

	//-----------------------------------------------------------------------------------------------------------------
	// Kiểu dữ liệu chuỗi: string
	// Kiểu dữ liệu chuỗi biểu thị một chuỗi các điểm mã Unicode. Hay nói cách khác, chúng là một chuỗi các byte bất biến
	var name string = "Le Quang Trung"
	fmt.Println("Name:", name)

	//-----------------------------------------------------------------------------------------------------------------
	// Kiểu dữ liệu logic: bool
	// Kiểu dữ liệu boolean chỉ đại diện cho một bit thông tin đúng hoặc sai
	var isRunning bool = true
	fmt.Println("Is Running:", isRunning)

	//-----------------------------------------------------------------------------------------------------------------
	// Giá trị mặc định của kiểu dữ liệu(Zero values)
	var numberInt int
	fmt.Println("Giá trị mặc định của kiểu dữ liệu int:", numberInt)
	var numberInt8 int8
	fmt.Println("Giá trị mặc định của kiểu dữ liệu int8:", numberInt8)
	var numberInt16 int16
	fmt.Println("Giá trị mặc định của kiểu dữ liệu int16:", numberInt16)
	var numberInt32 int32
	fmt.Println("Giá trị mặc định của kiểu dữ liệu int32:", numberInt32)
	var numberInt64 int64
	fmt.Println("Giá trị mặc định của kiểu dữ liệu int64:", numberInt64)

}
