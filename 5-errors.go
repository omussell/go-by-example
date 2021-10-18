// Errors
// Panic
// Recover

package main

import (
	"errors"
	"fmt"
)

func f1(arg int) (int, error) { // Errors are the last return value and have type `error`, a built in interface
	if arg == 42 {
		return -1, errors.New("cant work with 42") // errors.New constructs a basic error value with the given error message
	}
	return arg + 3, nil // A nil value in the error position indicates that there was no error
}

type argError struct {
	arg  int
	prob string
	/*
		Its possible to use custom types as errors by implementing the Error() method on them
	*/
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "cant work with it"}
	}
	return arg + 3, nil
}

func mayPanic() {
	panic("a problem")
}

func handleErrors() {
	/*
		The two loops below test out each of our error-returning functions. Note that the use of an inline error check on the if line is a common idiom in Go code.
	*/

	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	/*
		If you want to programmatically use the data in a custom error, you’ll need to get the error as an instance of the custom error type via type assertion.
	*/
	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}

	/*
		A panic is used when something goes unexpectedly wrong.

		A common use of panic is to abort if a function returns an error value that we dont know how to or want to handle.
	*/
	// panic("a problem")

	/*
		A recover can stop a panic from aborting the program and let it continue with execution instead

		recover must be called within a deferred function. When the enclosing function panics, the defer will activate and a recover call within it will catch the panic.

		The return value of recover is the error raised in the call to panic.
	*/
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	mayPanic()

	/*
		This code will not run, because mayPanic panics.

		The execution of handleErrors stops at the point of the panic and resumes in the deferred closure.
	*/
	fmt.Println("After mayPanic()")
}
