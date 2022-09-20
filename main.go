package main

import (
	"awesomeProject/sort"
	"fmt"
)

func main() {
	arr := []int{5, 2, 6, 3, 1}
	sort.Quick(arr)
	fmt.Println(arr)
}
