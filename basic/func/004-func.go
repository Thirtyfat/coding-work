package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)


// func
func eval(a,b int , op string) (int , error) {

	switch op {
	case "+":
		return a + b ,nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q,_ :=div(a,b)
		return q ,nil
	default:
		return 0,fmt.Errorf("unsupported operation: %s", op)
	}
}


// func 带余除法
//func div(a, b int) (int, int) {
func div(a, b int) (q,r int) {
	return a / b , a % b
}

func apply(op func(int, int) int , a , b  int) int  {
	// 反射
	// Pointer 指针
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s " +
		"with args " + "(%d , %d)\n", opName,a,b)

	return op(a,b)
}

func pow(a,b int)int  {
	return int (math.Pow(float64(a),float64(b)))
}

// 可变参数列表
func sum(values ...int) int {
	sum := 0
	for i :=range values {
		sum +=values[i]
	}
	return sum
}

func main() {

	//fmt.Println(eval(1,2,"*"))
	//div(21,3)
	//fmt.Println(div(21,3))
	//fmt.Println(eval(21,7,"*"))
	//if result ,err := eval(21,0,"/"); err != nil {
	//	fmt.Printf("error %s: " , err)
	//}else{
	//	fmt.Println(result)
	//}

	// 3 4次方
	fmt.Println(apply(func(a int, b int) int {
		return int (math.Pow(float64(a),float64(b)))
	},3,4))

	//fmt.Println(sum(1,2,3,4,5));
}
