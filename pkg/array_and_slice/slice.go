package array_and_slice

import "fmt"

func DemoSlice() {
	slice1 := []int{10, 20, 30, 40, 50}
	fmt.Println("slice1:", slice1)

	slice2 := make([]string, 3, 5)
	slice2[0] = "Go"
	slice2[1] = "is"
	slice2[2] = "fun"
	fmt.Println("slice2:", slice2)

	slice2 = append(slice2, "!")
	fmt.Println("slice2 after append:", slice2)
}
