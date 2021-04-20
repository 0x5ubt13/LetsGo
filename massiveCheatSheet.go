package main // Go scripts always start with this line, either main or something else

import (
	"fmt" // "format" - always import this to printing out
	//"go/ast"
	"io/ioutil" // "input outout utility" - to read the body of a request
	"log"       // to log errors
	"math"
	"net/http" // to make http Get requests
	"reflect"  // "reflections" - for extracting tags out of a field when dealing with structs
	"strconv"  // "string converter" - we will need this if we want to convert int/float to string with strconv.Itoa() (Integer to ascii)
)

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

	fmt.Print("\n")
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

	fmt.Print("\n")
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

func arraysAndSlices() {

	// Arrays
	// Declare first the number of indexes, then the type, then the values

	// Integer arrays
	grades := [3]int{97, 85, 93}
	fmt.Printf("Grades: %v \n", grades)

	gradesLiteral := [...]int{97, 85, 93}
	fmt.Printf("Grades: %v", gradesLiteral)

	fmt.Print("\n\n")

	// Strings arrays
	// Note: arrays' indexes look like python's lists: first index is index [0]
	var students [5]string
	fmt.Printf("Students: %v\n", students)
	students[0] = "Lisa"
	fmt.Printf("Students: %v\n", students)
	students[2] = "Emanuel"
	students[1] = "Gabriel"
	fmt.Printf("Students: %v\n", students)
	fmt.Printf("Student #1: %v\n", students[1])

	fmt.Printf("Students: %v\n", len(students))

	fmt.Print("\n")
	// Array arrays
	var identityMatrix [3][3]int = [3][3]int{[3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}}
	fmt.Printf("%v\n%v\n%v\n", identityMatrix[0], identityMatrix[1], identityMatrix[2])

	var identityMatrixTwo [3][3]int
	identityMatrixTwo[0] = [3]int{0, 0, 1}
	identityMatrixTwo[1] = [3]int{0, 1, 0}
	identityMatrixTwo[2] = [3]int{1, 0, 0}
	fmt.Println(identityMatrixTwo)

	fmt.Print("\n")
	// In Go, when declaring a new variable pointing to an existing list, it doesn't get referenced but
	// it gets fully copied instead (slice equivalent in python = list[:])
	arrayA := [...]int{1, 2, 3}
	arrayB := arrayA
	arrayB[1] = 11
	fmt.Println(arrayA) // output = [1 2 3]
	fmt.Println(arrayB) // output = [1 11 3]

	// In the other hand, if we want for the new variable to point to the list, we use the ampersand
	fmt.Print("\n")
	arrayC := [...]int{1, 2, 3}
	arrayD := &arrayC
	arrayD[1] = 11
	fmt.Println(arrayC)       // output = [1 11 3]
	fmt.Print(arrayD, "\n\n") // output = &[1 11 3]

	// Slices
	// When you don't want to make a fixed size array, use slices
	// They look like arrays and behave very similarly, with certain exceptions
	sliceA := []int{1, 2, 3}
	sliceB := sliceA // Now we don't need the ampersand, sliceB will be pointed to sliceA just like in Python's lists
	sliceB[1] = 22
	fmt.Println(sliceA)                         // output = [1 22 3]
	fmt.Printf("Length: %v\n", len(sliceA))     // output = 3
	fmt.Printf("Capacity: %v\n\n", cap(sliceA)) // output = 3

	// PoC slice slices
	sliceC := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sliceD := sliceC[:]   // slice of all elements
	sliceE := sliceC[3:]  // slice of all elements
	sliceF := sliceC[:6]  // slice first 6 elements
	sliceG := sliceC[3:6] // slice the 4th, 5th, and 6th elements
	sliceD[4] = 999
	fmt.Println(sliceC)
	fmt.Println(sliceD)
	fmt.Println(sliceE)
	fmt.Println(sliceF)
	fmt.Println(sliceG, "\n")

	// Slicing works the same way with an array. Outputs from both above and below snippets are the same
	arrayF := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sliceH := arrayF[:]   // slice of all elements
	sliceI := sliceH[3:]  // slice from 4th element to end
	sliceJ := sliceH[:6]  // slice first 6 elements
	sliceK := sliceH[3:6] // slice the 4th, 5th, and 6th elements
	sliceH[4] = 777
	fmt.Println(arrayF)
	fmt.Println(sliceH)
	fmt.Println(sliceI)
	fmt.Println(sliceJ)
	fmt.Println(sliceK)
	fmt.Print("\n")

	// Creating a slice with the make() function:
	// 1st arg = type of object, 2nd arg = length, 3rd arg (opt) = capacity
	makingA := make([]int, 3, 100)
	fmt.Println(makingA)
	fmt.Printf("Length: %v\n", len(makingA))
	fmt.Printf("Capacity: %v\n", cap(makingA))

	// Appending elements to a slice
	appending := []int{}
	fmt.Println(appending)
	fmt.Printf("Length: %v\n", len(appending))
	fmt.Printf("Capacity: %v\n", cap(appending))
	appending = append(appending, 1)
	fmt.Println(appending)
	fmt.Printf("Length: %v\n", len(appending))
	fmt.Printf("Capacity: %v\n", cap(appending))
	appending = append(appending, 2, 3, 4, 5)
	fmt.Println(appending)
	fmt.Printf("Length: %v\n", len(appending))
	fmt.Printf("Capacity: %v\n\n", cap(appending))

	// Concatenating slices:
	// Use the variadic function operator (called spread operator in JavaScript) (...)
	//appending = append(appending, []int{2, 3, 4, 5}) // This will throw an error
	appending = append(appending, []int{2, 3, 4, 5}...) // This will work perfectly
	fmt.Println(appending)
	fmt.Printf("Length: %v\n", len(appending))
	fmt.Printf("Capacity: %v\n\n", cap(appending))

	// Stack operations (append, pop)
	// Just do it with more slices! And remember to use the variadic operator
	stackOps := []int{1, 2, 3, 4, 5}
	fmt.Println(stackOps)                              // Output = [1 2 3 4 5]
	stackOps2 := append(stackOps[:2], stackOps[3:]...) // remember the 3 dots or it won't work
	fmt.Println(stackOps2)                             // Output = [1 2 4 5]
	fmt.Println(stackOps)                              // Output = [1 2 4 5 5] (???) (make sure you don't reference the first slice after this operation)

	// Arrays and Slices summary:
	//
	// - Arrays:
	// 		> Collection of items with same type
	//		> Fixed size
	//		> Declaration styles:
	//			* a := [3]int{1, 2, 3}
	//			* a := [...]int{1, 2, 3}
	//			* var a [3]int
	//		> Access via zero-based index
	//			* a := [3]int{1, 3, 5} // a[1] == 3
	//		> len() function returns size of array
	// 		> Copies refer to different underlying data (unexpected behaviour)
	//
	// - Slices:
	// 		> Backed by arrays
	// 		> Creation styles
	//			* Literal style
	//			* Via make function:
	//				· a := make([]int, 10)
	//				· a := make([]int, 10, 100)
	//		> len() func returns length of slice
	//		> cap() func returns capacity of slice
	//		> append() func to add elements to slice
	//			* May cause expensive copy operation if underlying array is too small
	//		> Copies refer to same underlying array

}

