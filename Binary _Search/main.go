package main

import "fmt"

// thuật toán tìm kiếm nhị phân để tìm ra giá trị trong 1 array
// tìm ra ví trị chính giữa của array , so sánh giá trị cần tìm với giá trị của phần tử chính giữa
// nếu nó thuộc nửa nhỏ hơn hay lớn hơn thì thực hiện tương tự với nửa tương ứng đến khi nào tìm ra được giá trị cần tìm

func binarySearch(array []int, key int) bool {
	startIndex := 0
	endIndex := len(array) - 1
	for startIndex <= endIndex {
		middle := (startIndex + endIndex) / 2
		if array[middle] < key {
			startIndex = middle + 1
		} else {
			endIndex = middle - 1
		}
	}
	if startIndex == len(array) || array[startIndex] != key {
		return false

	}
	return true

}

func main() {

	array := []int{1, 2, 9, 20, 31, 45, 63, 70}
	fmt.Println(binarySearch(array, 70))
}
