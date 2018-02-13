package examples

import (
	"fmt"
)

func ExampleEncode() {
	fmt.Println(Encode([]byte("Hi #golangsyd")))
}

func ExampleEncode_withOutput() {
	fmt.Print(Encode([]byte("Hi #golangsyd")))
	// Output: SGkgI2dvbGFuZ3N5ZA
}

func ExampleDecode() {
	out, err := Decode("SGkgI2dvbGFuZ3N5ZA")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", out)
	// Output: Hi #golangsyd
}

func ExampleDecode_fail() {
	out, err := Decode("SGkgI2dvbGFuZ3N5ZA")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", out)
	// Output: Hi #golandsyd
}

func ExampleUnordered() {
	ints := []int{0, 1, 2, 3}
	for _, i := range Unordered(ints) {
		fmt.Println(i)
	}
	// Unordered output:
	// 3
	// 2
	// 1
	// 0
}