func mapsAndStructs() {

	// Maps
	// key:value pairs, always with the same type as first declared

	// To create a map for later populating it, use the make() function:
	statePopulations := make(map[string]int) // "Make a map of strings as keys with integers as values
	statePopulations = map[string]int{
		"California": 39250017,
		"Texas":      27862596,
		"Florida":    20612439,
		"New York":   19745289,
		"Ohio":       11614373,
	}
	//map1 := map[[]int]string{} // This won't work. Slice isn't a valid type for this
	//testmap1 := map[[3]int]string{} // "Create a map of an array of integers over to strings"
	map2 := statePopulations
	delete(map2, "Ohio") // it also gets deleted from statePopulations

	// Checking whether a key in map2 exists with "comma okay" (, ok)
	_, ok := map2["California"]
	fmt.Println(ok) // Returns true

	// Adding a new key:value pair
	statePopulations["Pennsylvania"] = 12801539 // it also gets declared in map2
	fmt.Println(statePopulations)
	fmt.Println(map2)
	fmt.Println(map2["Ohio"])                // If key doesn't exist, a zero value returns
	fmt.Println(len(statePopulations), "\n") // Returns length of map (5)

	// Structs
	// Like maps but we can use the data type we want within them

	type Doctor struct {
		number     int      // number will be an integer
		actorName  string   // actorName will be a string
		companions []string // companions will be a slice of strings
	}

	aDoctor := Doctor{
		number:    3,
		actorName: "Jon Pertwee",
		companions: []string{
			"Liz Shaw",
			"Jo Grant",
			"Sarah Jane Smith",
		},
	}
	fmt.Println(aDoctor)

	// Creating anonymous struct for organising a one time only subset of data:
	oneTime := struct{ name string }{name: aDoctor.actorName}
	anotherTime := oneTime // if := &oneTime, Tom Baker would apply in both
	anotherTime.name = "Tom Baker"
	fmt.Println(oneTime)
	fmt.Println(anotherTime, "\n")

	// Embedding (pseudo-inheritance)
	type Animal struct {
		Name   string
		Origin string
	}

	type Bird struct {
		Animal   // We embed the abovementioned struct. Now we *have* Animal
		SpeekKPH float32
		CanFly   bool
	}

	embedding := Bird{}
	embedding.Name = "Emu"
	embedding.Origin = "Australia"
	embedding.SpeekKPH = 48
	embedding.CanFly = false
	fmt.Println(embedding)
	fmt.Println("Name:", embedding.Name, "\n")

	// Tags in embeds
	type NewAnimal struct {
		Name   string `required max:"100"`
		Origin string
	}

	// Import reflection package (reflect) to check Tags
	embedType := reflect.TypeOf(NewAnimal{})
	field, _ := embedType.FieldByName("Name")
	fmt.Println(field.Tag, "\n\n")

	// Summary:
	//
	// - Maps
	// 		> Collections of value types that are accessed via keys
	//		> Created via literals or make function
	//		> Members accessed via [key] syntax
	//			* myMap["key"] = "value"
	//		> Check for presence with "value, ok" form of result
	//		> Multiple assignments refer to same underlying data
	//
	//
	// - Structs
	//		> Collection of disparate data types that describe a single concept
	//		> Keyed by named fields
	// 		> Normally created as types, but anonymous structs are allowed
	// 		> Structs are value types
	//		> No inheritance, but can use composition via embedding
	//		> Tags can be added to struct fields to describe field
}

