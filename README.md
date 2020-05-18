# golang
## 3.1 Integers
List of different types of **integer types**. Signed numbers are represented in 2's-complement form.

* Eight sizes of signed/unsigned integers: 8, 16, 32, 64 bits.
* Two more for int and uint (can be iether 32 or 64 bits).
* rune synonym for int32; it indicates a value is a Unicode codepoint.
* byte synonym for uint8; it emphasizes the value is a piece of raw data.
* unsigned integer type uintptr; it is used only for low-level programming.
* binary operators 19 in total:Arithmetic, Comparison,Unary prefix, Bitwise binary, % only applies to integers. -5%3 == -5%-3 = -2.
* *overflow* of integer types: higher order bits are silently discarded.
* **Note** even if the sizes are the same one must perform a type conversion.

See `go doc fmt` for more info.
```
o := 0666
fmt.Printf("%d %[1]o %#[1]o\n", o) // "438 666 0666"
```

## 3.2 Floating-Point Numbers
There are only two sizes of *floating-point numbers*: `float32` and `float64`  whose arithmentic proproperties are governed by the [IEEE 754](https://en.wikipedia.org/wiki/IEEE_754) standard.
* Limits are found in the `math package`: `math.MaxFloat32`, `math.MaxFloat64`.
* Six decimal digits of precision and fifteen decimal digits of precision for float32 and float64 respectively.
* The smallest positive integer cannot be represented as float32 is not large. 16777216 == 2^24 == 1 << 24. See good explanation [here](https://stackoverflow.com/questions/12596695/why-does-a-float-variable-stop-incrementing-at-16777216-in-c?noredirect=1). TL;DR: highest integer that can be converted as a float32 is `16777216` any value higher than this integer, a float32 will not be be able to represent it.Try `float32(16777216) == float32(16777216 + n)` where n = 1,2,.. For any n, it will be `true`.

# General Notes
## [Code point](https://en.wikipedia.org/wiki/Code_point)
A **code point** or **code position** is any of the numerical values that make up the code space.

Many code points represent single characters but they can also have other meanings, such as for formatting.

The Unicode code space is divided into seventeen planes (the basic multilingual plane, and 16 supplementary planes)
* Code unit (UCS-4 encoding) any code point is encoded as 4-byte (octet) binary numbers
* In the UTF-8 encoding, different code points are encoded as sequences from one to four bytes long, forming a self-synchronizing code.
