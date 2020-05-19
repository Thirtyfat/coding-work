##函数式编程
###函数与闭包
#### 函数式编程 vs 函数指针
* 函数是一等公民，所谓**一等公民**也就是说，参数、变量、返回值都可以是函数
* 高阶函数 函数的参数还可以是函数
* 函数 -> 闭包


#### “正统”函数式编程
* 不可变性：不能有状态，只有常量和函数
* 函数只能有一个参数


> 闭包代码片段
```go
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
```