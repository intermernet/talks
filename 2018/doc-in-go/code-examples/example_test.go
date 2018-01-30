package examples

import (
	"fmt"
)

func ExampleEncode() {
	fmt.Println(Encode([]byte("Hi #golangsyd")))
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
