// Sorting
// Custom Sorting
package main

import (
	"fmt"
	"sort"
)

func commonFunctions() {

	// Sorting
	/*
	   Sorting of built in types

	   Sorting is in place, so it changes the given slice and doesnt return a new one
	*/
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs) // Prints Strings: [a b c]

	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:", ints) // Prints Ints: [2 4 7]

	a := sort.IntsAreSorted(ints)
	fmt.Println("Sorted:", a) // Prints Sorted: true. Check if a slice is already in sorted order.

	/*
	   Custom sorting
	*/

	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(byLength(fruits)) // Convert fruits slice to byLength, then use sort.Sort on that typed slice
	fmt.Println(fruits)         // Prints [kiwi peach banana]

}

type byLength []string

/*
Implement sort.Interface - Len, Less, and Swap - on the byLength type so we can use the sort packages generic Sort function.
*/
func (b byLength) Len() int {
	return len(b)
}

func (b byLength) Swap(c, d int) {
	b[c], b[d] = b[d], b[c]
}

func (b byLength) Less(c, d int) bool {
	return len(b[c]) < len(b[d]) // sort in order of increasing length
}
