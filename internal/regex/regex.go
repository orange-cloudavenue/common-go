package regex

import (
	"regexp"
	"sync"
)

const (
	UUID4RegexString        = `[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}`
	URNWithUUID4RegexString = `(?m)urn:[A-Za-z0-9][A-Za-z0-9-]{0,31}:([A-Za-z0-9()+,\-.:=@;$_!*']|%[0-9A-Fa-f]{2})+` + UUID4RegexString

	// * Cloudavenue
	// edgegateway(t1) name (tn01e02ocb0001234spt101)
	T0NameRegexString          = `^pr(vrf)?(?<siteCode>[0-9]{2})?(?<workloadType>[a-z]{1})(?<contractId>[a-z0-9]{10})(?<serviceType>[a-z]{2,6})(?<linkType>[a-z]{2})(?<increment>[0-9]{2,3})`
	EdgeGatewayNameRegexString = `^tn(?<siteCode>[0-9]{2})?(?<workloadType>[a-z]{1})(?<workload>[0-9]{2})(?<contractId>[a-z0-9]{10})(?<serviceType>[a-z]{2,6})t1(?<increment>[0-9]{2,5})`
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
	UUID4Regex           = lazyRegexCompile(UUID4RegexString)
	URNWithUUID4Regex    = lazyRegexCompile(URNWithUUID4RegexString)
	T0NameRegex          = lazyRegexCompile(T0NameRegexString)
	EdgeGatewayNameRegex = lazyRegexCompile(EdgeGatewayNameRegexString)
)
