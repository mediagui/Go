package saludo

import (
	"testing"
)

// TestGreeting prueba la función Greeting con casos básicos
func TestGreeting(t *testing.T) {
	expected := "Hello World"
	actual := Greeting("World")
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

// TestGreetingTableDriven prueba la función con múltiples casos
func TestGreetingTableDriven(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Nombre normal",
			input:    "World",
			expected: "Hello World",
		},
		{
			name:     "Nombre con un apellido",
			input:    "Juan García",
			expected: "Hello Juan García",
		},
		{
			name:     "Nombre vacío",
			input:    "",
			expected: "Hello ",
		},
		{
			name:     "Nombre mayúsculas",
			input:    "ALICE",
			expected: "Hello ALICE",
		},
		{
			name:     "Nombre con espacios",
			input:    "John Doe Smith",
			expected: "Hello John Doe Smith",
		},
		{
			name:     "Nombre con números",
			input:    "User123",
			expected: "Hello User123",
		},
		{
			name:     "Nombre con caracteres especiales",
			input:    "José María",
			expected: "Hello José María",
		},
		{
			name:     "Nombre corto",
			input:    "An",
			expected: "Hello An",
		},
		{
			name:     "Nombre con espacios al inicio",
			input:    " Bob",
			expected: "Hello  Bob",
		},
		{
			name:     "Nombre con espacios al final",
			input:    "Charlie ",
			expected: "Hello Charlie ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := Greeting(tt.input)
			if actual != tt.expected {
				t.Errorf("Greeting(%q) = %q, expected %q", tt.input, actual, tt.expected)
			}
		})
	}
}

// TestGreetingNotEmpty verifica que el resultado no esté vacío
func TestGreetingNotEmpty(t *testing.T) {
	result := Greeting("Test")
	if result == "" {
		t.Error("Greeting should not return empty string")
	}
}

// TestGreetingContainsInput verifica que el resultado contiene el input
func TestGreetingContainsInput(t *testing.T) {
	input := "Developer"
	result := Greeting(input)
	if !contains(result, input) {
		t.Errorf("Greeting(%q) should contain the input name, got %q", input, result)
	}
}

// TestGreetingFormat verifica que el resultado tiene el formato correcto
func TestGreetingFormat(t *testing.T) {
	tests := []string{"Alice", "Bob", "Charlie", "Diana", "Eve"}
	for _, name := range tests {
		result := Greeting(name)
		expected := "Hello " + name
		if result != expected {
			t.Errorf("Greeting(%q) = %q, expected %q", name, result, expected)
		}
	}
}

// Función auxiliar para verificar si una cadena contiene un substring
func contains(str, substr string) bool {
	for i := 0; i <= len(str)-len(substr); i++ {
		if str[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