func controlFlow() {
	// If - else if - else statements

	if true {
		fmt.Println("The test is true\n")
	}

	// Illustrating a control flow with previous code
	cfStatePopulations := map[string]int{
		"California": 39250017,
		"New York":   19745289,
		"Ohio":       11614373,
	}
	// if pop "comma okay" := (the test); is okay {go ahead and do this}
	// also, pop var will only exist within that if statement unless previously declared
	if cfPopulation, ok := cfStatePopulations["California"]; ok {
		fmt.Println(cfPopulation)
	}

	// Guessing a number game
	cfNumber := 50
	cfGuess := 70
	if cfGuess < cfNumber {
		fmt.Println("Too low")
	}
	if cfGuess > cfNumber {
		fmt.Println("Too high")
	}
	if cfGuess == cfNumber {
		fmt.Println("You got it!")
	}
	fmt.Println(cfNumber <= cfGuess, cfNumber >= cfGuess, cfNumber != cfGuess, "\n") // true false true

	// Adding logical operators
	// && = and
	// || = or
	// ! = not
	cfNumber2 := 40
	cfGuess2 := 39
	if cfGuess2 < 1 {
		fmt.Println("The guess must be greater than 1!")
	} else if cfGuess2 > 100 {
		fmt.Println("The guess must be less than 100!")
	} else {
		if cfGuess2 < cfNumber2 {
			fmt.Println("Too low")
		}
		if cfGuess2 > cfNumber2 {
			fmt.Println("Too high")
		}
		if cfGuess2 == cfNumber2 {
			fmt.Println("You got it!")
		}
		fmt.Println(cfNumber2 <= cfGuess2, cfNumber2 >= cfGuess2, cfNumber2 != cfGuess2, "\n") // false true true
	}

	// Short-circuiting in Go:
	// If you have an statement with 2 or more operators and the 1st one is enough to make the rest of the condition,
	// the rest are not going to be evaluated (in OR, the 1st being true; in AND, the 1st being false, etc)
	cfGuess3 := -5
	if cfGuess3 < 1 || returnTrue() || cfGuess3 > 100 { // it reads up to the first condition so returnTrue() never executes
		fmt.Println("The guess must be between 1 and 100!")
	}

	// Testing equality operators, problem with floating point numbers
	myNum := 0.1 // This will pass the if statement
	//myNum := 0.123456789  // This will not pass the if statement
	if myNum == math.Pow(math.Sqrt(myNum), 2) {
		fmt.Println("These are the same")
	} else {
		fmt.Println("These are different")
	}

	// For the float to pass the test, we need to cheat:
	myNum2 := 0.123456789
	if math.Abs(myNum2/math.Pow(math.Sqrt(myNum2), 2)-1) < 0.001 {
		fmt.Println("These are the same")
	} else {
		fmt.Println("These are different")
	}

	// Switch statements
	// switch "num" {case 1:, case 2:, default:}
	switch 2 {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("neither one nor two")
	} // output = "two"

	// In other languages, to compare against multiple cases they use "follow through". In Go,
	// this is explicitly forbidden. If you want this feature you must declare "fall through". What
	// we have though, is multiple testing in a single case, but they must be unique (non overlapping):
	switch i := 2 + 3; i { // = 5
	case 1, 5, 10:
		fmt.Println("one, five or ten")
	case 2, 4, 6:
		fmt.Println("two, four or six")
	default:
		fmt.Println("another number")
	} // output = "one, five or ten"

	// Switches with boolean conditions:
	// output = executes the first case because the second one overlaps, but triggers later so gets ignored.
	// this means "break" keyword is implied by design, as abovementioned with the fall through issue
	boolSwitch := 10
	switch {
	case boolSwitch <= 10:
		fmt.Println("less than or equal to ten")
	case boolSwitch >= 20:
		fmt.Println("less than or equal to twenty")
	default:
		fmt.Println("greater than twenty")
	}

	// Type switches
	var typeSwitch interface{} = 1
	switch typeSwitch.(type) {
	case int:
		fmt.Println("typeSwitch is an int")
		break
		fmt.Println("This will not get printed")
	case float64:
		fmt.Println("typeSwitch is a float64")
	case string:
		fmt.Println("typeSwitch is a string")
	default:
		fmt.Println("typeSwitch is another type")
	}

	// Control flow summary
	//
	// - If statements
	// 		> Initializer
	//		> Comparison operators (<, >, ==, >=, <=, !=)
	//		> Logical operators(&&, ||, !)
	//		> Short circuiting
	//		> If - else statements
	//		> If - else if statements
	//		> If - else if - else statements
	//		> Equality and floats (cheat using comparison against an error function)
	//
	// - Switch statements
	//		> Switching on a tag (variable)
	//		> Cases with multiple tests
	//		> Initializers generating tags
	//		> Switches with no tag
	// 		> Fallthrough (implicit breaks, explicit fallthrough)
	//		> Type switches
	//		> Breaking out early

}

