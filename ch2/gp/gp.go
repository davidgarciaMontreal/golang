package main

import (
	"bufio"
	"fmt"
	"github.com/davidgarciaMontreal/golang/ch2/tempconv"
	"os"
	"strconv"
)

func procString(s string) {
	x, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
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
		// TODO: not sure if it can iterate over input.Scan and os.Args in procString...
		for input.Scan() {
			procString(input.Text())
		}
	} else {
		for _, i := range os.Args[1:] {
			procString(i)
		}
	}
}
