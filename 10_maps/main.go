package main

import "fmt"

// Map cũng là một kiểu tập hợp như Array hay Slice, lưu trữ dữ liệu dạng key:value
// Tất cả các key trong Map phải duy nhất
// Tất cả các key và value phải phải cùng kiểu dữ liệu với nhau
// Map hỗ trợ phương thức truy xuất nhanh chóng thông qua key

func main() {
	// Khai báo một Map
	var firstMap map[string]string
	fmt.Printf("The default value of map %#v\n", firstMap)
	fmt.Printf("Length of map %d\n", len(firstMap))

	// Trước khi sử dụng Map cần phải khởi tạo Map, sử dụng function make() hoặc khởi tạo giá trị luôn cho Map
	firstMap = make(map[string]string)
	firstMap["Trung"] = "Quang Le"
	firstMap["Hieu"] = "Quang Le"
	fmt.Println("The value of firstMap:", firstMap)

	secondMap := map[string]int{}
	secondMap["Trung"] = 26
	fmt.Println("The value of secondMap:", secondMap)

	thirdMap := map[int]string{
		2: "hello",
	}
	fmt.Println("The value of thirdMap:", thirdMap)

	// Kiểm tra key có trong Map hay không
	value, isExists := secondMap["Trung"]
	fmt.Println("isExists:", isExists, "value:", value)

	// Trong trường hợp key không tồn tại thì value sẽ trả về giá trị mặc định của kiểu dữ liệu value
	value, isExists = secondMap["Quang"]
	fmt.Println("isExists:", isExists, "value:", value)

	// Xóa một cặp key value trong Map
	delete(secondMap, "Trung")
	value, isExists = secondMap["Trung"]
	fmt.Println("isExists:", isExists, "value:", value)

}
