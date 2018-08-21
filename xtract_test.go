package xtract

import (
	"reflect"
	"testing"
)

func TestExtract(t *testing.T) {
	v := "<p>This is a paragraph <a>This is a link inside a paragraph</a></p>"
	assertEq("This is a paragraph This is a link inside a paragraph", Value(v), t)

	assertEq("This is a paragraph", ValueLim(v, 4), t)
}

func assertEq(x, y interface{}, t *testing.T) {
	if !reflect.DeepEqual(x, y) {
		t.Errorf("expected: %v, but got: %v", x, y)
	}
}
