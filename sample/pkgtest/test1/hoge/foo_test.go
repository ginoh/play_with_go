package hoge

import (
	"testing"
)

func TestFooEcho(t *testing.T) {
	expect := "foo"
	actual := FooEcho()

	if expect != actual {
		t.Errorf("%s != %s", expect, actual)
	}
}
