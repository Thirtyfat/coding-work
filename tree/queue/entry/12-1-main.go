package main

import (
	"coding-work/tree/queue"
	"fmt"
)

func main() {
	// 先进先出队列
	q := queue.Queue{1}
	// q每次执行对象都不相同，因为属于指针接收者
	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	q.Push("wow")
	fmt.Println(q.IsEmpty())
	q.Pop()
	fmt.Println(q.IsEmpty())


}
