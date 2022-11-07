package main

import "fmt"

// để tìm kiếm 1 phần từ trong array or slice thì t duyệt từng phần tử trong array đó là so sánh với phầm từ cần tìm kiếm

// giả sử chúng ta cần tìm số 99 trong array

func LinerSearch(array []int, key int) bool {

	for _, value := range array {
		if value == key {
			return true
		}
	}
	return false
}

func main() {
	array := []int{95, 78, 46, 58, 45, 86, 99, 251, 320}
	fmt.Println(LinerSearch(array, 99))
}
