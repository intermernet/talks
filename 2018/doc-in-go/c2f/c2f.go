package c2f

type Celsius float64

func (c Celsius) Fahrenheit() float64 {
	return float64(c*1.8 + 32.0)
}
