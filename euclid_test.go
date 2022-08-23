package main

import (
	"testing"
)

func Test1(t *testing.T) {
	got := gcd(527,1258)[0]
	want := 17

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

