package strcmp

// MatchNaively implements string match in naive way.
//
// This algorithm iteratively compare every chars
// between p and t.
//
// This algorithm takes O(nm) time complexity,
// where n is the length of p(attern),
// and m is the length of t(ext).
func MatchNaively(p, t string) bool {
	m := len(p)
	n := len(t)

	i := 0
	j := 0

	for i < m && j < n {
		if p[i] == t[j] {
			i++
			j++
		} else {
			j = j - i + 1
			i = 0
		}
	}

	return i == m
}
