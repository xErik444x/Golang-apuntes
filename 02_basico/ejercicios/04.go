package main

import "fmt"

func main() {
	var arr [3]int = [3]int{1, 2, 3}
	arr2 := [4]int{4, 5, 6, 7}
	arr3 := [...]int{8, 9, 10} // Aquí no usamos `var`

	fmt.Println(arr)
	fmt.Println(arr2)
	fmt.Println(arr3)
}
