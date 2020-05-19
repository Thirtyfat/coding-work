package main

import "fmt"

func main() {
	arr :=[...]int {0,1,2,3,4,5,6,7}
	s1 :=arr[2:6]
	fmt.Println(s1) 							// [2,3,4,5]
	s2 := s1[3:5]
	fmt.Println(s2) 							// [5,6]

	s3 := append(s2,10)
	fmt.Println(s3) 							// [5,6,10]

	s4 := append(s3,11)
	fmt.Println(s4) 							// [5,6,10,11]

	s5 := append(s4,12)
	fmt.Println(s5) 							// [5 6 10 11 12]

	// s4 and s5 no longer view arr.
	// 向slice当中添加元素
	// 1. 添加元素时如果超越cap，系统会重新分配更大的底层数组
	// 2. 由于是值传递的关系，必须接受append的返回值
	fmt.Println(arr) 							// [0,1,2,3,4,5,6,10]  s4 and s5 no longer view arr.
	// TODO 由于 s2 := s1[3:5]获取的值是[5,6] 当前cap当中还存在元素7
	// TODO 当执行s3 := append(s2,10) 将当前view替换成[0,1,2,3,4,5,6,10],所以arr就是[0,1,2,3,4,5,6,10]
}
