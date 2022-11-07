package main

// thuận toán insertion Sot là thuật toán sắp xếp chèn thuật toán này duyệt từng giá trị trong mảng và tìm ví trị phù hợp để chèn vào
// bản chất thì cũng giống như bubble sort
// độ phức tạp của thuật toàn là o(n2) vs n là số phần tử trong mảng

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	slice := generateSlice(20)
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	insertionSort(slice)
	fmt.Println("\n--- Sorted ---\n\n", slice, "\n")
}

// Tạo ra 1 slice ngẫu nhiên
func generateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

func insertionSort(items []int) {
	var n = len(items)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if items[j-1] > items[j] {
				items[j-1], items[j] = items[j], items[j-1]
			}
			j = j - 1
		}
	}
}
