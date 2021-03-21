package main // Go scripts always start with this line, either main or something else

import "fmt"     // "format" - always import this to printing out
import "strconv" // "string converter" - we will need this if we want to convert int/float to string with strconv.Itoa() (Integer to ascii)

func basicVariables() {

	// Printing Hello, World!
	fmt.Println("Hello, World!")

	// How to declare variables
	// "long" method:
	var varsA int
	varsA = 42
	fmt.Println(varsA)

	// "medium" method:
	var varsB int = 43
	fmt.Println(varsB)

	// "short" method:
	varsC := 44
	fmt.Println(varsC)

	// "group" method:
	var (
		varsFoo       = "1"
		varsBar       = "2"
		varsSomething = "else"
	)

	fmt.Println(varsFoo, varsBar, varsSomething)

	// Printf (f-strings). %v for "value", %T for "type":
	fmt.Printf("%v, %T\n\n", varsC, varsC)

	// Converting types
	var varsD int = 1337
	fmt.Printf("%v, %T\n", varsD, varsD)

	var varsE = float64(varsD)
	fmt.Printf("%v, %T\n", varsE, varsE)

	// int -> float checks, but float -> int you scrap the decimal places and thus lose information
	// int -> string, we will need to import the library "strconv":

	var varsF = strconv.Itoa(varsD)
	fmt.Printf("%v, %T\n\n\n", varsF, varsF)

	// 	Signed Integers
	// 	Left blank, compiler will choose the adecquate one
	// 	Should we want to choose it ourselves, here's the min/max values:
	// 	int8  -> from -128 up to 128
	// 	int16 -> from -32,768 up to 32,767
	// 	int32 -> from -2,147,483,648 up to 2,147,483,647
	// 	int64 -> from -9,223,372,036,854,775,808 up to 9,223,372,036,854,775,807
	// 	for larger numbers, we will need to use the math library

	//  Variables Summary:
	//
	//  Variable declaration:
	//   - var foo int
	//   - var foo int = 42
	//   - foo := 42
	//
	//  Can't redeclare variables, but can shadow them
	//
	//  All variables must be used
	//
	//  Visibility:
	//   - lowercase first letter for package scope
	//   - uppercase first letter to export
	//   - no private scope
	//
	//  Naming conventions:
	//   - Pascal or camelCase
	//       > Capitalize acronyms (HTTP, URL)
	//   - As short as reasonable
	//       > Longer names for longer lives
	//
	//  Type conversions:
	//   - destinationType(variable)
	//   - use "strconv" library for strings

}

