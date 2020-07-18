# go-strcmp

A simple go library for string matching algorithms.

## Usage

	import strcmp "github.com/kenju/go-strcmp"

	pattern := "abcabac"
	text := "ababcababcabac"

    strcmp.Match(pattern, text) // returns true
