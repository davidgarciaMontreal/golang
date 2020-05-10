package main

import (
	"bufio"
	"fmt"
	"github.com/davidgarciaMontreal/golang/ch2/tempconv"
	"os"
	"strconv"
)

func procString(x float64) {
	m := tempconv.Meter(x)
	f := tempconv.Feet(x)
	k := tempconv.Kilogram(x)
	p := tempconv.Pound(x)
	fmt.Printf("%s = %s\n", m, tempconv.MToF(m))
	fmt.Printf("%s = %s\n", f, tempconv.FToM(f))
	fmt.Printf("%s = %s\n", k, tempconv.KiToP(k))
	fmt.Printf("%s = %s\n", p, tempconv.PToKi(p))
}
func main() {
	if n := len(os.Args[1:]); n == 0 {
		input := bufio.NewScanner(os.Stdin)
		fmt.Println("Zero Arguments need to read from STDIN")
		for input.Scan() {
			x := input.Text()
			f, err := strconv.ParseFloat(x, 64)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			procString(f)
		}
	} else {
		for _, i := range os.Args[1:] {
			f, err := strconv.ParseFloat(i, 64)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			procString(f)
		}
	}
}
