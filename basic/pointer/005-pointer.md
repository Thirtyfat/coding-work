###指针相关
```go
	var a int = 2
	var pa *int = &a
	*pa = 3
	fmt.Println(a)
```

* go语言不能运算

####参数传递
> go语言的参数是值传递还是引用传递？
* c、c++ 即可以值传递也可以引用传递
* java和python绝大部分都是引用传递，除了一些系统的自建类型以外

> 什么是值传递什么是引用传递？
```c
void pass_by_val(int a){
    a++
}

void pass_by_ref(int& a){
    a++
}

int main(){
   int a = 3;
   pass_by_val(a); // 3
   pass_by_ref(a); // 4
}
```
* pass_by_val 属于值传递，copy到函数的时候，值发生变化，但是此时main函数当中的值并没有发生变化
* pass_by_ref 属于引用传递，函数跟main方法同时引用同一个变量，当函数内发生变化时，则直接影响main函数当中的值


> go语言只有值传递一种方式
* go语言在自定义类型的时候，就需要考虑到是当做指针来使用还是当成值来用。


```go
func swap(a, b int)  {
	a,b = b,a
}

a,b := 3,4
swap(a,b)
fmt.Println(a,b)

reponse
3 4
```

```go
func swapPointer(a, b *int)  {
	*a,*b = *b,*a
}
a,b := 3,4
swapPointer(&a,&b)
fmt.Println(a,b)

reponse
4 3
```

```go
func swapReturn(a, b int) (int,int)  {
	return b,a
}
a,b := 3,4
a,b = swapReturn(a,b)
fmt.Println(a,b)

reponse
4 3 
```