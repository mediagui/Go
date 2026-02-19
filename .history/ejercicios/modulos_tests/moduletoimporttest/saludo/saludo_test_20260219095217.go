package saludo

import (
	"testing"
)

func TestGreeting(t *testing.T) {
	expected := "Hello, World!"
	actual := Greeting("World")
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}
