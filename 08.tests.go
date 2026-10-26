package main

import (
	"testing"
)

func TestUp(t *testing.T) {
	r := Up([]string{"salut", "(up)"})
	if r[0] != "SALUT" {
		t.Error("Up marche pas", r)
	}
}

func TestLow(t *testing.T) {
	r := Low([]string{"SALUT", "(low)"})
	if r[0] != "salut" {
		t.Error("Low marche pas", r)
	}
}

