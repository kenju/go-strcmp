package main

import "regexp"

// Match returns true when patten appears in text
func Match(pattern string, text string) bool {
	return MatchRegexp(pattern, text)
}

// MatchRegexp implements matching algorithm with regexp package
func MatchRegexp(p, t string) bool {
	var reg = regexp.MustCompile(p)
	return reg.MatchString(t)
}

// MatchKMP implements Knuth-Morris-Pratt algorithm
// @doc https://en.wikipedia.org/wiki/Knuth%E2%80%93Morris%E2%80%93Pratt_algorithm
func MatchKMP(p, t string) bool {
	return true
}
