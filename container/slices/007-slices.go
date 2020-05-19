package main

import "fmt"

// 半开半闭区间 slice



//func slice(array [...]int) (s[]int){
//	 return array[2:6]
//}


func updateSlice(a [] int){
	a[0] = 100
}


func main() {
	// slice 是对array的一个view
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	fmt.Println("arr[2:6] = ", arr[2:6]) // 2,3,4,5
	fmt.Println("arr[:6] = ", arr[:6])   // 0,1,2,3,4,5
	s2 := arr[2:]							  // 2,3,4,5,6,7
	fmt.Println("arr[2:] = ", s2)		  // 2,3,4,5,6,7
	sAll := arr[:] 							  // 0, 1, 2, 3, 4, 5, 6, 7
	fmt.Println("arr[:] = ", sAll)	      // 0, 1, 2, 3, 4, 5, 6, 7

	fmt.Printf("update s2 After slice \n")

	// S2 =  [2 3 4 5 6 7]
	updateSlice(s2)
	fmt.Printf("s2 %d\n", s2)      		//  [100 3 4 5 6 7]
	fmt.Printf("s2 arr %d\n", arr) 		//  [0 1 100 3 4 5 6 7]

	fmt.Printf("update sAll After slice \n")
	updateSlice(sAll)
	fmt.Printf("sAll %d\n", sAll) 		// [100 1 100 3 4 5 6 7]
	//fmt.Printf("sAll arr %d\n" , arr)

	// Reslice
	fmt.Printf("Reslice \n")
	fmt.Println(sAll) 							// [100 1 100 3 4 5 6 7]
	s3 := sAll[:5]
	fmt.Println(s3) 							// [100 1 100 3 4]
	s4 := s3[2:]
	fmt.Println(s4) 							// [100 3 4]

	// extending slice
	fmt.Printf("extending slice  \n")
	arr1 := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s5 := arr1[2:6] 							// [2,3,4,5]
	// s6 := s5[3:7] 							// panic: runtime error: slice bounds out of range [:7] with capacity 6
	s6 := s5[3:5] 								// [5,6]

	// TODO 1. slice 可以向后扩展，但是不可以向前扩展
	// TODO 2. s[i]不可以超越len(s)，向后扩展不可以超越底层数组 cap(s)
	fmt.Printf("s5=%v,len(s5)=%d,cap(s5)=%d\n",
		s5, len(s5), cap(s5)) 					// s5=[2 3 4 5],len(s5)=4,cap(s5)=6
	fmt.Printf("s6=%v,len(s6)=%d,cap(s6)=%d\n",
		s6, len(s6), cap(s6)) 					// s6=[5 6],len(s6)=2,cap(s6)=3

	s7 := s6[0:3]
	fmt.Printf("s7=%v,len(s7)=%d,cap(s7)=%d\n",
		s7, len(s7), cap(s7))					// 5,6,7

}
