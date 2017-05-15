package main

import "testing"

func Test_getMode(t *testing.T) {
	var args []string
	if getMode(args) != "readAll" {
		t.Fail()
	}

	args = []string{"lala", "blabla"}
	if getMode(args) != "read" {
		t.Fail()
	}

	args = []string{"lala=abc", "blabla=123"}
	if getMode(args) != "write" {
		t.Fail()
	}
}

func Test_splitWriteArg(t *testing.T) {
	exp := [2]string{"haha", "123"}
	if splitWriteArg("haha=123") != exp {
		t.Fail()
	}
}