func returnTrue() bool {
	// Declaring this function to showcase short-circuiting in control flow statements. See line 626.

	fmt.Println("Returning True")
	return true
}

func loops() {

	// For loops:

	// initialise; condition; increment by {something to loop}
	for loopOne := 0; loopOne < 5; loopOne++ {
		fmt.Println(loopOne)
	}

	fmt.Println("\n")

	// Incrementing by other than 1:
	for loopTwo := 0; loopTwo < 5; loopTwo = loopTwo + 2 {
		fmt.Println(loopTwo)
	}

	fmt.Println("\n")

	// Looping two variables at the same time
	for loopThree, loopFour := 0, 0; loopThree < 5; loopThree, loopFour = loopThree+1, loopFour+2 {
		fmt.Println(loopThree, loopFour)
	}

	fmt.Println("\n")

	// Playing around with the counter (bad practise but valuable trick to know about)
	for loopFive := 0; loopFive < 5; loopFive++ {
		fmt.Println(loopFive)
		if loopFive%2 == 0 {
			loopFive /= 2
		} else {
			loopFive = 2*loopFive + 1
		}
	}

	fmt.Println("\n")

	// Simplifying the initialisation:
	loopSix := 0
	for ; loopSix < 5; loopSix++ {
		fmt.Println(loopSix)
	}

	fmt.Println("\n")

	// More simplifying, variable scoped now only to the For loop.
	for loopSeven := 0; loopSeven < 5; {
		fmt.Println(loopSeven)
		loopSeven++
	}

	fmt.Println()

	// The following is also equivalent to Do ... while loops
	loopEight := 0
	for loopEight < 5 {
		fmt.Println(loopEight)
		loopEight++
	}

	fmt.Println("\n")

	// The following is also equivalent to While True loops
	loopNine := 0
	for {
		fmt.Println(loopNine)
		loopNine++
		if loopNine == 5 {
			break
		}
	}

	// "Continue" statement. Skipping even numbers with continue
	for loopTen := 0; loopTen < 10; loopTen++ {
		if loopTen%2 == 0 {
			continue
		}
		fmt.Println(loopTen)
	}

	fmt.Println("\n")

	// Nested loops, breaking to a tag. (I.e.: tag "Loop")
Loop:
	for loopEleven := 1; loopEleven <= 3; loopEleven++ {
		for loopTwelve := 1; loopTwelve <= 3; loopTwelve++ {
			fmt.Println(loopEleven * loopTwelve)

			if loopEleven*loopTwelve >= 3 {
				break Loop
			}
		}
	}

	// Looping through collections
	sliceLoopOne := []int{1, 2, 3}
	for key, value := range sliceLoopOne {
		fmt.Println("index:", key, "value:", value)
	}
	fmt.Println("\n")

	sliceLoopTwo := "Hello Go!"
	for key, value := range sliceLoopTwo {
		fmt.Println("index:", key, "value:", string(value))
	}

	// Loops summary
	//
	// - For statements
	//		> Simple loops
	//			* for initialiser; test; incrementer {}
	//			* for test {}
	//			* for {}
	//		> Exiting early
	//			* break
	//			* continue
	//			* labels
	//		> Looping over collections
	//			* Arrays, slices, maps, strings, channels
	//			* for k, v := range collection {}
	//				· Only the keys: for k :=
	//				· Only the values: for _, v :=
}

