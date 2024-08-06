package main

import "testing"

func TestSoma(t *testing.T) {
	result := Soma(200, 20)
	if result != 220 {
		t.Errorf("Soma(20, 20) = %d; want 40", result)
	}
}
