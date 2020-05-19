```go
// 值传递
func (node treeNode) print()  {
	fmt.Print(node.value , " ")
}
//  值传递 方法功能同上
func print(node treeNode)  {
	fmt.Print(node.value, " ")
}
// 值传递 func (node treeNode) setValue(value int)
// 引用传递 func (node *treeNode) setValue(value int)
#使用指针作为方法接收者，不然的话该结构只是值传递，值传递是一个copy，copy之后的数据在外层也不会显示，作用域只限于函数内部
func (node *treeNode) setValue(value int)  {
	node.value = value
}
```

* 普通函数不增加指针就是值传递，增加指针就是引用传递
* 显示定义和命名方法接收者
* 只有使用指针裁才可以改变结构内容
* nil指针也可以调用方法
> nil 指针方法调用
```go
func (node *treeNode) setValue(value int)  {
	if node == nil {
		fmt.Println("Setting value to nil node. Ignored." )
		return
	}
	node.value = value
}

var tRoot *treeNode
tRoot.setValue(200)
tRoot = &node
tRoot.setValue(300)
pRoot.print()
```
* go语言仅支持封装，不支持继承和多态
* go语言没有class，只有struct

### 值接收者 指针接收者
* 要改变内容，必须使用指针接收者，值接收者只是对值的拷贝
* 结构过大时也需要考虑使用**指针接收者**
* 一致性问题，如果有指针接收者，最好都是指针接收者
* 值接收者是go语言特有
* c++也存在指针接收者，java 的this对象引用也是属于一个指针
* 值/指针接收者均可接收值/指针




