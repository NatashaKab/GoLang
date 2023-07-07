/*
Writing a test is just like writing a function, with a few rules

	It needs to be in a file with a name like xxx_test.go
	The test function must start with the word Test
	The test function takes one argument only t *testing.T
	In order to use the *testing.T type, you need to import "testing", like we did with fmt in the other file
*/
package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want) //We are calling the Errorf method on our t which will print out a message and fail the test. The f stands for format which allows us to build a string with values inserted into the placeholder values %q.
	})

	t.Run("empty string defaults to 'world'", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Elodie", "French")
		want := "Bonjour, Elodie"
		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
