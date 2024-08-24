package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("DK", "")
		want := "Hello, DK"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello World' when an empty strin gis supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Korean", func(t *testing.T) {
		got := Hello("DK", "Korean")
		want := "Ahn nyeong, DK"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("DK", "Spanish")
		want := "Hola, DK"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Got %q, Want %q", got, want)
	}
}
