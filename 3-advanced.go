// Variadic Functions
// Anonymous Functions and Closures
// Recursion
// Pointers
// Structs
// Methods
// Interfaces
// Errors
package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	// Variadic Functions
	/*
	   Call functions with any number of trailing
	   arguments.
	*/
	sum(1, 2)    // Prints [1 2] 3
	sum(1, 2, 3) // Prints [1 2 3] 6

	nums := []int{1, 2, 3, 4}
	sum(nums...) // Prints [1 2 3 4] 10

	// Anonymous Functions and Closures
	/*
	   Anonymous functions are useful when you want to define a function inline without having to name it.

	   The intSeq function returns another function, which we define anonymously in the body of intSeq. The returned function closes over the variable i to form a closure.
	*/
	nextInt := intSeq()
	/*
	   We call intSeq, assigning the result, a function, to nextInt. This function value captures its own i value, which will be updated each time we call nextInt.
	*/

	fmt.Println(nextInt()) // Prints 1
	fmt.Println(nextInt()) // Prints 2
	fmt.Println(nextInt()) // Prints 3

	newInts := intSeq()
	fmt.Println(newInts()) // Prints 1

	// Recursion
	/*
	   Function calls itself until it reaches fact(0)
	*/
	fmt.Println(fact(7)) // Prints 5040

	// Pointers
	/*
	   Pointers allow you to pass references to values
	*/
	i := 1
	fmt.Println("initial:", i) // Prints initial: 1

	zeroval(i)
	fmt.Println("zeroval:", i) // Prints zeroval: 1

	zeroptr(&i)
	fmt.Println("zeroptr:", i) // Prints zeroptr: 0
	/*
	   zeroptr has an *int parameter meaning it takes an int pointer. The *iptr code in the function body dereferences the pointer from its memory address to the current value at that address.

	   Assigning a value to a dereferenced pointer changes the value at the referenced address.
	*/

	fmt.Println("pointer:", &i) // Prints pointer: 0x42131100

	// Structs
	/*
	   Structs are typed collections of fields. Useful for grouping data together to form records.
	*/
	fmt.Println(person{"Bob", 20})              // Create a new struct
	fmt.Println(person{name: "Alice", age: 30}) // You can name fields when initializing a struct
	fmt.Println(person{name: "Fred"})           // Omitted fields will be zero-valued
	fmt.Println(&person{name: "Ann", age: 40})  // An & prefix yields a pointer to the struct
	fmt.Println(newPerson("Jon"))               // Its idiomatic to encapsulate new struct creation in constructor functions

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name) // Access struct fields with a dot

	sp := &s
	fmt.Println(sp.age) // You can also use dots with struct pointers. The pointers are automatically dereferenced.

	sp.age = 51
	fmt.Println(sp.age) // Structs are mutable.

	/*
	   Prints:
	   {Bob 20}
	   {Alice 30}
	   {Fred 0}
	   &{Ann 40}
	   &{Jon 42}
	   Sean
	   50
	   51
	*/

	// Methods

	r := rect{width: 10, height: 5}

	fmt.Println("area: ", r.area())   // Prints area: 50
	fmt.Println("perim: ", r.perim()) // Prints perim: 30

	rp := &r
	fmt.Println("area: ", rp.area())   // Prints area: 50
	fmt.Println("perim: ", rp.perim()) // Prints perim: 30

	/*
	   You may want to use a pointer receiver type to avoid copying on method calls or to allow the method to mutate the receiving struct.
	*/

	// Interfaces
	/*
	   Interfaces are named collections of method signatures

	   To implement an interface, we need to implement all of the methods in the interface. So the geometry interface has area and perim methods, so we need to create area and perim methods for each of our structs.

	   If a variable has an interface type, then we can call methods that are in the named interface e.g. g.area().
	*/
	re := rectangle{width: 3, height: 4}
	ci := circle{radius: 5}

	measure(re)
	measure(ci)

	/*
	   Prints:
	   {3 4}
	   12
	   14
	   {5}
	   78.5398
	   31.4159
	*/

	// Errors
	for _, i := range []int{7, 42} {
		if res, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", res)
		}
	}
	for _, i := range []int{7, 42} {
		if res, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", res)
		}
	}

	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}

	/*
	   Prints:
	   f1 worked: 10
	   f1 failed: Cant work with 42
	   f2 worked: 10
	   f2 failed: 42 - Cant work with it
	   42
	   Cant work with it
	*/
}

// Variadic Functions
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

// Anonymous Functions and Closures
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// Recursion
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

// Pointers
func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

// Structs
type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

// Methods
type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

// Interfaces
type geometry interface {
	area() float64
	perim() float64
}

type rectangle struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (re rectangle) area() float64 {
	return re.width * re.height
}
func (re rectangle) perim() float64 {
	return 2*re.width + 2*re.height
}

func (ci circle) area() float64 {
	return math.Pi * ci.radius * ci.radius
}
func (ci circle) perim() float64 {
	return 2 * math.Pi * ci.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

// Errors
func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("Cant work with 42")
	}
	return arg + 3, nil
}

type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "Cant work with it"}
	}
	return arg + 3, nil
}
