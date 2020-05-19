package queue

type Queue [] interface{}

func (q *Queue) Push(v interface{}){
	// q 所指向的slice被改变
	*q = append(*q,v)
}


func (q *Queue) Pop() interface{}{
	head :=(*q)[0]
	*q  = (*q)[1:]
	return head
}

func (q *Queue) IsEmpty()bool  {
	return len(*q) == 0
}