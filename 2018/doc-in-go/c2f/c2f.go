package c2f

// Celsius represents a temperature value
type Celsius float64

// Fahrenheit returns the temperature in fahrenheit
func (c Celsius) Fahrenheit() float64 {
	return float64(c*1.8 + 32.0)
}

// New creates a new Celsius with a given value
func New(temp float64) Celsius {
	return Celsius(temp)
}
