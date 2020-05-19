###循环语句
* for的条件当中不需要括号
* for的条件当中可以省略初始条件，结束条件，递增表达式

> 省略初始条件
```go
func convertToBin(n int ) string {
	result := ""
	for ; n > 0 ; n /= 2{
		// 最低位 % 2
		lsb := n % 2
		// strconv.Itoa() 转字符串 
		result = strconv.Itoa(lsb) + result;
	}
	return result;
}

junit test > 
fmt.Println(
		convertToBin(10),
		convertToBin(13),
		convertToBin(9527),

	)
response: 1010 1101 10010100110111
```
> 省略递增条件
```go
func printFile(filename string){
	// 打开文件
	file , err := os.Open(filename)
	if err != nil{
		// error
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	// while 
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
```
>  省略结束条件 无限循环
```go
func forever()  {
	for {
		fmt.Println("aaa")
	}
}
```
















