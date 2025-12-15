package array_and_slice

import "fmt"

func DemoArray() {
	arr1 := [3]int{1, 2, 3}
	fmt.Println("arr1:", arr1)

	var arr2 [5]string
	arr2[0] = "Go"
	arr2[1] = "is"
	arr2[2] = "awesome"
	arr2[3] = "!"
	fmt.Println("arr2:", arr2)
}
