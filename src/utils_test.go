package main

import (
	"testing"
)

func TestisRealDir(t *testing.T) {

	test(t, ".", false)
	test(t, "coucou", true)

}


func test (t *testing.T, root string, isDir bool) {
	got := isRealDir(root)

	if got != isDir{
		t.Errorf("%s must be a fake dir ? %t and the real value is %t ", root, isDir, got)
	}

}