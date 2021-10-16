// Slices
// Maps
// Ranges
// Functions
// Multiple Return Values
package main

import "fmt"

func beyondBasics() {

	// Slices
	a := make([]string, 3)
	fmt.Println("emp:", a) // Prints emp: [  ]
	/*
	   Unlike arrays, slices are typed only by the elements they contain, not the number of elements. To create an empty slice with non-zero length, use make. Here a slice of strings with length 3, initially zero-valued.
	*/

	a[0] = "a" // Can set and get just like arrays
	a[1] = "b"
	a[2] = "c"
	fmt.Println("set:", a)    // prints set: [a b c]
	fmt.Println("get:", a[2]) // prints get: c

	fmt.Println("len:", len(a)) // prints len: 3. Length of the slice.

	a = append(a, "d")
	a = append(a, "e", "f")
	fmt.Println("apd:", a) // Prints apd: [a b c d e f]
	/*
	   Slices can use append, which returns a slice containing one or more new values. We need to accept a return value from append since we may get a new slice value.
	*/

	b := make([]string, len(a))
	copy(b, a)
	fmt.Println("cpy:", b) // Prints cpy: [a b c d e f]. Create an empty slice b of the same length as a and copy from a into b.

	c := a[2:5]
	fmt.Println("sl1:", c) // Prints sl1: [c d e]
	/*
	   Slices support a slice operator with the syntax slice[low:high]. This gets a slice of the elements a[2], a[3] and a[4].
	*/

	c = a[:5]
	fmt.Println("sl2:", c) // Prints [a b c d e]. Up to, but excluding 5.

	c = a[2:]
	fmt.Println("sl3:", c) // Prints sl3: [c d e f]. Up from and including 2.

	d := []string{"g", "h", "i"}
	fmt.Println("dcl:", d) // Prints dcl: [g h i]. Declare and initialise a variable for slice in a single line.

	twoD := make([][]int, 3)
	for e := 0; e < 3; e++ {
		innerLen := e + 1
		twoD[e] = make([]int, innerLen)
		for f := 0; f < innerLen; f++ {
			twoD[e][f] = e + f
		}
	}
	fmt.Println("2d: ", twoD) // Prints 2d:  [[0] [1 2] [2 3 4]].
	/*
	   Slices can create multi-dimensional data structures. The length of the inner slices can vary while arrays cannot do this.
	*/

	// Maps
	// same as dict

	g := make(map[string]int) // Create empty map: make(map[key-type]value-type).

	g["k1"] = 7
	g["k2"] = 13

	fmt.Println("map:", g) // Prints map: map[k1:7 k2:13]. Set key/value pairs using name[key] = value.

	h1 := g["k1"]
	fmt.Println("h1: ", h1) // prints v1:  7

	fmt.Println("len:", len(g)) // prints len: 2

	delete(g, "k2")
	fmt.Println("map:", g) // Prints map: map[k1:7]. Remove key/value from map.

	_, prs := g["k2"]
	fmt.Println("prs:", prs) // Prints prs: false. The optional second return value when getting a map value indicates if the key was present in the map. This differentiates between missing keys and keys with zero values. We dont need the value itself so we ignore it with the blank identifier _.

	i := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", i) // Prints map: map[bar:2 foo:1]. Declare and initialise a new map in one line.

	// Ranges

	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums { // ignore index with _
		sum += num
	}
	fmt.Println("sum:", sum) // Prints sum: 9. Sums the numbers in a slice.

	for j, num := range nums {
		if num == 3 {
			fmt.Println("index:", j)
		}
	} // Prints index: 1. range on arrays and slices gives both the index and value for each item.

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	} // Prints a -> apple \n b -> banana. range on map iterates over key/value pairs.

	for k := range kvs {
		fmt.Println("key:", k)
	} // Prints key: a \n key: b. range can iterate over keys of map.

	for i, c := range "go" {
		fmt.Println(i, c)
	} // Prints 0 103 \n 1 111. range on strings iterates over Unicode. First value is the byte index of the rune. Second is the rune itself.

	// Functions
	functionTest()

	// Multiple Return Values
	multiReturnValues()

}

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func functionTest() {
	res := plus(1, 2)
	fmt.Println("1+2 =", res) // prints 1+2 = 3

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res) // prints 1+2+3 = 6
}

func vals() (int, int) {
	return 3, 7
}

func multiReturnValues() {
	a, b := vals()
	fmt.Println(a) // prints 3
	fmt.Println(b) // prints 7

	_, c := vals()
	fmt.Println(c)
}
