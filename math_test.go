package main

import "testing"

func TestSoma(t *testing.T) {
	result := Soma(200, 40)
	if result != 240 {
		t.Errorf("Soma(200, 40) = %d; want 240", result)
	}
}
