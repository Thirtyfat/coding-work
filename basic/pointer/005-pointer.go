package main

import "fmt"

func swap(a, b int)  {
	a,b = b,a
}
func swapPointer(a, b *int)  {
	*a,*b = *b,*a
}
func swapReturn(a, b int) (int,int)  {
	return b,a
}
/*
指针不能运算
参数传递： c++ 既可以值传递也可以引用传递 ， java，python是属于引用传递
*/

func main() {
	a,b := 3,4
	swap(a,b)
	fmt.Println(a,b)

	//swapPointer(&a,&b)
	//fmt.Println(a,b)

	a,b = swapReturn(a,b)
	fmt.Println(a,b)
}
