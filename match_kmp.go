package main

// MatchKMP implements Knuth-Morris-Pratt algorithm
// @doc https://en.wikipedia.org/wiki/Knuth%E2%80%93Morris%E2%80%93Pratt_algorithm
func MatchKMP(p, t string) bool {
	m := len(p)
	n := len(t)

	i := 1
	j := 1

	failures := computeKMPFailures(p)

	for i <= m && j <= n {
		if i == 0 || p[i-1] == t[j-1] {
			i++
			j++
		} else {
			i = failures[i-1]
		}
	}

	return i == m+1
}

func computeKMPFailures(p string) []int {
	m := len(p)
	failures := make([]int, m)

	i := 0
	j := 1

	for j < m {
		if i == 0 || p[i-1] == p[j-1] {
			i++
			j++
			failures[j-1] = i
		} else {
			i = failures[i-1]
		}
	}

	return failures
}
