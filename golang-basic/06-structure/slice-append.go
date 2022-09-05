package main

import "fmt"

func main() {
	sliceA := []int{1, 2, 3}
	sliceB := append(sliceA, 4, 5)
	fmt.Println(sliceA, sliceB)

	slice1 := []int{1, 2, 3}
	slice1 = append(sliceA, 4, 5)
	slice2 := make([]int, 3)
	copy(slice2, slice1)
	fmt.Println(slice1, slice2)
}
