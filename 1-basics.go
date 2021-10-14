// Summary of docs from https://gobyexample.com/
//
// Imports
// Constants
// Values
// Variables
// Loops
// If/Else
// Switch
// Arrays
package main

// import "fmt" // Single import
/*
Multi line import
import (
    "fmt"
    "math"
)
*/
import (
	"fmt"
	"math"
	"time"
)

// Constants
const s string = "constant"

func main() {
	fmt.Println("hello world") // Print string

	// Values
	fmt.Println("go" + "lang") // Add strings together

	fmt.Println("1+1 =", 1+1)         // Add integers
	fmt.Println("7.0/3.0 =", 7.0/3.0) // Divide floats

	fmt.Println(true && false) // false
	fmt.Println(true || false) // true
	fmt.Println(!true)         // false

	// Variables
	var a = "initial"
	fmt.Println(a) // prints initial

	var b, c int = 1, 2 // Declare multiple variables at once
	fmt.Println(b, c)   // prints 1 2

	var d = true
	fmt.Println(d) // prints true

	var e int
	fmt.Println(e) // Prints 0. Variables without initialisation are zero-valued. The zero value for an int is 0.

	f := "apple"
	fmt.Println(f) // Prints apple. := is shorthand for declaring and initialising a variable. This is the same as saying `var f string = "apple"`

	// := can only be used inside a function. var can be used inside and outside of functions. Its just shorthand to make it easier to declare variables.

	// Constants
	fmt.Println(s) // prints constant

	const x = 500000000 // Constants can appear anywhere a var statement can

	const g = 3e20 / x // Constant expressions perform arithmetic with arbitrary precision
	fmt.Println(g)     // prints 6e+11

	fmt.Println(int64(g)) // Prints 600000000000. A numeric constant has no type until its given one, like by an explicit conversion

	fmt.Println(math.Sin(x)) // Prints -0.28470407323754404. A number can be given a type by using it in a context that requires one, such as a variable assignment or function call. Here, math.Sin expects a float64.

	// Loops
	// The only type of loop is `for`. There is no `while` etc. just variations on `for`.
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	} // Basic loop, with a single condition
	/*
	   Prints:
	   1
	   2
	   3
	*/

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	} // initial/condition/after for loop
	/*
	   Prints:
	   7
	   8
	   9
	*/

	for {
		fmt.Println("loop")
		break
	} // while style for loop. Will loop repeatedly until you break out of the loop or return from the enclosing function.
	/*
	   Prints:
	   loop
	*/

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	} // Continue to the next iteration of the loop
	/*
	   Prints:
	   1
	   3
	   5
	*/

	// If/Else
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	} // Prints 7 is odd. Basic if/else

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	} // Prints 8 is divisible by 4. If without else.

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	} // Prints 9 has 1 digit. A statement can precede conditionals. Any variables declared in this statement are available in all branches.

	// Switch
	k := 2
	fmt.Print("Write ", k, " as ")
	switch k {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	} // Prints Write 2 as two. Basic switch.

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Its the weekend")
	default:
		fmt.Println("Its a weekday")
	} // Prints Its a weekday. Commas separate multiple expressions in the same case statement.

	l := time.Now()
	switch {
	case l.Hour() < 12:
		fmt.Println("Its before noon")
	default:
		fmt.Println("Its after noon")
	} // Prints Its after noon. Switch without an expressions is like if/else. Also, case expression can be non-constants.

	whatAmI := func(m interface{}) {
		switch l := m.(type) {
		case bool:
			fmt.Println("Im a bool")
		case int:
			fmt.Println("Im an int")
		default:
			fmt.Printf("Dont know type %T\n", l)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
	/*
	   Prints:
	   Im a bool
	   Im an int
	   Dont know type string
	   A type switch compares types instead of values.
	*/

	// Arrays
	var n [5]int
	fmt.Println("emp:", n) // Prints emp: [0 0 0 0 0]
	/*
	   An array that holds exactly 5 ints. The type of elements and length are both part of the arrays type. By default an array is zero-valued.
	*/
	n[4] = 100                // Set value at index. Arrays start at 0.
	fmt.Println("set:", n)    // prints set: [0 0 0 0 100]
	fmt.Println("get:", n[4]) // prints get: 100

	fmt.Println("len:", len(n)) // prints len: 5

	o := [5]int{1, 2, 3, 4, 5} // declare and initialise array in one line
	fmt.Println("dcl:", o)     // Prints dcl: [1 2 3 4 5]

	var twoD [2][3]int
	for p := 0; p < 2; p++ {
		for q := 0; q < 3; q++ {
			twoD[p][q] = p + q
		}
	}
	fmt.Println("2d: ", twoD) // Prints 2d:  [[0 1 2] [1 2 3]]
	/*
	   Arrays are one dimensional but you can compose types to build multidimensional data structures.
	*/

}
