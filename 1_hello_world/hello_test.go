package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("sying hello to people", func(t *testing.T) {
		got := Hello("David", "")
		want := "Hello, David"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in french", func(t *testing.T) {
		got := Hello("DK", "French")
		want := "Bonjour, DK"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	// Tells the test suite that this function is a test helper.
	// When it fails, the line number reported will be from the caller of this function.
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
