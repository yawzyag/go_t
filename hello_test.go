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
		got := Hello("Chris")
		want := "Hello, Chris"
		assertMessage(t, got, want)
	})

	t.Run("Say hello world if no variable passed to function", func(t *testing.T) {
		got := Hello("")
		want := "Hello, world"
		assertMessage(t, got, want)
	})
}
