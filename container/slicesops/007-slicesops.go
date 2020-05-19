package main

import "fmt"

// 添加元素如果超过cap，系统会重新分配更大的底层数组
func printSlice(s []int) {
	fmt.Printf("v=%v,len=%d , cap=%d \n" ,s, len(s),cap(s))
}

func main() {
	fmt.Printf("creating  slice ")
	var s []int  											// zero value  for slice is nil
	for i :=0 ; i<100;i++{
		printSlice(s)										//
		s = append(s, i * 2 +1)
	}
	fmt.Println(s)


	s1 := []int{2,4,6,8}
	printSlice(s1) 											// v=[2 4 6 8],len=4 , cap=4

	s2 := make([]int, 16)									// 内建函数
	printSlice(s2) 											// v=[0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0],len=16 , cap=16
	// len 10 cap 32
	s3 := make([]int, 10 , 32)								// cap
	printSlice(s3) 											// v=[0 0 0 0 0 0 0 0 0 0],len=10 , cap=32

	fmt.Printf("Copying  slice \n")					//
	copy(s2,s1)
	printSlice(s2) 											// [2 4 6 8 0 0 0 0 0 0 0 0 0 0 0 0],len=16 , cap=16

	fmt.Printf("Deleting elements from slice \n")
	// deleting elements 8
	// TODO  printSlice(s2[:3], s2[4:]...) 如果使用该输出的话，那么当前s2的指针指向则还是之前的
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2)											// v=[2 4 6 0 0 0 0 0 0 0 0 0 0 0 0],len=15 , cap=16


	fmt.Printf("Popping  from  front slice \n")
	front := s2[0] 											// response 2  poping len[0] value[2]
	s2 = s2[1:] 											//  [4 6 0 0 0 0 0 0 0 0 0 0 0 0 0]
	fmt.Println(front)
	printSlice(s2) 											// v=[4 6 0 0 0 0 0 0 0 0 0 0 0 0 0],len=15 , cap=15


	fmt.Printf("Popping  from  back slice \n")
	back := s2[len(s2) - 1 ] 								// response 0
	s2 = s2[:len(s2) - 1 ] 									//  [4 6 0 0 0 0 0 0 0 0 0 0 0 0 ]
	fmt.Println(back)
	printSlice(s2) 											// v=[4 6 0 0 0 0 0 0 0 0 0 0 0 ],len=14 , cap=15
}
