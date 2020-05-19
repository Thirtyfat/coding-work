package stock

import (
	"net/http"
	"net/http/httputil"
	"time"
)

type StockRetriever struct {
	UserAgent string
	TimeOut time.Duration
}

// 指针接收者
func (s *StockRetriever) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(nil)
	}

	response, err := httputil.DumpResponse(resp, true)
	resp.Body.Close()
	if err != nil {
		panic(nil)
	}
	return string(response)

}



