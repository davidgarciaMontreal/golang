package main

import "fmt"

const boilingF = 212.0
const tempS = "\u00B0"

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point =%g%sF or %g%sC\n", f, tempS, c, tempS)
}
