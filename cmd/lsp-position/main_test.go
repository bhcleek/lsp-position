package main

import (
	"strings"
	"testing"
	"unicode/utf16"
)

func TestCodeUnit(t *testing.T) {
	tests := []struct {
		name     string
		in       string
		expected uint64
	}{
		{
			name:     "simple",
			in:       "ascii only",
			expected: uint64(len("ascii only")),
		}, {
			name:     "two place of interest",
			in:       "⌘⌘",
			expected: 2,
		}, {
			name:     "Deseret Capital Letter Long I",
			in:       "𐐀",
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, r := range tt.in {
				r1, r2 := utf16.EncodeRune(r)
				t.Logf("%c surrogate: %v", r, utf16.IsSurrogate(r))
				t.Logf("%c = (%x, %x)", r, r1, r2)
			}

			s := tt.in + " "
			actual, err := codeUnitOf(strings.NewReader(s), uint64(len(s)-1))
			if err != nil {
				t.Fatal(err)
			}

			if actual != tt.expected {
				t.Errorf("codeUnitOf(...) = %v; want %v", actual, tt.expected)
			}
		})
	}
}
