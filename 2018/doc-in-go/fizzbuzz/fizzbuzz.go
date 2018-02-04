package main

import "fmt"

func main() {
	var mask uint = 810092048
	s := []int{70, 105, 122, 122, 66, 117, 122, 122}
	for i := 1; i <= 100; i++ {
		c := mask & 3
		switch {
		case c > 0:
			for _, i := range s[(c-1)%2*4 : 8-2*(1>>(c-1)*2)] {
				fmt.Printf("%c", i)
			}
			fmt.Println()
		default:
			fmt.Println(i)
		}
		mask = mask>>2 | c<<28
	}
}
