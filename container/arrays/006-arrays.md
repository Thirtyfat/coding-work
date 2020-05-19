###数组，切片和容器
```go
    var arr1 [5] int
	arr2 :=[3]int{1,3,5}
	arr3 :=[...]int {2,4,6,8,10}
reponse
[0 0 0 0 0] [1 3 5] [2 4 6 8 10]
```
> 二元数组
```go
var grid [4][5] int
reponse
[[0 0 0 0 0] [0 0 0 0 0] [0 0 0 0 0] [0 0 0 0 0]]
```

> 遍历数组
```go
for i :=0;i<len(arr3);i++{
	fmt.Println(arr3[i])
}
# 获取下标 or 获取下标以及值
for _,v := range arr3{
    fmt.Println(v)
}
```
```go
	maxi := -1
	maxValue := -1
	for i , v :=range arr3{
		if v > maxValue{
			maxi , maxValue = i,v
		}
	}
	fmt.Printf("maxi :  %d\n", maxi)
	fmt.Printf("maxValue :  %d\n",maxValue)
```

### 值类型数组
```go
var arr1 [5] int
arr2 :=[3]int{1,3,5}
arr3 :=[...]int {2,4,6,8,10}

func printArray(arr[5] int)  {
	arr[0] = 100
	for i,v := range arr{
		fmt.Println(i,v)
	}
}

response
printArray(arr1)    // 0 100
                    //   1 0
                    //   2 0
                    //   3 0
                    //   4 0

printArray(arr2)    //  cannot use arr2 (type [3]int) as type []int in argument to printArray
printArray(arr3)     // 0 100
                    //   1 4
                    //   2 6
                    //   3 8
                    //   4 10
fmt.Println(arr1,arr3) // [0 0 0 0 0] [2 4 6 8 10]
```
* [10]int 和[20]int 是不同类型
* 调用func f (arr [10] int) 会拷贝数据

> 采用指针方式调用数组 (引用传递)
```go
func printArray(arr *[5] int)  {
	arr[0] = 100
	for i,v := range arr{
		fmt.Println(i,v)
	}
}
printArray(&arr1)  // 0 100
                  //   1 0
                  //   2 0
                  //   3 0
                  //   4 0
//printArray(arr2) //  cannot use arr2 (type [3]int) as type []int in argument to printArray
printArray(&arr3) // 0 100
                  //   1 4
                  //   2 6
                  //   3 8
                  //   4 10
fmt.Println(arr1,arr3) // [100 0 0 0 0] [100 4 6 8 10]
```
* 在go语言当中一般不使用数组，一般使用切片

