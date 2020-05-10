package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64
type Feet float64
type Meter float64
type Kilogram float64
type Pound float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string {
	return fmt.Sprintf("%.3g°C", c)
}
func (f Fahrenheit) String() string {
	return fmt.Sprintf("%.3g°F", f)
}
func (k Kelvin) String() string {
	return fmt.Sprintf("%.3g°K", k)
}
func (f Feet) String() string {
	return fmt.Sprintf("%.4gft", f)
}
func (m Meter) String() string {
	return fmt.Sprintf("%.4gm", m)
}
func (k Kilogram) String() string {
	return fmt.Sprintf("%.4gkg", k)
}
func (p Pound) String() string {
	return fmt.Sprintf("%.4glb", p)
}
