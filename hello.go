// Packages are ways of grouping up related Go code together.
package main

import "fmt" //With import fmt we are importing a package which contains the Println function that we use to print.
//The func keyword is how you define a function with a name and a body.

// Constants are defined like so
const (
	french  = "French"
	spanish = "Spanish"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

func Hello(name string, language string) string { //means this function returns a string.
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello(spanishHelloPrefix, spanish))
}
