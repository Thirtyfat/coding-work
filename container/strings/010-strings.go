package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "瑞信咖啡~!!"												// utf-8
	fmt.Println(len(s)) 											// 15
	fmt.Printf("%s\n",[]byte(s)) 							// 瑞信咖啡~!!
	fmt.Printf("%X\n",[]byte(s)) 							// E7919EE4BFA1E59296E595A17E2121
	for _,b := range []byte(s){
		fmt.Printf("%X ",b)									// E7 91 9E E4 BF A1 E5 92 96 E5 95 A1 7E 21 21
	}
	fmt.Println()

	for i,ch := range s {									// ch is a rune
		fmt.Printf("(%d %X)",i,ch)	// (0 745E)(3 4FE1)(6 5496)(9 5561)(12 7E)(13 21)(14 21)
											// E7 91 9E to 745E 	(utf8 ---> unicode)
	}
	// utf8 ----> unicode
	fmt.Println()

	fmt.Println(utf8.RuneCountInString(s)) 					// 7
	bytes :=[]byte(s)
	for len(bytes) > 0{
		ch,size :=utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ",ch)							// 瑞 信 咖 啡 ~ ! !
	}
	fmt.Println()

	for i,ch := range [] rune(s){
		fmt.Printf("(%d %c)",i,ch)					// (0 瑞)(1 信)(2 咖)(3 啡)(4 ~)(5 !)(6 !)

	}
	fmt.Println()
}
