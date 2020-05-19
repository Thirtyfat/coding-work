package main

import (
	"coding-work/tree"
	"fmt"
)

type myTreeNode struct {
	nodes *tree.Node
}

// func (myNode *myTreeNode) postOrder()  指针接收者
func (myNode *myTreeNode) postOrder() {
	if myNode ==nil || myNode.nodes == nil {
		return
	}
	// Fields are assigned without explicit names
	left := myTreeNode{
		myNode.nodes.Left,
	}
	left.postOrder()
	right := myTreeNode{
		myNode.nodes.Right,
	}
	right.postOrder()
	myNode.nodes.Print()
}


func main() {
	var root tree.Node
	root = tree.Node{Value : 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5,nil,nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2);
	root.Right.Left.SetValue(4)



	fmt.Printf("In-order traverse: ")
	//
	root.Traverse()

	fmt.Println()

	fmt.Printf("My own post-order traverse: ")
	myRoot := myTreeNode{&root}
	myRoot.postOrder()
}