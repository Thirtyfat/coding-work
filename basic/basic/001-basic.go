package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func variableZeroValue()  {
	var a int
	var b string
	fmt.Printf("%d %q\n" , a,b)
}

func variableInitialValue(){
	var a,b,c int = 1,2,3
	var d string = "abc"
	fmt.Println(a,b,c,d)
}
func variableTypeDeduction(){
	var a,b,c,d,e,f = 3,4,5,true,false,"true"
	fmt.Println(a,b,c,d,e,f)
}

func variableShorter()  {
	a,b,c,d,e,f := 3,4,5,true,false,"true"
	// 赋值
	b = 6
	fmt.Println(a,b,c,d,e,f)
}




// 欧拉公式
func euler()  {
	//c := 3 + 4i  (4i复数)
	//fmt.Println(cmplx.Abs(c)) // abs to result 5
	// 欧拉公式
	fmt.Println(cmplx.Pow(math.E , 1i  * math.Pi) + 1)
	// result 0+1.2246467991473515e-16i 浮点数的标准都是不准确的
	// complex64 的实部和虚部都是float32  complex128的实部和虚部都是float64
	fmt.Printf("%.3f\n",cmplx.Exp(1i * math.Pi) + 1)
}

// 强制类型转换
func triangle() {
	a,b := 11,22
	var c int
	c = int(math.Sqrt(float64(a * a + b * b)))
	fmt.Println(c)
}

// 常量
func consts() {
	const fileName = "abc.txt"
	const a , b = 3,4
	var c int
	c = int(math.Sqrt(a * a + b*b))
	fmt.Println(fileName,c)
}

// 枚举
func enums() {
	const (
		java = iota
		_
		golang
		cpp
		javascript
	)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(java,javascript,golang,cpp)
	fmt.Println(b,kb,mb,gb,tb,pb)
}

func main() {
	//variableZeroValue()
	//variableInitialValue()
	//variableTypeDeduction()
	//variableShorter()
	//euler()
	//triangle()
	//consts()
	enums()
}
