package main

/*
	This is just here to clean up after the other tests.
	It's named z_test.go so it runs last.
*/

import (
	"os"
	"testing"
)

func TestCleanup(t *testing.T) {
	os.RemoveAll("testdata/test-filesystem/gopath/pkg")
}
