package crawler

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
)

func setDomainRegex(parsed *url.URL) {
	if parsed == nil {
		reDomain = reDomainDefault
		return
	}

	parts := strings.Split(parsed.Host, hostSeparator)
	if parts[0] == www {
		parts = parts[1:]
	}

	var domain, extFirst, extSecond string

	domain = parts[0]

	if len(parts) == 2 {
		extFirst = parts[1]

		reDomain = regexp.MustCompile(fmt.Sprintf(regexSimpleDomainPattern, domain, extFirst))
		return
	}

	if len(parts) > 2 {
		extFirst = parts[1]
		extSecond = parts[2]
	}

	reDomain = regexp.MustCompile(fmt.Sprintf(regexMultiDomainPattern, domain, extFirst, extSecond))
}
