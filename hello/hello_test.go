package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	}

	t.Run("say hello to people", func(t *testing.T) {
		got := Hello("NMSL", "")
		want := "Hello, NMSL"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello world when an empty string is suppplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Alex", "French")
		want := "Bonjour, Alex"
		assertCorrectMessage(t, got, want)
	})

}
