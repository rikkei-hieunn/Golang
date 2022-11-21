package main

import (
	"fmt"
	"math/rand"
	"time"
)

// thuật toán Quick Sort còn được gọi là chia để trị
//Chia: chọn là phần từ chốt p (pivot) rồi chia làm 2 dãy con , dãy bên trái là những phần tử nhỏ hơn P và bên phải là những phần tử lớn hơn P
// Trị : lặp lại đệ quy thuật toán
// có nhiều cách lấy điểm P

func main() {

	slice := generateSlice(20)
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	quickSort(slice)
	fmt.Println("\n--- Sorted ---\n\n", slice, "\n")
}

// Generates a slice of size, size filled with random numbers
func generateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

//C1: chọn chốt ngẫu nhiên
func quickSort(slice []int) []int {
	if len(slice) < 2 {
		return slice
	}
	left, right := 0, len(slice)-1
	pivot := rand.Int() % len(slice)
	slice[pivot], slice[right] = slice[right], slice[pivot]

	for index, _ := range slice {
		if slice[index] < slice[right] {
			slice[left], slice[index] = slice[index], slice[left]
			left++
		}
	}
	slice[left], slice[right] = slice[right], slice[left]
	quickSort(slice[:left])
	quickSort(slice[left+1:])
	return slice

}
