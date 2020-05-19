package mock

type OrderRetriever struct {
	Contents string
}

// 值接收者
func (r *OrderRetriever) Get(url string) string {
	return r.Contents
}

func (p *OrderRetriever) Post(url string,from map[string]string) string {
	p.Contents = from["contents"]
	return p.Contents
}
