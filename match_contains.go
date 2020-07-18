package strcmp

import "strings"

// MatchContains implement string match algorithm using "strings" package.
//
// @see strings.Contains https://github.com/golang/go/blob/0951939fd9e4a6bc83f23c42e8ddff02b29c997e/src/strings/strings.go#L60-L63
// @see bytes.Index https://github.com/golang/go/blob/0951939fd9e4a6bc83f23c42e8ddff02b29c997e/src/bytes/bytes.go#L1076-L1162
func MatchContains(p, t string) bool {
	return strings.Contains(t, p)
}
