package dating

import (
	"fmt"
	"testing"
)

// Review maps and slices

func TestF(t *testing.T) {
	Mytest()
}

func TestWarmup(t *testing.T) {
	r := Warmup()
	fmt.Printf("r=%v\n", r)

	if val, ok := r["bozo"]; ok {
		fmt.Printf("yes it works...: %v\n", val)
	}

	if r["mousey"] != 3 {
		t.Fatalf("Bad value")
	}
}

func TestPref(t *testing.T) {
	r := CreateDatingThingy()
	if r["Mouse"].Animal == "cat" {
		fmt.Printf("yay! it works!")
	}
}
