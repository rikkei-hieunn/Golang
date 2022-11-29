package main

import "fmt"

//Interface trong Golang là một kiểu được định nghĩa bởi tập hợp của các method (hàm trong Golang)..
// Interface có thể chứa bất kỳ giá trị gì miễn là nó có implement các method này
//type Name interface {
//    Name_method() value_return
//}

func main() {
	//Interface thay thế cho mọi kiểu dữ liệu
	//interface{} là sử dụng làm value cho kiểu dữ liệu map của Golang
	//ta sẽ có một kiểu map có key luôn là string và value là bất kỳ kiểu dữ liệu nào cũng được
	product := make(map[string]interface{}, 0)
	product["name"] = "Iphone 13 Pro Max"
	product["price"] = 31000000
	product["quantity"] = 40
	fmt.Println(product)

	//để check interface có dùng đúng kiểu dữ liệu muốn ép về, nó sẽ trả về kiểu dữ liệu và giá trị true/false
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	//interface nên được khai báo cùng package sử dụng nó

}
