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

	//	  Signed Integers
	//	  Left blank, compiler will choose the adecquate one
	//	  Should we want to choose it ourselves, here's the min/max values:
	//	  int8  -> from -128 up to 128
	//	  int16 -> from -32,768 up to 32,767
	//	  int32 -> from -2,147,483,648 up to 2,147,483,647
	//	  int64 -> from -9,223,372,036,854,775,808 up to 9,223,372,036,854,775,807
	//	  for larger numbers, we will need to use the math library

	//

}

func main() {
	basicVariables()
	primitives()
}
