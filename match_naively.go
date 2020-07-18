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

	i := 1
	j := 1

	for i <= m && j <= n {
		if p[i-1] == t[j-1] {
			i++
			j++
		} else {
			j = j - i + 2
			i = 1
		}
	}

	return i == m+1
}
