/*
 * SPDX-FileCopyrightText: Copyright (c) 2025 Orange
 * SPDX-License-Identifier: Mozilla Public License 2.0
 *
 * This software is distributed under the MPL-2.0 license.
 * the text of which is available at https://www.mozilla.org/en-US/MPL/2.0/
 * or see the "LICENSE" file for more details.
 */

package regex

import "regexp"

const (
	// T0 name (pr01e02ocb0001234spt101)
	T0NameRegexString = `^pr(vrf)?(?<siteCode>[0-9]{2})?(?<workloadType>[a-z]{1})(?<contractId>[a-z0-9]{10})(?<serviceType>[a-z]{2,6})(?<linkType>[a-z]{2})(?<increment>[0-9]{2,3})`
	// edgegateway(t1) name (tn01e02ocb0001234spt101)
	EdgeGatewayNameRegexString = `^tn(?<siteCode>[0-9]{2})?(?<workloadType>[a-z]{1})(?<workload>[0-9]{2})(?<contractId>[a-z0-9]{10})(?<serviceType>[a-z]{2,6})t1(?<increment>[0-9]{2,5})`
	// Organization name (cav01ev01ocb0001234) - https://regex101.com/r/hgYkSE/1
	OrganizationNameRegexString = `^cav(?<siteCode>[0-9]{2})(?<customerType>[i,e,v])v([0-9]{2})ocb(?<contractId>[0-9]{7})`
	// VDC name (<alphanumeric> with hyphen and minus, with max length 27 and min length 2) - https://regex101.com/r/NgL6X0/1
	VDCNameRegexString = `^[a-zA-Z0-9-_]{2,27}$`
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
		{
			Key:         "organization",
			Description: "Organization name (cav01ev01ocb0001234)",
			RegexString: OrganizationNameRegexString,
			RegexP:      OrganizationNameRegex(),
		},
		{
			Key:         "vdc",
			Description: "VDC name (<alphanumeric> with - _ character and with max length 27 and min length 2)",
			RegexString: VDCNameRegexString,
			RegexP:      VDCNameRegex(),
		},
	}
)
