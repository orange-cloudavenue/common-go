package regex

import "regexp"

const (
	// T0 name (pr01e02ocb0001234spt101)
	T0NameRegexString = `^pr(vrf)?(?<siteCode>[0-9]{2})?(?<workloadType>[a-z]{1})(?<contractId>[a-z0-9]{10})(?<serviceType>[a-z]{2,6})(?<linkType>[a-z]{2})(?<increment>[0-9]{2,3})`
	// edgegateway(t1) name (tn01e02ocb0001234spt101)
	EdgeGatewayNameRegexString = `^tn(?<siteCode>[0-9]{2})?(?<workloadType>[a-z]{1})(?<workload>[0-9]{2})(?<contractId>[a-z0-9]{10})(?<serviceType>[a-z]{2,6})t1(?<increment>[0-9]{2,5})`
)

type CavResourceName struct {
	Key         string
	Description string
	RegexString string
	RegexP      *regexp.Regexp
}

var (
	ListCavResourceNames = []CavResourceName{
		{
			Key:         "t0",
			Description: "T0 name (pr01e02ocb0001234spt101)",
			RegexString: T0NameRegexString,
			RegexP:      T0NameRegex(),
		},
		{
			Key:         "edgegateway",
			Description: "Edge Gateway name (tn01e02ocb0001234spt101)",
			RegexString: EdgeGatewayNameRegexString,
			RegexP:      EdgeGatewayNameRegex(),
		},
	}
)