///////////////////////////////
//       Control flow:       //
//  Defer, Panic and Recover //
///////////////////////////////

// Deferring:
// We use this keyword to delay the execution of a function until right before hitting the "return".

func deferOne() {
	fmt.Println("start")
	defer fmt.Println("middle\n")
	fmt.Println("end")
	// output: start, end, middle. Middle gets deferred to the end.
}

func deferTwo() {
	defer fmt.Println("start\n")
	defer fmt.Println("middle")
	defer fmt.Println("end")
	// When deferring multiple statements, they work in LIFO order
	// output: end, middle, start
}

// Deferring is very useful when we have an open object. Deferring the Close() statement will keep it open
// until we don't need to work with it anymore
// Be mindful of deferring a lot of open resources, in that case you'd probably be better off not deferring
func deferThree() {
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
	fmt.Println("\n")
}

// Also, the defer statement evaluates and then gets delayed.
// In this example, the output is "start"
func deferFour() {
	a := "start"
	defer fmt.Println(a)
	a = "end"
}

func panicking() {
	// In Go we have no thrown exceptions, but in some situations we have "panic"
	// Panic One:
	// a, b := 1, 0
	// ans := a / b
	// fmt.Println(ans)
	// Output 1: panic: runtime error: integer divide by zero

	// Panic Two:
	// fmt.Println("start")
	// panic("something bad has happened")
	// fmt.Println("end")
	// Output 2: panic: runtime error: integer divide by zero

	// Panic Three:
	// Setting up a webserver:
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//  	w.Write([]byte("Hello Go!"))
	// })
	// err := http.ListenAndServe(":8080, nil")
	//		if err != nil {
	//			panic(err.Error())
	//		}
	// If we run the application twice, it will return the Panic error

	// Panic Four:
	// fmt.Println("start")
	// defer fmt.Println("this was deferred")
	// panic("something bad happened")
	// fmt.Println("end") // <- this will never get printed
}

