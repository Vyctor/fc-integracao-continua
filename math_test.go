package main

import "testing"

func TestSoma(t *testing.T) {
	result := Soma(20, 20)
	if result != 40 {
		t.Errorf("Soma(20, 20) = %d; want 40", result)
	}
}
