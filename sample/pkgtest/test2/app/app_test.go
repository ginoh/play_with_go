package app

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	println("before m.Run")

	code := m.Run()

	println("after m.Run")

	os.Exit(code)
}

func TestAppEcho(t *testing.T) {
	expect := "App"
	actual := AppEcho()

	if expect != actual {
		t.Errorf("%s != %s", expect, actual)
	}
}
