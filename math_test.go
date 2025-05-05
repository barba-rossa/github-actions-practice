package main

import "testing"

func TestSoma(t *testing.T) {

	total := sum(15, 14)

	if total != 30 {
		t.Errorf("Invalid sum result: got %d, expected %d", total, 30)
	}
}
