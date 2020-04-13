package main

import "testing"

func TestHello(t *testing.T) {

	assertMessage := func(t *testing.T, got string, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("Say hello + variable passed to function", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertMessage(t, got, want)
	})

	t.Run("Say hello world if no variable passed to function", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"
		assertMessage(t, got, want)
	})
	t.Run("Say Hola, segundo parametro para lenguaje", func(t *testing.T) {
		got := Hello("Yesid", "Spanish")
		want := "Hola, Yesid"
		assertMessage(t, got, want)
	})
	t.Run("Say Bonjour, second parameter languaje", func(t *testing.T) {
		got := Hello("Yesid", "French")
		want := "Bonjour, Yesid"
		assertMessage(t, got, want)
	})
	t.Run("Say Привет, second argument languaje", func(t *testing.T) {
		got := Hello("хуанхо", "Ruso")
		want := "Привет, хуанхо"
		assertMessage(t, got, want)
	})
}
