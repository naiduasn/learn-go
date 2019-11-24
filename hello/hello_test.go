package main

import (
	"testing"
)

func TestHello(t *testing.T) {

	assertMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("Expected %q Got %q", got, want)
		}
	}

	t.Run("Saying Hello to people", func(t *testing.T) {
		getFunc := Hello("Naidu")
		toTest := "Hello, Naidu"

		assertMessage(t, getFunc, toTest)
	})

	t.Run("Saying Hello World to When an empty String is supplied", func(t *testing.T) {
		getFunc := Hello("")
		toTest := "Hello, World"

		assertMessage(t, getFunc, toTest)
	})

}
