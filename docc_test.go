package main

import (
	"testing"
)

func TestConvertToHTTPS(t *testing.T) {
	expected := "https://github.com/tcnksm/docc"
	got := ""

	got = convertURLToHTTPS("github.com:tcnksm/docc.git")
	if got != expected {
		t.Errorf("Expected '%v', but got '%v'.", expected, got)
	}

	got = convertURLToHTTPS("git@github.com:tcnksm/docc.git")
	if got != expected {
		t.Errorf("Expected '%v', but got '%v'.", expected, got)
	}

	got = convertURLToHTTPS("github.com:/~user/tcnksm/docc.git")
	if got != expected {
		t.Errorf("Expected '%v', but got '%v'.", expected, got)
	}

	got = convertURLToHTTPS("ssh://github.com/tcnksm/docc.git")
	if got != expected {
		t.Errorf("Expected '%v', but got '%v'.", expected, got)
	}

	got = convertURLToHTTPS("ssh://git@github.com/tcnksm/docc.git")
	if got != expected {
		t.Errorf("Expected '%v', but got '%v'.", expected, got)
	}

	got = convertURLToHTTPS("ssh://github.com:9418/tcnksm/docc.git")
	if got != expected {
		t.Errorf("Expected '%v', but got '%v'.", expected, got)
	}

	got = convertURLToHTTPS("ssh://github.com/~user/tcnksm/docc.git")
	if got != expected {
		t.Errorf("Expected '%v', but got '%v'.", expected, got)
	}

	got = convertURLToHTTPS("git://github.com/tcnksm/docc.git")
	if got != expected {
		t.Errorf("Expected '%v', but got '%v'.", expected, got)
	}

	got = convertURLToHTTPS("git://github.com:9418/tcnksm/docc.git")
	if got != expected {
		t.Errorf("Expected '%v', but got '%v'.", expected, got)
	}

	got = convertURLToHTTPS("git://github.com/~user/tcnksm/docc.git")
	if got != expected {
		t.Errorf("Expected '%v', but got '%v'.", expected, got)
	}

	got = convertURLToHTTPS("https://github.com/tcnksm/docc.git")
	if got != expected {
		t.Errorf("Expected '%v', but got '%v'.", expected, got)
	}

	got = convertURLToHTTPS("https://github.com:9418/tcnksm/docc.git")
	if got != expected {
		t.Errorf("Expected '%v', but got '%v'.", expected, got)
	}
}
