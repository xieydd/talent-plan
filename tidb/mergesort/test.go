package main

import (
	"fmt"
)

func Print(arr []int) {
	for _, v := range arr {
		fmt.Printf("%d\n", v)
	}
	arr[3] = 5
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("%v\n", arr[0:7])
	Print(arr)
	fmt.Printf("%d\n", arr[3])
	a := 4
	b := 5
	c := (a + b) / 2
	fmt.Printf("%d\n", c)
}
