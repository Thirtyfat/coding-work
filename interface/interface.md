###接口

* 面向接口编程
* 使用google guice 实现依赖注入


### duck typing
* 描述事物的外部行为而非内部结构
* 严格说go语言属于结构化类型系统，类似duck typing
* go语言的duck typing
    * 同时需要实现Readable,Appendable 该如何解决？(apache polygene)
    * 同时具有python/c++的duck typing 的灵活性
    * 具备java 的类型检查

* go语言接口由使用者定义
* 接口的实现是隐式的
### 接口变量里面有什么？
* 接口变量当中存在**实现者的类型**以及**实现者的值**，或者可以说实现者的指针指向实现者
* 接口变量自带指针
* 接口变量采用值传递，指针接收者实现只能以指针方式使用，值接收者则都可以

### 如何查看接口变量
* 表示任何类型 interface{}
* Type Assertion
* Type Switch

### 接口组合，修改接口由值接收改为指针引用
```go
mock.OrderRetriever {James}
Contents :  James
*stock.StockRetriever &{9527 1m0s}
UserAgent :  9527
contents error &{9527 1m0s}
1m0s
9527
session 
James



*mock.OrderRetriever &{James}
Contents :  James
*stock.StockRetriever &{9527 1m0s}
UserAgent :  9527
contents error &{9527 1m0s}
1m0s
9527
session 
another faked mogulajiao.com

```
