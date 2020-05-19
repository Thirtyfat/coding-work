package infra
//
//import (
//	"io/ioutil"
//	"net/http"
//)
//
//type Retriever struct {}
//
//func (Retriever) Get(url string) string{
//	resp, err := funcName(url)
//	if err != nil || resp == nil{
//		panic(err)
//	}
//	defer resp.Body.Close()
//	bytes, _ := ioutil.ReadAll(resp.Body)
//	return string(bytes)
//}
//
//func funcName(url string) (*http.Response, error) {
//	return http.Get(url)
//}
