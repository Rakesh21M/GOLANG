package main

import "fmt"

func factorial(num int) int {
	if num == 1 {
		return 1
	} else {
		return (factorial(num) * factorial(num-1))
	}
}
func main() {
	var num int = 4
	val := factorial(num)
	fmt.Println(val)
}
