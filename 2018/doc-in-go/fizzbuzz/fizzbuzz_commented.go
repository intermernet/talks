package main

import "fmt"

func main() {
	// mask is the decimal representation of the first 15 results of the standard "FizzBuzz"
	// challenge in reverse order. These 15 results can repeat infinitely.
	// Binary representation (30 bits long): 11 00 00 01 00 10 01 00 00 01 10 00 01 00 00
	// Decimal representation when split into bit-pairs: 3 0 0 1 0 2 1 0 0 1 2 0 1 0 0
	var mask uint = 810092048
	// s is a a slice of integers repesenting the Unicode characters "F i z z B u z z"
	s := []int{70, 105, 122, 122, 66, 117, 122, 122}
	for i := 1; i <= 100; i++ {
		// c is the result of a bitwise AND of 3 (binary 11) on mask.
		// It returns the current right-most bit-pair from mask.
		c := mask & 3
		// If c != 0 then it must be 1, 2 or 3
		if c > 0 {
			// To get the slice bounds, do some bit-twiddling and then loop over s
			// converting each value in the slice to Unicode.			//
			// Output table: c  (c-1)%2*4  8-2*(1>>(c-1)*2  Output  Unicode
			//               1          0                4  s[0:4]  "Fizz"
			//               2          4                8  s[4:8]  "Buzz"
			//               3          0                8  s[0:8]  "FizzBuzz"
			for _, i := range s[(c-1)%2*4 : 8-2*(1>>(c-1)*2)] {
				fmt.Printf("%c", i)
			}
			fmt.Println() // Go to a new line
		} else { // If c == 0 then just print the value of i
			fmt.Println(i)
		}
		// Pop off the 2 right-most bits of mask and put the previous result of c back onto the left. This makes the 15 bit-pairs loop forever.
		mask = mask>>2 | c<<28
	}
}
