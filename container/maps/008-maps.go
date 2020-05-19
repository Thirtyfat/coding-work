package main

import "fmt"
/*
 * 创建：make(map[string]int)
 * 获取元素 ： m[key]
 * key不存在时，获得value类型的初始值
 * 用value,ok := m[key] 来判断key是否存在
 * 用delete删除一个key

 * 使用range遍历key，或者遍历key，value对
 * 不保证遍历顺序，如需顺序，需要手动对key进行排序
 * 使用len获得元素个数

 * map使用哈希表，必须可以比较相等
 * 除了slice，map，function的内建内省都可以作为key
 * struct类型不包含上述字段，也可以作为key
 */
func main() {
	m :=map[string]string{
		"name":"ccmouse",
		"course":"golang",
		"site":"imooc",
		"quality":"notbad",
	}
	fmt.Println(m) 								// map[course:golang name:ccmouse quality:notbad site:imooc]

	m2 := make(map[string]int) 					// m2 == empty map
	var m3 map[string]int 						// m3 == nil

	fmt.Println(m,m2,m3)						// map[course:golang name:ccmouse quality:notbad site:imooc] map[] map[]

	fmt.Printf("traversing map... \n")
	for k,v :=range m{							// range
		fmt.Println(k,v)
	}

	fmt.Printf("Getting values... \n")
	valuesName ,ok:= m["name"] 					// ccmouse true
	fmt.Println(valuesName,ok)
	if coursesName , ok := m["couse"]; ok {		// false  result  ===> key does not exist !
		// true
		fmt.Println(coursesName)
	}else {
		// false
		fmt.Println("key does not exist ! ")
	}

	fmt.Printf("Deleting values... \n")
	if valueName ,ok := m["name"]; ok{			// ccmouse true
		println(valueName,ok) 					// true
		delete(m,"name")					// delete value by maps keys
	}else{
		fmt.Println("key does not exist ! ")
	}
	name , ok := m["name"]						// false
	println(name,ok)
}
