package main

// Match returns true when patten appears in text
func Match(pattern string, text string) bool {
	return MatchKMP(pattern, text)
}
