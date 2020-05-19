package tree

import "fmt"

type Node struct {
	Value int
	Left , Right *Node
}

// 工厂函数
func CreateNode (value int ) *Node {
	// 返回的是局部变量的地址
	return &Node{Value: value}
}

// 值传递
func (node Node) Print()  {
	fmt.Print(node.Value , " ")
}
//  值传递 方法功能同上
func Print(node Node)  {
	fmt.Print(node.Value, " ")
}
// 值传递 func (node Node) setValue(value int)
// 引用传递 func (node *Node) setValue(value int)
func (node *Node) SetValue(value int)  {
	if node == nil {
		fmt.Println("Setting value to nil node. Ignored." )
		return
	}
	node.Value = value
}
