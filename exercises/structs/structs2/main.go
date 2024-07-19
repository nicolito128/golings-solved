// structs2
// Make me compile!
package main

import "fmt"

type Person struct {
	// don't just create the phone field here. embed a new struct
	Phone
	name string
	age  int
}

type Phone struct {
	phone string
}

func main() {
	// contactDetails := ContactDetails{}
	person := Person{name: "John", age: 32, Phone: Phone{"+54 9 222 555"}}
	fmt.Printf("%s is %d years old and his phone is %s\n", person.name, person.age, person.phone)
}
