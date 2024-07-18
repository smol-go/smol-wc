package main

import (
	"testing"
)

var info = get_info("test.txt")

func TestNumLines(t *testing.T) {
	want := 7137
	if info.numLines != want {
		t.Fatalf("want %d, got %d", want, info.numLines)
	}
}

func TestNumWords(t *testing.T) {
	want := 58159
	if info.numWords != want {
		t.Fatalf("want %d, got %d", want, info.numWords)
	}
}

func TestNumChars(t *testing.T) {
	want := 339120
	if info.numChars != want {
		t.Fatalf("want %d, got %d", want, info.numChars)
	}
}

func TestNumBytes(t *testing.T) {
	want := 341836
	if info.numBytes != want {
		t.Fatalf("want %d, got %d", want, info.numBytes)
	}
}
