##内建变量类型
* bool，string
* (u) int,(u)int8,(u)int16,(u)int32,(u)int64,uintptr
* byte,rune
* float32,float64,complex64,complex128
* 复数回顾，欧拉公式
#### 欧拉公式
```go
// 欧拉公式
func euler()  {
	// c := 3 + 4i  (4i复数)
	// fmt.Println(cmplx.Abs(c)) // abs to result 5
	// 欧拉公式
	fmt.Println(cmplx.Pow(math.E , 1i  * math.Pi) + 1)
	// result 0+1.2246467991473515e-16i 浮点数的标准都是不准确的
	// complex64 的实部和虚部都是float32  complex128的实部和虚部都是float64
	fmt.Printf("%.3f\n",cmplx.Exp(1i * math.Pi) + 1)
}
```

###强制类型转换
> 转换时必须是指定类型转换，go不像java，可以显示类型转换
> math.Sqrt(a * a + b * b) ×                <br/>
> int(math.Sqrt(float64(a * a + b * b))) √
```go
func triangle() {
	a,b := 11,22
	var c int
	c = int(math.Sqrt(float64(a * a + b * b)))
	fmt.Println(c)
}
```
###常量定义
const 数值可作为各种类型使用
```go
func consts() {
	const fileName = "abc.txt"
	const a , b = 3,4
	var c int
	c = int(math.Sqrt(a * a + b*b))
	fmt.Println(fileName,c)
}
```

###枚举类型
`iota` 表达式
* 普通枚举类型
* 自增枚举类型 
```go
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
```

###变量定义要点回顾
* 变量类型写在变量名之后
* 编译器可推测变量类型
* 没有char，只有rune
* 原生支持复数类型
