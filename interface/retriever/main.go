package main

import (
	"coding-work/interface/retriever/mock"
	"coding-work/interface/retriever/stock"
	"fmt"
	"time"
)

type Retriever interface {
	Get(url string)string
}
const url = "https://www.mogulajiao.com"
// 接口由使用者定义
func download(r Retriever) string {
	return r.Get(url)
}


type Poster interface {
	Post(url string , form map [string] string) string
}
func Post(poster Poster){
	poster.Post(
		url,
		map[string]string{
			"100010":"jack",
			"200020":"James",
		},
	)
}

type RetrieverPost interface {
	Poster
	Retriever
}

func session(r RetrieverPost) string {
	r.Post(url, map[string]string{
		"contents":"another faked mogulajiao.com",
	})
	return r.Get(url)
}



func main() {
	var r Retriever

	mockRetriever := mock.OrderRetriever{Contents: "James"}
	r = &mockRetriever
	inspect(r)

	r = &stock.StockRetriever{
		UserAgent: "9527",
		TimeOut: time.Minute,
	}
	inspect(r)

	if contents , ok := r.(*mock.OrderRetriever) ; ok{
		fmt.Println(contents)
	}else{
		fmt.Printf("contents error ")
	}
	fmt.Println(r)


	// Type assertion
	orderRetriever := r.(*stock.StockRetriever)
	timeOut , userAgent := orderRetriever.TimeOut,orderRetriever.UserAgent
	fmt.Println(timeOut)
	fmt.Println(userAgent)

	fmt.Printf("session \n")
	fmt.Println(session(&mockRetriever))

}

func inspect(r Retriever) {
	fmt.Printf("%T %v\n", r, r)
	switch v := r.(type) {
	case *mock.OrderRetriever:
		fmt.Println("Contents : ", v.Contents)
	case *stock.StockRetriever:
		fmt.Println("UserAgent : ", v.UserAgent)
	}
}
