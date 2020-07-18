package main_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	strcmp "github.com/kenju/go-strcmp"
)

/**
 * Unit Tests
 */

func TestMatch(t *testing.T) {
	testMatchAlgorithm(t, strcmp.Match)
}

func TestMatchRegexp(t *testing.T) {
	testMatchAlgorithm(t, strcmp.MatchRegexp)
}

func TestMatchKMP(t *testing.T) {
	testMatchAlgorithm(t, strcmp.MatchKMP)
}

/**
 * Benchmark
 */

func benchMatchAlgorithm(b *testing.B, fn matchFunc) {
	pattern := "abcabac"
	text := "ababcababcabac"
	for n := 0; n < b.N; n++ {
		fn(pattern, text)
	}
}

func BenchmarkMatchRegexp(b *testing.B) {
	benchMatchAlgorithm(b, strcmp.MatchRegexp)
}

func BenchmarkMatchKMP(b *testing.B) {
	benchMatchAlgorithm(b, strcmp.MatchKMP)
}

/**
 * Test Utility
 */

type testCase struct {
	pattern  string
	text     string
	expected bool
}

func testCases() map[string]testCase {
	return map[string]testCase{
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
}

type matchFunc func(pattern string, text string) bool

func testMatchAlgorithm(t *testing.T, fn matchFunc) {
	for name, tt := range testCases() {
		t.Run(name, func(t *testing.T) {
			actual := fn(tt.pattern, tt.text)
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
