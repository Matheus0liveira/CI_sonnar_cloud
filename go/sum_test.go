package main

import "testing"

func TestSum(t *testing.T) {
	total := Sum(10, 15)

	if total != 25 {
		t.Errorf("Result error")
	}

}