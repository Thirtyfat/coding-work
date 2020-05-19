package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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

func forever()  {
	for {
		fmt.Println("aaa")
	}
}


func main() {
	//fmt.Println(
	//	convertToBin(10),
	//	convertToBin(13),
	//	convertToBin(9527),
	//
	//)
	// printFile("abc.txt")
	forever()
}
