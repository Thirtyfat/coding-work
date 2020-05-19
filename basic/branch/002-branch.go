package main

import (
	"fmt"
)

// switch 会自动break, 除非使用fallthrough
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

func main() {
	//const filename = "abc.txt"
	//if contents, err := ioutil.ReadFile(filename) ; err !=nil {
	//	fmt.Println(err)
	//}else{
	//	fmt.Printf("%s\n" , contents)
	//}
	fmt.Println(
		grade(100),
		grade(90),
		grade(80),
		grade(70),
		grade(-1),
		)

}