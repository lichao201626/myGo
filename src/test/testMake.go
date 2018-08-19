package main

import "fmt"

/* func main() {
	arr := make([]int, 5)
	printSlice("a", arr)
}
*/
func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
