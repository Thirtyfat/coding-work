package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// 1,1,2,3,5,8,13,21...
// a,b
// 	 a,b
// 	   a,b
func fibonacci() intGen  {
	a , b :=0,1
	return func() int {
		a , b = b , a +b
		return a
	}
}

type intGen func() int

//  函数实现接口
func (g intGen) Read(
		p []byte)(n int ,err error) {
	next := g()
	// 增加停止条件
	if next > 10000{
		return 0,io.EOF
	}
	s := fmt.Sprintf("%d\n" , next)
	// TODO: incorrect if p is too small!
	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader){
	newScanner := bufio.NewScanner(reader)
	for newScanner.Scan(){
		fmt.Println(newScanner.Text())
	}
}


func main() {
	f :=fibonacci()
	printFileContents(f)
}
