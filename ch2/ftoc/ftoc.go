package main

import "fmt"

const tempS = "\u00B0"

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g%sF = %g%sC\n", freezingF, tempS, fToC(freezingF), tempS)
	fmt.Printf("%g%sF = %g%sC\n", boilingF, tempS, fToC(boilingF), tempS)
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
