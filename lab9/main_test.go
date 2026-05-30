package main

import (
	"testing"
)

func TestToYAML(t *testing.T) {
	tests := []struct {
		name     string
		input    any
		expected string
	}{
		{
			name: "Повна структура Server",
			input: Server{
				Host:       "localhost",
				Port:       8080,
				Debug:      true,
				AllowedIPs: []string{"192.168.1.1", "10.0.0.1"},
			},
			expected: "host: \"localhost\"\nport: 8080\ndebug: true\nallowed_ips:\n  - \"192.168.1.1\"\n  - \"10.0.0.1\"",
		},
		{
			name: "Структура з порожнім масивом",
			input: Server{
				Host:       "127.0.0.1",
				Port:       3000,
				Debug:      false,
				AllowedIPs: []string{},
			},
			expected: "host: \"127.0.0.1\"\nport: 3000\ndebug: false\nallowed_ips: []",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ToYAML(tt.input)

			if err != nil {
				t.Fatalf("отримано помилку: %v", err)
			}

			if result != tt.expected {
				t.Errorf("\nОЧІКУВАЛОСЬ:\n%s\n\nОТРИМАНО:\n%s", tt.expected, result)
			}
		})
	}
}


var benchData = Server{
	Host:       "localhost",
	Port:       8080,
	Debug:      true,
	AllowedIPs: []string{"192.168.1.1", "10.0.0.1"},
}

func BenchmarkToYAML(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = ToYAML(benchData)
	}
}

func BenchmarkToJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = ToJSON(benchData)
	}
}
