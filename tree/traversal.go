package tree

// 中序遍历，先左，再中，再右
func (node *Node) Traverse(){
	node.TraverseFunc(func(node *Node) {
		node.Print()
	})
}



func (node *Node) TraverseFunc(f func(node *Node)){
	if node == nil{
		return
	}

	node.Left.TraverseFunc(f)
	f(node)
	node.Right.TraverseFunc(f)
}