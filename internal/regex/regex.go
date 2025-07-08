package regex

import (
	"regexp"
	"sync"
)

const (
	uUID4RegexString        = `[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}`
	uRNWithUUID4RegexString = `(?m)urn:[A-Za-z0-9][A-Za-z0-9-]{0,31}:([A-Za-z0-9()+,\-.:=@;$_!*']|%[0-9A-Fa-f]{2})+` + uUID4RegexString
)

func lazyRegexCompile(str string) func() *regexp.Regexp {
	var regex *regexp.Regexp
	var once sync.Once
	return func() *regexp.Regexp {
		once.Do(func() {
			regex = regexp.MustCompile(str)
		})
		return regex
	}
}

var (
	UUID4Regex        = lazyRegexCompile(uUID4RegexString)
	URNWithUUID4Regex = lazyRegexCompile(uRNWithUUID4RegexString)
)
