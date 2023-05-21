package main

import (
	"regexp"
	"testing"
)

func TestHello(t *testing.T) {
	msg := Hello()
	res, err := regexp.MatchString("Welcome.", msg)
	if !res {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
