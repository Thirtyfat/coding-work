###函数
* func eval (a,b int,op string) int 
```go
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
if result ,err := eval(21,0,"/"); err != nil {
		fmt.Printf("error %s: " , err)
	}else{
		fmt.Println(result)
	}
}s
```

> 带余除法
```go
func div(a, b int) (q,r int) {
	return a / b , a % b
}
fmt.Println(div(21,3))
```
* 函数可返回多个值
* 函数返回多个值时可以取名字
* 仅用于非常简单的函数

> 匿名函数
```go
func apply(op func(int, int) int , a , b  int) int  {
	// 反射
	// Pointer 指针
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s " +
		"with args " + "(%d , %d)\n", opName,a,b)

	return op(a,b)
}
fmt.Println(apply(func(a int, b int) int {
		return int (math.Pow(float64(a),float64(b)))
	},3,4))

reponse
Calling function main.main.func1 with args (3 , 4)
81
```
* 函数作为参数

> 可变参数列表
```go
func sum(values ...int) int {
	sum := 0
	for i := range values {
		sum +=values[i]
	}
	return sum
}
fmt.Println(sum(1,2,3,4,5));
```

###函数语法要点回顾
* 返回值类型写在最后面
* 可返回多个值
* 函数可作为参数
* 没有默认参数，可选参数，只有一个可变参数列表