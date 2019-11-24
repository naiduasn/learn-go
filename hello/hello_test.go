package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	getFunc := Hello("Naidu")
	toTest := "Hello Naidu"

	if getFunc != toTest {
		t.Errorf("Expected %q Got %q", toTest, getFunc)
	}
}

