package main

import (
	"testing"
)

func TestAppEcho(t *testing.T) {
	expect := "App"
	actual := AppEcho()

	if expect != actual {
		t.Errorf("%s != %s", expect, actual)
	}
}
