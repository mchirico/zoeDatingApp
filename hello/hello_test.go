package hello

import (
	"testing"
	"fmt"
)

func TestBozo(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Bozo(); got != tt.want {
				t.Errorf("Bozo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDating(t *testing.T) {
	fmt.Printf("this is a test...")
}

