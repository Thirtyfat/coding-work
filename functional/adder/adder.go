package main

import "fmt"

// 函数体存在局部变量
func adder() func(int) int  {
	sum :=0 // sum 自由变量
	return func(v int) int {
		sum += v // v 局部变量
		return sum
	}
}

func main() {
	a := adder()
	for i := 0 ; i< 10 ; i++{
		fmt.Println(a(i))
	}
}
