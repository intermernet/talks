package c2f

import "fmt"

func ExampleCelsius_Fahrenheit() {
	var c Celsius = 40.0
	fmt.Println(c.Fahrenheit())
}

func ExampleNew_freezing() {
	fmt.Printf("%f", New(0.0))
}

func ExampleNew_boiling() {
	fmt.Printf("%f", New(100.0))
}
