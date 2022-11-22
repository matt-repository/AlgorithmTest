package Test

import "fmt"

func Test() {
	var map1 map[int]int
	fmt.Println(map1)
	fmt.Println(map1 == nil)
	var map2 map[int]int = map[int]int{}

	fmt.Println(map2 == nil)
}
