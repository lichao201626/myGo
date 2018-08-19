package main

import "fmt"

/* func main() {
	arr := []int64{5, 2, 124, 1, 4, 13, 52}
	sort(arr)
} */

func sort(ar []int64) {
	fmt.Println("hello world", ar)
	for i := 1; i < len(ar); i++ {
		for j := 0; j < len(ar)-1; j++ {
			if ar[j] < ar[j+1] {
				tmp := ar[j]
				ar[j] = ar[j+1]
				ar[j+1] = tmp
			}
		}
	}

	for _, k := range ar {
		// fmt.Println(v)
		fmt.Printf("%v \n", k)
		// fmt.Println(k)
		// fmt.Println(ar[v])
	}
}
