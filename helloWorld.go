package main // Go scripts always start with this line, either main or something else

import "fmt"     // "format" - always import this to printing out
import "strconv" // "string converter" - we will need this if we want to convert int/float to string with strconv.Itoa() (Integer to ascii)

func basicVariables() {

	// Printing Hello, World!
	fmt.Println("Hello, World!")

	// How to declare variables
	// "long" method:
	var i int
	i = 42
	fmt.Println(i)

	// "medium" method:
	var j int = 43
	fmt.Println(j)

	// "short" method:
	k := 44
	fmt.Println(k)

	// "group" method:
	var (
		foo       = "1"
		bar       = "2"
		something = "else"
	)

	fmt.Println(foo, bar, something)

	// Printf (f-strings). %v for "value", %T for "type":
	fmt.Printf("%v, %T\n\n", k, k)

	// Converting types
	var a int = 1337
	fmt.Printf("%v, %T\n", a, a)

	var b = float64(a)
	fmt.Printf("%v, %T\n", b, b)

	// int -> float checks, but float -> int you scrap the decimal places and thus lose information
	// int -> string, we will need to import the library "strconv":

	var c = strconv.Itoa(a)
	fmt.Printf("%v, %T\n\n\n", c, c)

	//	  Signed Integers
	//	  Left blank, compiler will choose the adecquate one
	//	  Should we want to choose it ourselves, here's the min/max values:
	//	  int8  -> from -128 up to 128
	//	  int16 -> from -32,768 up to 32,767
	//	  int32 -> from -2,147,483,648 up to 2,147,483,647
	//	  int64 -> from -9,223,372,036,854,775,808 up to 9,223,372,036,854,775,807
	//	  for larger numbers, we will need to use the math library

	//   Variables Summary:
	//
	//   Variable declaration:
	//    - var foo int
	//    - var foo int = 42
	//    - foo := 42
	//
	//   Can't redeclare variables, but can shadow them
	//
	//   All variables must be used
	//
	//   Visibility:
	//    - lowercase first letter for package scope
	//    - uppercase first letter to export
	//    - no private scope
	//
	//   Naming conventions:
	//    - Pascal or camelCase
	//        > Capitalize acronyms (HTTP, URL)
	//    - As short as reasonable
	//        > Longer names for longer lives
	//
	//   Type conversions:
	//    - destinationType(variable)
	//    - use "strconv" library for strings

}

func primitives() {

	// Booleans - true or false
	var l bool = true
	n := 1 == 1
	m := 1 == 2

	fmt.Printf("%v, %T\n", l, l)
	fmt.Printf("%v, %T\n", n, n)
	fmt.Printf("%v, %T", m, m)

	// Bitwise ops
	a := 10             // 1010
	b := 3              // 0011
	fmt.Println(a & b)  // 0010
	fmt.Println(a | b)  // 1011
	fmt.Println(a ^ b)  // 1001
	fmt.Println(a % ^b) // 0100

	// Bit shifting
	a := 8              // 2^3
	fmt.Println(a << 3) //  2^3 * 2^3 = 2^6
	fmt.Println(a >> 3) //  2^3 / 2^3 = 2^0

	// Floating point literals

	// Floating point literals
	// float32 -> from +-1.18e-38 up to +-3.4e38
	// float64 -> from +-2.23e-308 up to +-1.8e308

	var n float32 = 3.14
	// n = 13.7e72 <-- doesn't work with float32
	var n float64 = 13.7e72
	n = 2.1e14
	fmt.Printf("%v, %T\n", n, n)

	// Reminder: float32 and float64 don't compute, both must be same type.
	//           no modulus              |
	//           no bit shifting         |      only available for integers
	//           no bitwise operations   |

	// Complex type (imaginary numbers)
	// Very powerful for data science
	// 2 types available:
	// complex128
	// complex64
	// Same operations available as float32 & 64

	var n complex64 = 1 + 2i
	fmt.Printf("%v, %T\n", n, n)
	fmt.Printf("%v, %T\n", real(n), real(n))
	fmt.Printf("%v, %T\n", imag(n), imag(n))

	// Putting together a real number and an imaginary number to make a complex number
	var n complex128 = complex(5, 12)
	fmt.Printf("%v, %T\n", n, n) // output = (5+12i), complex128

	// Strings
	// In Go, strings are aliases for UTF-8 encoded bytes
	s := "this is a string"
	//s[2] = "u" // error, cannot use "u" as type byte in assignment
	b = []byte(s)                               // will print all bytes corresponding to each of the letters
	fmt.Println("%v, %T\n", s[2], s[2])         // output = 105, uint8
	fmt.Println("%v, %T\n", string(s[2]), s[2]) // output = i, uint8

	// Runes
	// Like strings but UTF-32 instead, alias for int32
	// They must be enclosed between single quotes ('')
	var r rune = 'a'
	fmt.Printf("%v, %T\n", r, r) // output = 97, int32
	// For reading off of a rune, watch https://golang.org/pkg/strings/#Reader.ReadRune
}

func main() {
	basicVariables()
	primitives()
}
