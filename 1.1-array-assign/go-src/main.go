package main

import "fmt"

func main() {
	var arr []int = []int{1, 2, 3}
	var arr2 []int = arr

	arr2[1] = 20
	fmt.Println("arr= ", arr)
	fmt.Println("arr2= ", arr2)

}
