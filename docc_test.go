package main

import (
	"testing"
)

func TestConvertToHTTPS(t *testing.T) {
	
	args := []string{
		"github.com:tcnksm/docc.git",
		"git@github.com:tcnksm/docc.git",
		"github.com:/~user/tcnksm/docc.git",
		"ssh://github.com/tcnksm/docc.git",
		"ssh://git@github.com/tcnksm/docc.git",
		"ssh://github.com:9418/tcnksm/docc.git",
		"ssh://github.com/~user/tcnksm/docc.git",
		"git://github.com/tcnksm/docc.git",
		"git://github.com:9418/tcnksm/docc.git",
		"git://github.com/~user/tcnksm/docc.git",
		"https://github.com/tcnksm/docc.git",
		"https://tcnksm@github.com/tcnksm/docc.git",
		"https://tcnksm:PASS@github.com/tcnksm/docc.git",
		"https://github.com:9418/tcnksm/docc.git",
	}

	expected := "https://github.com/tcnksm/docc"

	for _, arg := range args {
		got := convertURLToHTTPS(arg)
		if got != expected {
			t.Errorf("Expected '%v', but got '%v'.", expected, got)
		}
	}
}
