// slices1
// Make me compile!

package main

import "fmt"

func main() {
	a := make([]complex128, 2, 5*3) // play with length and capacity
	fmt.Println("length of 'a':", len(a))
	fmt.Println("capacity of 'a':", cap(a))
	fmt.Println(a)
}
