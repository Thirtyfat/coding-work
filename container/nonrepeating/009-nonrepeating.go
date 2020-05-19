package main

import "fmt"

func lengthOfNonRepeatingSubstr(s string) int{
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0
	//  ch int32
	for i ,ch :=range [] rune(s){
		//if lastI,ok := lastOccurred[ch]; ok && lastI >= start{
		//	start = lastI + 1
		//}
		if lastOccurred[ch] >= start{
			start = lastOccurred[ch] + 1
		}
		if i -start +1 > maxLength{
			maxLength = i -start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

func main() {
	var s1 string = "abcabcbb"
	var s2 string = "瑞信咖啡瑞信咖啡瑞信咖啡"
	//var s3 string = "pwwkew"

	fmt.Println(
		lengthOfNonRepeatingSubstr(s1),
		lengthOfNonRepeatingSubstr(s2),
		//lengthOfNonRepeatingSubstr(s3),
		)

	//fmt.Println(s1,s2,s3)
	//
	//start := 0
	//maxLength := 0
	//
	//for k,v :=range []rune(s1){
	//	if k -start +1 > maxLength{
	//		maxLength = k -start + 1
	//	}
	//	fmt.Println(k,v)
	//}

}
