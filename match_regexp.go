package strcmp

import (
	"regexp"
)

// MatchRegexp implements matching algorithm with regexp package
func MatchRegexp(p, t string) bool {
	var reg = regexp.MustCompile(p)
	return reg.MatchString(t)
}
