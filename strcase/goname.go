package strcase

import (
	"strings"
	"unicode"

	"github.com/kr/pretty"
)

// ToPublicGoName returns the Go public name of the given string.
func ToPublicGoName(s string) string {
	return toGoName(s)
	// return toGoName(TitleFirstWord(s))
}

// ToPrivateGoName returns the Go private name of the given string.
func ToPrivateGoName(s string) string {
	return toGoName(lowerCaseFirstLetterOrAcronyms(s))
}

// toGoName returns a different name if it should be different.
// Now supports "{index}" and "{key}" as special path segments.
func toGoName(name string) (should string) {
	name = strings.ReplaceAll(name, " ", "_")
	name = strings.ReplaceAll(name, "-", "_")

	// Fast path for simple cases: "_" and all lowercase.
	if name == "_" {
		return name
	}
	allLower := true
	for _, r := range name {
		if !unicode.IsLower(r) {
			allLower = false
			break
		}
	}
	if allLower {
		return name
	}

	// Handle special segments "{index}" and "{key}" (leave them as is)
	parts := strings.Split(name, ".")
	for i, part := range parts {
		if part == "{index}" || part == "{key}" {
			continue
		}
		// Process the part as before
		runes := []rune(part)
		w, j := 0, 0 // index of start of word, scan
		for j+1 <= len(runes) {
			eow := false // whether we hit the end of a word
			switch {
			case j+1 == len(runes):
				eow = true
			case runes[j+1] == '_':
				eow = true
				n := 1
				for j+n+1 < len(runes) && runes[j+n+1] == '_' {
					n++
				}
				if j+n+1 < len(runes) && unicode.IsDigit(runes[j]) && unicode.IsDigit(runes[j+n+1]) {
					n--
				}
				copy(runes[j+1:], runes[j+n+1:])
				runes = runes[:len(runes)-n]
			case unicode.IsLower(runes[j]) && !unicode.IsLower(runes[j+1]):
				eow = true
			}
			j++
			if !eow {
				continue
			}

			// [w,j) is a word.
			word := string(runes[w:j])
			u := strings.ToUpper(word)
			if commonInitialisms[u] {
				// Keep consistent case, which is lowercase only at the start.
				if w == 0 && unicode.IsLower(runes[w]) {
					u = strings.ToLower(u)
				}
				copy(runes[w:], []rune(u))
			} else if specialCase, exist := customInitialisms[u]; exist {
				if w == 0 && unicode.IsLower(runes[w]) {
					u = specialCase[1]
				} else {
					u = specialCase[0]
				}
				copy(runes[w:], []rune(u))
			} else if w > 0 && strings.ToLower(word) == word {
				// already all lowercase, and not the first word, so uppercase the first character.
				runes[w] = unicode.ToUpper(runes[w])
			}
			w = j
		}
		parts[i] = string(runes)
		// Always force the first letter to uppercase for each part (except {index}/{key})
		if len(parts[i]) > 0 {
			r := []rune(parts[i])
			r[0] = unicode.ToUpper(r[0])
			parts[i] = string(r)
		}
	}
	return strings.Join(parts, ".")
}

// Deprecated: this list should not be completed as it affects generation for our Go SDK only.
//
// commonInitialisms is a set of common initialisms.
// Only add entries that are highly unlikely to be non-initialisms.
// For instance, "ID" is fine (Freudian code is rare), but "AND" is not.
var commonInitialisms = map[string]bool{
	"ACL":   true,
	"API":   true,
	"ASCII": true,
	"CPU":   true,
	"CSS":   true,
	"DHCP":  true,
	"DNS":   true,
	"EOF":   true,
	"GUID":  true,
	"HTML":  true,
	"HTTP":  true,
	"HTTPS": true,
	"ID":    true,
	"IP":    true,
	"JSON":  true,
	"LB":    true,
	"LHS":   true,
	"QPS":   true,
	"RAM":   true,
	"RHS":   true,
	"RPC":   true,
	"SLA":   true,
	"SMTP":  true,
	"SQL":   true,
	"SSD":   true,
	"SSH":   true,
	"TCP":   true,
	"TLS":   true,
	"TTL":   true,
	"UDP":   true,
	"UI":    true,
	"UID":   true,
	"UUID":  true,
	"URI":   true,
	"URL":   true,
	"UTF8":  true,
	"VM":    true,
	"XML":   true,
	"XMPP":  true,
	"XSRF":  true,
	"XSS":   true,
}

// customInitialisms is a set of common initialisms we use at Cloudavenue.
// value[0] is the uppercase replacement
// value[1] is the lowercase replacement
var customInitialisms = map[string][2]string{
	"ACLS":  {"ACLs", "acls"},
	"APIS":  {"APIs", "apis"},
	"CPUS":  {"CPUs", "cpus"},
	"IDS":   {"IDs", "ids"},
	"IPS":   {"IPs", "ips"},
	"IPV":   {"IPv", "ipv"}, // handle IPV4 && IPV6
	"LBS":   {"LBs", "lbs"},
	"UIDS":  {"UIDs", "uids"},
	"UUIDS": {"UUIDs", "uuids"},
	"URIS":  {"URIs", "uris"},
	"URLS":  {"URLs", "urls"},
}

// TitleFirstWord upper case the first letter of a string.
func TitleFirstWord(s string) string {
	pretty.Println("=================================", s)
	if s == "" {
		return s
	}

	// split the string by dot and capitalize the first letter of each part (if not {index} or {key})
	parts := strings.Split(s, ".")
	for i, part := range parts {
		if part == "{index}" || part == "{key}" {
			continue
		}
		parts[i] = titleFirstWord(part)
	}
	return strings.Join(parts, ".")
}

func titleFirstWord(s string) string {
	if s == "" {
		return s
	}

	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])

	return string(r)
}

// UntitleFirstWord lower case the first letter of a string.
func UntitleFirstWord(s string) string {
	if s == "" {
		return s
	}

	r := []rune(s)

	firstWord := strings.Split(s, " ")[0]
	_, isCommonInitialism := commonInitialisms[firstWord]
	_, isCustomInitialism := customInitialisms[firstWord]
	if !isCommonInitialism && !isCustomInitialism {
		r[0] = unicode.ToLower(r[0])
	}

	return string(r)
}

// lowerCaseFirstLetterOrAcronyms lower case the first letter of a string.
func lowerCaseFirstLetterOrAcronyms(s string) string {
	r := []rune(s)
	if len(r) == 0 {
		return ""
	}

	for i := 0; len(r) > i && unicode.IsUpper(r[i]); i++ {
		word := string(r[:i+1])
		if u := strings.ToUpper(word); commonInitialisms[u] {
			copy(r[0:], []rune(strings.ToLower(u)))
			break
		}
	}
	r[0] = unicode.ToLower(r[0])

	return string(r)
}
