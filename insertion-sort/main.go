package main

import "fmt"

func main() {
	// solution 1
	// x := make([]int, 6, 6)
	// x := [6]int{5, 2, 4, 6, 1, 3}
	// for i := 1; i < len(x); i++ {
	// 	j := i
	// 	for j > 0 {
	// 		if x[j-1] > x[j] {
	// 			x[j-1], x[j] = x[j], x[j-1]
	// 		}
	// 		j = j - 1
	// 	}
	// }
	// fmt.Println(x) // answer

	// solution 2
	var items [10]int = [10]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	for i := 1; i < len(items); i++ {
		j := i - 1
		key := items[i]
		for j >= 0 && items[j] > key {
			items[j+1] = items[j]
			j = j - 1
		}
		items[j+1] = key
	}
	fmt.Println(items)
}
