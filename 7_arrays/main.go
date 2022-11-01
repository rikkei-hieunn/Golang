package main

import "fmt"

// Một Array là một tập hợp, được đánh chỉ mục(index) lưu trữ tập hợp các phần tử có cùng kiểu dữ liệu với nhau
// Một Array luôn được cố định kích thức
// Toàn bộ các phần tử trong một Array phải cùng kiểu dữ liệu với nhau
// Trong Golang các phần tử trong cùng một Array có địa chỉ ô nhớ kế tiếp nhau
// Kích thước của một Array được xác định tại thời điểm biên dịch(Compile time) và không thể thay đổi

func main() {
	// Khai báo một Array integer gồm 4 phần tử
	var numbers [4]int

	// Giá trị mặc định(zero value) của Array numbers
	fmt.Printf("Default value of numbers: %#v\n", numbers)

	// Cú pháp khai báo ngắn một Array
	strs := [2]string{"hello", "world"}
	fmt.Printf("The value of strs: %#v\n", strs)

	// Khi khai báo một mảng, ta có thể không cần gán một giá trị cụ thể cho kích thước mảng
	// Mà kích thước mảng sẽ được tính toán dựa vào phần value khi khơi tạo mảng
	// => để làm được điều này, sử dụng toán tử Ellipsis(Ellipsis operator)
	listNumber := [...]float32{1.2, 2.3, 3.4, 4.5}
	fmt.Printf("The value of listNumber: %#v\n", listNumber)
	// Để lấy được số lượng phần tử trong mảng, sử dụng phương thức len() của Array
	fmt.Printf("The length of listNumber: %d\n", len(listNumber))

	// Một số toán tử sử dụng với Array
	// Toán tử gán nhằm thay đổi giá trị của một phần tử trong mảng
	fmt.Printf("The first element of listNumber before: %.1f\n", listNumber[0])
	listNumber[0] = 10
	fmt.Printf("The first element of listNumber after: %.1f\n", listNumber[0])

	// Duyệt toàn bộ phần tử trong mảng, sử dụng range
	for index, value := range listNumber {
		fmt.Printf("The value of index %d is %.1f\n", index, value)
	}

	// Duyệt toàn bộ phần tử trong mảng, không sử dụng range
	for index := 0; index < len(listNumber); index++ {
		fmt.Printf("The value of index %d is %.1f\n", index, listNumber[index])
	}

	// Mảng 2 chiều(matrix)
	matrix := [3][3]int{
		{1, 2, 3},
		{2, 3, 4},
		{4, 5, 6},
	}

	fmt.Printf("The value of matrix %#v\n", matrix)
	for rowIndex := 0; rowIndex < len(matrix); rowIndex++ {
		for colIndex := 0; colIndex < len(matrix[0]); colIndex++ {
			fmt.Printf("\t%d", matrix[rowIndex][colIndex])
		}
		fmt.Printf("\n")
	}

	// So sánh 2 Array
	firstArray := [4]int{1, 2, 3, 4}
	secondArray := [4]int{1, 2, 3, 4}
	fmt.Println("firstArray is equal to secondArray:", firstArray == secondArray)
	firstArray[0] = 0
	fmt.Println("firstArray is equal to secondArray:", firstArray == secondArray)

	// Khởi tạo mảng với key là index
	thirdArray := [4]int{
		1: 10,
		0: 40,
		3: 99,
		2: 9,
	}
	fmt.Printf("The value of third array: %#v\n", thirdArray)

	// Các index không được gán giá trị sẽ có giá trị mặc định
	fouthArray := [3]int{
		2: 19,
	}
	fmt.Printf("The value of fouth array: %#v\n", fouthArray)

	// Trong trường hợp khai báo mảng sử dụng cả 2 cách gán index và ko gán index
	// Khi này kích thước của mảng = max index + số lượng phần tử không được gán index
	fifthArray := [...]int{
		2:  19,
		20, // index 3
	}
	fmt.Printf("The value of fifth array: %#v\n", fifthArray)
}
