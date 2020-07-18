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
