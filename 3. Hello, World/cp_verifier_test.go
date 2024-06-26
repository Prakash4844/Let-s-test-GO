package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Zaphkiel!!!", "")
		want := "Hello, Zaphkiel!!!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Japanese", func(t *testing.T) {
		got := Hello("ichigo", "Japanese")
		want := "Mosi mosi, ichigo"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Hindi", func(t *testing.T) {
		got := Hello("Imli", "Hindi")
		want := "Namaste, Imli"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Ouga buga", func(t *testing.T) {
		got := Hello("Bouga", "Ouga")
		want := "Ouga, Bouga"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if want != got {
		t.Errorf("got %q want %q", got, want)
	}
}
