package go_map

import "fmt"

func DemoMap() {
	map1 := map[string]int{
		"apple":  5,
		"banana": 10,
		"orange": 15,
	}
	fmt.Println("map1:", map1)

	map2 := make(map[int]string)
	map2[1] = "one"
	map2[2] = "two"
	map2[3] = "three"
	fmt.Println("map2:", map2)

	value, exists := map1["banana"]
	if exists {
		fmt.Println("The value for 'banana' is:", value)
	} else {
		fmt.Println("'banana' does not exist in map1")
	}

	delete(map1, "orange")
	fmt.Println("map1 after deleting 'orange':", map1)
}
