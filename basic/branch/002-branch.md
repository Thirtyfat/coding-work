###条件分支语句 if
```go
func main() {
	const filename = "abc.txt"
	contents, err := ioutil.ReadFile(filename)
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Printf("%s\n" , contents)
	}
}
```
* if条件语句当中可以赋值
* if条件语句当中赋值的作用域只存在这个if块当中 
```go
func main() {
	const filename = "abc.txt"
	if contents, err := ioutil.ReadFile(filename) ; err !=nil {
		fmt.Println(err)
	}else{
		fmt.Printf("%s\n" , contents)
	}
}
```
###switch
* switch 会自动break, 除非使用fallthrough
```go
func grade(score int) string{
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score : %d\n" , score))
	case score <= 60:
		g =  "D"
	case score <= 70:
		g =  "C"
	case score <= 80:
		g =  "B"
	case score <= 100:
		g =  "A"
	}
	return g
}
```
* for，if条件当中可以允许没有括号
* if条件语句当中可以赋值
* 没有while
* switch不需要break，也可以直接switch多个条件