func primitives() {

	// Booleans - true or false
	var primsL bool = true
	primsN := 1 == 1
	primsM := 1 == 2

	fmt.Printf("%v, %T\n", primsL, primsL)
	fmt.Printf("%v, %T\n", primsN, primsN)
	fmt.Printf("%v, %T\n\n", primsM, primsM)

	// Bitwise ops
	primsA := 10 // 1010
	primsB := 3  // 0011

	fmt.Println(primsA & primsB)  // 0010
	fmt.Println(primsA | primsB)  // 1011
	fmt.Println(primsA ^ primsB)  // 1001
	fmt.Println(primsA &^ primsB) // 1000 (called bitclear or AND NOT)

	// Bit shifting
	primsX := 8                  //  2^3
	fmt.Println("\n", primsX<<3) //  2^3 * 2^3 = 2^6
	fmt.Println(primsX >> 3)     //  2^3 / 2^3 = 2^0

	// Floating point literals

	// Floating point literals
	// float32 -> from +-1.18e-38 up to +-3.4e38
	// float64 -> from +-2.23e-308 up to +-1.8e308

	// var y float32 = 3.14
	// y = 13.7e72 <-- doesn't work with float32
	var primsY float64 = 13.7e72
	primsZ := 2.1e14
	fmt.Printf("\n%v, %T\n", primsY, primsY)
	fmt.Printf("%v, %T\n", primsZ, primsZ)

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

	var primsComplex complex64 = 1 + 2i
	fmt.Printf("%v, %T\n", primsComplex, primsComplex)
	fmt.Printf("%v, %T\n", real(primsComplex), real(primsComplex))
	fmt.Printf("%v, %T\n", imag(primsComplex), imag(primsComplex))

	// Putting together a real number and an imaginary number to make a complex number
	var primsT complex128 = complex(5, 12)
	fmt.Printf("%v, %T\n", primsT, primsT) // output = (5+12i), complex128

	// Strings
	// In Go, strings are aliases for UTF-8 encoded bytes
	primsS := "this is a string"
	//s[2] = "u" // error, cannot use "u" as type byte in assignment
	primsStr := []byte(primsS)                               // will print all bytes corresponding to each of the letters
	fmt.Printf("%v, %T\n", primsStr[2], primsStr[2])         // output = 105, uint8
	fmt.Printf("%v, %T\n", string(primsStr[2]), primsStr[2]) // output = i, uint8

	// Runes
	// Like strings but UTF-32 instead, alias for int32
	// They must be enclosed between single quotes ('')
	var primsRune rune = 'a'
	fmt.Printf("%v, %T\n", primsRune, primsRune) // output = 97, int32
	// For reading off of a rune, watch https://golang.org/pkg/strings/#Reader.ReadRune

	// Primitives summary
	//
	// - Boolean type:
	//		> Values are true or false
	//		> Not an alias for other types (e.g. int)
	//		> Zero value is false
	//
	// - Numeric type:
	// 		> Integers:
	//			* Signed integers:
	//				¬ Positive or negative
	//				¬ int type has varying size, but min 32 bits
	//				¬ More control: 8 bit (int8) through 64 bit (int64)
	// 			* Unsigned integers:
	//				¬ Only positive
	//				¬ 8 bit (uint8) through 32 bit (uint32)
	// 			* Arithmetic operations
	//				¬ add, sub, multi, div, mod
	//			* Bitwise ops:
	//				¬ AND, OR, XOR, NOT, AND NOT
	//			* Zero value is 0
	//			* Can't mix types in same family! (uint16 + uint32 = error)
	//
	//		> Floating point numbers:
	//			* Follow IEEE-754 standard
	//			* Zero value is 0
	//			* 32 and 64 bit versions
	// 			* Literal styles:
	//				¬ Decimal (3.14)
	//				¬ Exponential (13e18 or 2E10)
	// 				¬ Mixed (13.7e12)
	//			* Arithmetic ops
	//				¬ Add, sub, multi, div
	//		> Complex numbers
	//			* Zero value is (0+0i)
	//			* 64 and 128 bit versions
	//			* Built-in functions:
	//				¬ complex - make complex number from two floats
	//				¬ real - get real part as float
	//				¬ imag - get imaginary part as float
	//			* Arithmetic ops:
	//				¬ Add, sub, multi, div
	//
	//		> Text types
	//			* Strings
	//				¬ UTF-8
	//				¬ Immutable
	//				¬ Can be concatenated with plus (+) operator
	//				¬ Can be converted to []byte
	//			* Runes
	//				¬ UTF-32
	//				¬ Alias for int32
	//				¬ Special methods normally required to process, e.g. strings.Reader#ReadRune
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

	// Constants theory summary:
	//
	// - Immutable, but can be shadowed
	// - Replaced by the compiler at compile time, so value must be calculable at compile time
	// - Named like variables:
	// 		> PascalCase for exported constants (outside of the package)
	// 		> camelCase for internal constants
	// - Typed constants work like immutable variables, so:
	// 		> Can interoperate only with same type
	// - Untyped constants work like literals, so:
	// 		> Can interoperate only with similar types
	//
	// - Enumerated constants:
	// 		> Special symbol iota allows related constants to be created easily
	//		> Iota starts at 0 in each const block and increments by one
	//		> Watch out for constant values that match zero values for variables, may cause subtle bugs
	//
	// - Enumerated expressions:
	// 		> Operations that can be determined at compile time are allowed:
	//			* Arithmetic
	//			* Bitwise operations
	//			* Bitshifting

}

func main() {
	basicVariables()
	primitives()
	constants()
}
