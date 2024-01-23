package main

import "testing"

func TestCmdFormatter(t *testing.T) {
	text := "command firstArg secondArg"
	expected := []string{"command", "firstarg", "secondarg"}
	textFormatted := cmdFormatter(text)

	if textFormatted == nil {
		t.Error("cache is nil")
	}

	if textFormatted[0] != expected[0] {
		t.Error("command did not match")
	}

	if textFormatted[1] != expected[1] {
		t.Error("first arg did not match")
	}

	if textFormatted[2] != expected[2] {
		t.Error("second arg did not match")
	}
}
