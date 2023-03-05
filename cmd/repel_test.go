package main

import (
	"strconv"
	"testing"
)

type A struct {
}

func (receiver A) toString() string {
	return "A"
}

type B struct {
	A
}

func (receiver B) toString() string {
	return "B"
}

func TestName(t *testing.T) {
	val, _ := strconv.Atoi("nope")
	if val != 0 {
		t.Fatalf("Should have be set to default int")
	}
}
