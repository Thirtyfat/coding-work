###包结构
* 名字一般使用CamelCase 
    * 首字母大写代表public
    * 首字母小写代表private 
    * public和private针对包而言
* 每个目录一个包，main包包含可执行入口
* 为结构定义的方法必须放在同一个包内，可以是不同文件


### 如何扩充系统类型或者别人的类型？
#### 定义别名
```go
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
```
#### 组合
```go
type Queue [] int

func (q *Queue) Push(v int){
	// q 所指向的slice被改变
	*q = append(*q,v)
}


func (q *Queue) Pop() int{
	head :=(*q)[0]
	*q  = (*q)[1:]
	return head
}

func (q *Queue) IsEmpty()bool  {
	return len(*q) == 0
}
```

#### 使用内嵌(Embedding)方式扩展已有的类型
```go
// 继承当中可以把子类类型赋值给父类的指针或者引用
// var baseRoot *tree.Node
// 在java当中最通用的做法 ： 将子类指针赋值给基类指针
// 但是在go语言当中，如下语句会抛出一个异常
// baseRoot :=&node // cannot use &node (type *myTreeNode) as type *tree.Node in assignment
```


