// Packages are ways of grouping up related Go code together.
package main

import "fmt" //With import fmt we are importing a package which contains the Println function that we use to print.
//The func keyword is how you define a function with a name and a body.

func Hello(name string) string { //means this function returns a string.
	return "Hello, " + name
}

func main() {
	fmt.Println(Hello("world"))
}