func recovering() {
	// Panic Five / Recovery showcase
	// Nesting deferred recovery in anonymous func inside of a nested anonymous func
	fmt.Println("Start")
	func() {
		fmt.Println("About to panic")
		defer func() {
			if err := recover(); err != nil {
				log.Println("Error:", err)
			}
		}() // <- Parenthesis to call the anon func
		panic("something bad happened")
		fmt.Println("done panicking")
	}() // <- Parenthesis to call the anon func
	fmt.Println("end")

	// Summary
	//
	// 	- Defer:
	// 		> Used to delay execution of a statement until function exists
	//		> Useful to group "open" and "close" functions together
	//			* Be careful in loops though
	// 		> Run in LIFO order
	//		> Args evaluated at time defer is executed, not at time of called function execution
	//
	// 	- Panic:
	//		> Occurs when program can't continue at all
	//			* Don't use when file can't be opened, unless it's critical
	//			* Use for unrecoverable events - can't obtain TCP port for web server
	//		> Function will stop executing
	//			* Deferred functions will still fire
	// 		> If nothing handles panic, program will exit
	//
	// 	- Recover:
	//		> Used to recover from panics
	//		> Only useful in deferred functions
	//		> Current function will not attempt to continue, but higher functions in call stack will!
}

func pointers() {

	// Pointers:

	// De-referencing operator
	var pa int = 42
	var pb *int = &pa    // <-- The * here is declaring a pointer to data of the int type
	fmt.Println(pa, *pb) // <-- The * here is de-referencing the memory location to pull the value
	pa = 27
	fmt.Println(pa, *pb)
	*pb = 14
	fmt.Println(pa, *pb, "\n")

	// Pointer arithmetic
	// If we need to do this, we should import the package "unsafe"
	// By design, Go is built in by simplicity, hence pointer arithmetics are not allowed by default
	// pc := [3]int{1, 2, 3}
	// pd := &pc[0]
	// pe := &pc[1] - 4 // <-- This would throw error - invalid operation - mismatched types *int and int.
	// fmt.Printf("%v %p %p\n", pc, pd, pe)

	// Pointers to objects - example 1:
	type pointersMyStruct struct {
		foo int
	}
	var pms *pointersMyStruct
	pms = &pointersMyStruct{foo: 42}
	fmt.Println(pms)        // output: &{42}
	fmt.Println(*pms, "\n") // output: {42}

	// Pointers to objects - example 2:
	type pointersMyStruct2 struct {
		foo int
	}
	var pms2 *pointersMyStruct2
	pms2 = new(pointersMyStruct2)
	pms2.foo = 42
	fmt.Println(pms2.foo, "\n") // output: 42

	// Pointers to slices
	psa := [3]int{1, 2, 3}
	psb := psa            // copy of psa, won't apply changes in psa
	psc := []int{1, 2, 3} // Slices are actually pointers to an underlying array
	psd := psc            // pointed to psc, will apply changes in both
	fmt.Println(psa, psb)
	fmt.Println(psc, psd)
	psa[1] = 42
	psc[1] = 42
	fmt.Println(psa, psb) // output: [1 42 3] [1 2 3]
	fmt.Println(psc, psd) // output: [1 42 3] [1 42 3]

	// Pointers to maps
	pma := map[string]string{"foo": "bar", "baz": "buz"}
	pmb := pma
	fmt.Println(pma, pmb)
	pma["foo"] = "qux"
	fmt.Println(pma, pmb) // Same behaviour than slices, it will change both maps

	// Pointers summary
	//
	//	- Creating pointers
	//		> Pointer types use an asterisk (*) as a prefix to type pointed to
	//			(*) *int - a pointer to an integer
	//		> Use the "addressof" operator (&) to get the memory address of a variable in memory
	//
	//	- Dereferencing pointers
	//		> Dereferece a pointer by calling it with a preceding asterisk (*)
	//		> Complex types (e.g. structs) are automatically dereferenced
	//
	// 	- Create pointers to objects
	//		> Can use the "addressof" operator (&) if value type already exists
	//			* Example:
	//				ms := myStruct{foo: 42}
	//				p  := &ms
	//		> Use addressof operator before initialiser:
	//			* Example:
	//				&myStruct{foo: 42}
	//		> Use the "new" keyword
	//			* Can't initialize fields at the same time though.
	//
	// 	- Types with internal pointers:
	//		> All assignment operations in Go are copy operations
	// 		> Slices and maps contain internal pointers, so copies point to the same underlying data
}

