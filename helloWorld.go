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
	x := 8              //  2^3
	fmt.Println(x << 3) //  2^3 * 2^3 = 2^6
	fmt.Println(x >> 3) //  2^3 / 2^3 = 2^0

	// Floating point literals

	// Floating point literals
	// float32 -> from +-1.18e-38 up to +-3.4e38
	// float64 -> from +-2.23e-308 up to +-1.8e308

	// var y float32 = 3.14
	// y = 13.7e72 <-- doesn't work with float32
	var y float64 = 13.7e72
	z := 2.1e14
	fmt.Printf("%v, %T\n", y, y)
	fmt.Printf("%v, %T\n", z, z)

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

	var u complex64 = 1 + 2i
	fmt.Printf("%v, %T\n", u, u)
	fmt.Printf("%v, %T\n", real(u), real(u))
	fmt.Printf("%v, %T\n", imag(u), imag(u))

	// Putting together a real number and an imaginary number to make a complex number
	var t complex128 = complex(5, 12)
	fmt.Printf("%v, %T\n", t, t) // output = (5+12i), complex128

	// Strings
	// In Go, strings are aliases for UTF-8 encoded bytes
	s := "this is a string"
	//s[2] = "u" // error, cannot use "u" as type byte in assignment
	f := []byte(s)                             // will print all bytes corresponding to each of the letters
	fmt.Printf("%v, %T\n", f[2], f[2])         // output = 105, uint8
	fmt.Printf("%v, %T\n", string(f[2]), f[2]) // output = i, uint8

	// Runes
	// Like strings but UTF-32 instead, alias for int32
	// They must be enclosed between single quotes ('')
	var r rune = 'a'
	fmt.Printf("%v, %T\n", r, r) // output = 97, int32
	// For reading off of a rune, watch https://golang.org/pkg/strings/#Reader.ReadRune
}

func constants() {

	// Constants:

	// Very similar concept than vars, they can be shadowed and acquire same types however they are immutable
	// Interesting feature is IOTA:
	// 	  - Naming constants and assigning iota to the first constant in a constants block
	// 		keeps track of the occurrences of every each one of them like a counter but we
	//		can assign a special value to that counter. Examples:

	// iota example 1:
	const (
		_ = iota + 5
		catSpecialist
		dogSpecialist
		snakeSpecialist
	)

	fmt.Printf("%v\n", snakeSpecialist) // output = 8

	// iota example 2:
	type ByteSize float64
	const (
		_  = iota // ignore first value by assigning to blank identifier
		KB = 1 << (10 * iota)
		MB
		GB
		TB
		PB
		EB
		ZB
		YB
	)

	fileSize := 4000000000.
	fmt.Printf("%.2fGB\n", fileSize/GB)

	// Setting boolean flags inside of a single byte with bit shifting and iota

	const (
		isAdmin = 1 << iota
		isHeadquarters
		canSeeFinancials

		canSeeAfrica
		canSeeAsia
		canSeeEurope
		canSeeNorthAmerica
		canSeeSouthAmerica
	)

	var roles byte = isAdmin | canSeeFinancials | canSeeAsia | canSeeSouthAmerica
	fmt.Printf("%b\n", roles)
	fmt.Printf("Is Admin? %v\n", isAdmin&roles == isAdmin)
	fmt.Printf("Is Headquarters? %v\n", isHeadquarters&roles == isHeadquarters)
	fmt.Printf("Can see Europe? %v\n", canSeeEurope&roles == canSeeEurope)

}

func main() {
	basicVariables()
	primitives()
	constants()
}
