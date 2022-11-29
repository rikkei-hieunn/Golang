package main

import "fmt"

//Generics giúp định nghĩa dược những method/function chấp nhận các tham số chung,
//không cần chỉ định rõ nó thuộc kiểu dữ liệu gì, khi được gọi thì sẽ quyết định việc này

//  Generics gồm 2 phần :

// Print 1 biến có thể khai báo cụ thể 1 kiểu dữ liệu nhất định
func Print[T int](s ...T) {
	for _, v := range s {
		fmt.Print(v)
	}
}

// PrintA Print 2  nếu biến k cần rằng buộc gì thì chúng ta có thể dùng "any"
func PrintA[T any](s ...T) {
	for _, v := range s {
		fmt.Print(v)
	}
}

//	K và V sẽ là 2 kiểu dữ liệu generics cho hàm Map. Chúng ta chưa rõ nó là gì nhưng chúng ta muốn mọi kiểu dữ liệu đều có thể dùng được nên cả 2 đều là any.
//	Ở phần định nghĩa tham số truyền vào chúng ta sẽ cần một mảng K: []K và một hàm transform nhận kiểu K và trả về kiểu V.
//	Hàm này là để người gọi muốn làm gì với mỗi item K thì làm, miễn return V là okie.
//	Cuối cùng hàm Map phải trả về một mảng mới: []V
//	Ở hàm main, mình chạy thử với 2 trường hợp mảng int và mảng string. Các bạn có thể đổi lại tuỳ mục đích biến đổi item trong mảng ban đầu nhé.

func Map[K, V any](s []K, transform func(K) V) []V {
	rs := make([]V, len(s))
	for i, v := range s {
		rs[i] = transform(v)
	}
	return rs
}

//Sử dụng constraint Generics trong Go:
// khi khai báo biến vs "any"  thì không phải kiểu dữ liệu nào chúng ta cũng có thể thực hiện được công việc mong muôn
//vd ( kp value nào cũng thực hiện so sánh được )
//Giải pháp

type SignedInteger interface {
	int | int8 | int16 | int32 | int64
}

//Bây giờ T bắt buộc phải thuộc một trong những kiểu int | int8 | int16 | int32 | int64 thì mới dùng phép toán so sánh được.

func Smallest[T SignedInteger](s []T) T {
	r := s[0]
	for _, v := range s[1:] {
		if r > v {
			r = v
		}
	}
	return r
}
func main() {
	arr := []int{1, 2, 3}
	resultArr := Map(arr, func(v int) int { return v * 2 })
	fmt.Println(resultArr)

	arr2 := []string{"a", "b", "c"}
	resultArr2 := Map(arr2, func(v string) string { return v + v })
	fmt.Println(resultArr2)

}
