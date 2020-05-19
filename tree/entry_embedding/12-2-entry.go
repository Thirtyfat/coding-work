package main

import (
	"coding-work/tree"
	"fmt"
)

type myTreeNode struct {
	*tree.Node // Embedding 内嵌
}

// func (myNode *myTreeNode) postOrder()  指针接收者
func (myNode *myTreeNode) postOrder() {
	if myNode ==nil || myNode.Node == nil {
		return
	}
	// Fields are assigned without explicit names
	left := myTreeNode{myNode.Left}
	right := myTreeNode{myNode.Right}
	left.postOrder()
	right.postOrder()
	myNode.Print()
}

// shadowed
func (myNode * myTreeNode) Traverse()  {
	fmt.Println("This method is shadowed")
}



func main() {

	node := myTreeNode{&tree.Node{Value: 3}}
	node.Left = &tree.Node{}
	node.Right = &tree.Node{5,nil,nil}
	node.Right.Left = new(tree.Node)
	node.Left.Right = tree.CreateNode(2);
	node.Right.Left.SetValue(4)

	fmt.Printf("In-order traverse: ")
	node.Traverse()

	fmt.Printf("My own post-order traverse: ")
	node.postOrder()
	fmt.Println()

	// 继承当中可以把子类类型赋值给父类的指针或者引用
	// var baseRoot *tree.Node
	// 在java当中最通用的做法 ： 将子类指针赋值给基类指针
	// 但是在go语言当中，如下语句会抛出一个异常
	// baseRoot :=&node // cannot use &node (type *myTreeNode) as type *tree.Node in assignment
}
