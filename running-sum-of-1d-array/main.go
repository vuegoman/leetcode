package main

import "fmt"

func main() {
	num := [4]int{1, 3, 6, 10}
	result := runningSum(num)
	fmt.Println(result)
}

func runningSum(num [4]int) [4]int {
	for i := 1; i < len(num); i++ {
		num[i] += num[i-1]
	}
	return num
}
