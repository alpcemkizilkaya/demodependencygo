package foo_test

import (
	"testing"

	"github.com/alpcemkizilkaya/demodependencygo/foo"
)

func TestFoo(t *testing.T) {
	test_string := foo.Foo()
	if test_string != "Hello dep" {
		t.Error("String is not equal to class value")
	}

}