/////////////////////////////
/// Functions masterclass ///
/////////////////////////////

// Basic syntax of functions:
// func - name - ( - arguments - ) - return type - {
// (indent) something we want to do
// }
func basics() {
	for ind := 0; ind < 5; ind++ {
		sayMessage("Hello, Go!", ind)
	}
	bfGreeting := "Heya"
	bfName := "Stacey"
	sayGreeting(&bfGreeting, &bfName)
	fmt.Println(bfName, "\n")
}

func sayMessage(msg string, idx int) {
	fmt.Println(msg)
	fmt.Println("The value of index is", idx, "\n")
}

// We can declare several args separated with comma and then set the type:
// We can also use pointers to manipulate the original variables:
func sayGreeting(greeting, name *string) {
	fmt.Printf("%v, %v!\n", *greeting, *name)
	*name = "Ted"
	fmt.Println(*name)
}

// Variatic parameters:
func variatics() {
	result := addingUp(1, 2, 3, 4, 5)
	fmt.Println("The sum is:", *result)
	result2 := addingUp2(1, 2, 3, 4, 5)
	fmt.Println("The sum is:", result2, "\n")

}

// Return pointer example
func addingUp(values ...int) *int {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	return &result // This value is pointing to a var located in the heap
}

// Named return values
func addingUp2(values2 ...int) (result2 int) {
	fmt.Println(values2)
	for _, val := range values2 {
		result2 += val
	}
	return
}

