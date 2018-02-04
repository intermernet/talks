package c2f

import "fmt"

func ExampleCelsius_Fahrenheit() {
	var c Celsius = 40
	fmt.Println(c.Fahrenheit())
}
