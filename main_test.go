package main_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	strcmp "github.com/kenju/go-strcmp"
)

func TestMatch(t *testing.T) {
	tests := map[string]struct {
		pattern  string
		text     string
		expected bool
	}{
		"empty": {
			pattern:  "",
			text:     "",
			expected: true,
		},
		"single char": {
			pattern:  "-",
			text:     "-",
			expected: true,
		},
		"same input": {
			pattern:  "abcabac",
			text:     "abcabac",
			expected: true,
		},
		"simple true": {
			pattern:  "abcabac",
			text:     "ababcababcabac",
			expected: true,
		},
		"simple false": {
			pattern:  "abab",
			text:     "abcabc",
			expected: false,
		},
		"number true": {
			pattern:  "0123012",
			text:     "010380123012",
			expected: true,
		},
		"number false": {
			pattern:  "0123012",
			text:     "010102030101",
			expected: false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			actual := strcmp.Match(tt.pattern, tt.text)
			comp(t, tt.expected, actual)
		})
	}
}

func comp(t *testing.T, expected, actual interface{}) {
	diff := cmp.Diff(expected, actual)
	if diff != "" {
		t.Fatalf(diff)
	}
}

func BenchmarkMatchRegexp(b *testing.B) {
	pattern := "abcabac"
	text := "ababcababcabac"
	for n := 0; n < b.N; n++ {
		strcmp.MatchRegexp(pattern, text)
	}
}