// Multiple returns
func multipleReturns() {
	mrd, err := mrdivide(5.0, 0.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(mrd)
}

func mrdivide(divideA, divideB float64) (float64, error) {
	if divideB == 0.0 {
		return 0.0, fmt.Errorf("Can't divide by zero!") // returning error instead of panic to allow execution
	}
	return divideA / divideB, nil // <-- returning nil as error because error wasn't present
}

func functionsMasterclass() {
	basics()
	variatics()
	multipleReturns()
	_, _ = mrdivide(5.0, 2.5)

	// Immediately invoked anonymous function
	func() {
		msg := "This is an immediately invoked function"
		fmt.Println(msg)
	}()
	// Declared variable func
	f := func() {
		fmt.Println("This is a declared variable function")
	}
	f()

	// Function signature for a divide function
	var fdivide func(float64, float64) (float64, error)
	fdivide = func(fa, fb float64) (float64, error) {
		if fb == 0.0 {
			return 0.0, fmt.Errorf("Can't divide by zero")
		} else {
			return fa / fb, nil
		}
	}
	fd, err := fdivide(5.0, 0.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(fd)

	// Methods declaration (cut & paste this snippet into the main func)
	g := greeter{
		greeting: "Heya",
		name:     "Golang",
	}
	g.greet()
	fmt.Println("The new name is:", g.name)

	// Functions summary
	//		- Basic syntax (if name is capitalised, function will be able to be exported
	//			> func foo() {
	//				...
	//			  }
	//
	//		- Parameters
	//			> Comma delimited list of variables and types
	//				· func foo(bar string, baz int)
	//			> Parameters of same type list type once
	//				· func foo(bar, baz int)
	//			> When pointers are passed in, the function can change the value in the caller
	//				· This is always true for data of slices and maps
	//			> Use variadic parameters to send list of same types in
	//				· Must use last parameter
	//				· Received as a slice
	//				· func foo (bar string, baz ...int)
	//
	//		- Return values
	//			> Single return values just list type
	//				· func foo() int
	//			> Multiple return value list types surrounded by parentheses
	//				· func foo() (int, error)
	//				· The (result type, error) paradigm is a very common idiom
	//			> Can use named return values
	//				· Initialises returned variable
	//				· Return using return keyword on its own
	//			> Can return addresses of local variables
	//				· Automatically promoted from local memory (stack) to shared memory (heap)
	//
	//		- Anonymous functions
	//			> Functions don't have names if they are:
	//				· Immediately invoked
	//					¬ func() {
	//						...
	//					  }()
	//				· Assigned to a variable or passed as an argument to a function
	//					¬ a := func() {
	//						...
	//					  }
	//					  a()
	//
	//		- Functions as types
	//			> Can assign functions to variables or use as arguments and return values in functions
	//			> Type signature is like function signature, with no parameter names. First parentheses are types coming
	//	  		  in, and second parenthesis are types coming out
	//				· var f func(string, string, int) (int, error)
	//
	//		- Methods
	//			> Function that executes in context of a type
	//			> Format:
	//				· func (g greeter) greet() {
	//					...
	//				  }
	//			> Receiver can be either value or pointer
	//				· Value receiver gets copy of type
	//				· Pointer receiver gets pointer to type
}

// Methods declaration
type greeter struct {
	greeting string
	name     string
}

func (g *greeter) greet() {
	fmt.Println(g.greeting, g.name)
	g.name = "pointer to greeter function!"
}

//////////////////
/// Interfaces ///
//////////////////

func interfaces() {
	// Basics

	// Composing interfaces

	// Type conversion
	// The empty interface
	// Type switches

	// Implementing with values vs pointers

	// Best practices
}

func main() {
	//basicVariables()
	//primitives()
	//constants()
	//arraysAndSlices()
	//mapsAndStructs()
	//controlFlow()
	//loops()
	//Defer, Panic and Recover:
	//deferOne()
	//deferTwo()
	//deferThree()
	//deferFour()
	//panicking()
	//recovering()
	//pointers()
	//functionsMasterclass()
	interfaces()
}
