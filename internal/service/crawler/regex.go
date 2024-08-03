package crawler

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
)

var reDomain *regexp.Regexp

func setDomainRegex(parsed *url.URL) {
	if parsed == nil {
		reDomain = reDomainDefault
		return
	}

	splitted := strings.Split(parsed.Host, hostSeparator)

	var domainName, domainExtension, domainExtensionSecond string

	domainName = splitted[0]

	if domainName == www {
		domainName = splitted[1]
		domainExtension = splitted[2]
		domainExtensionSecond = splitted[3]

		reDomain = regexp.MustCompile(fmt.Sprintf(regexMultiDomainPattern, domainName, domainExtension, domainExtensionSecond))
		return
	}

	if len(splitted) == 2 {
		domainExtension = splitted[1]

		reDomain = regexp.MustCompile(fmt.Sprintf(regexSimpleDomainPattern, domainName, domainExtension))
		return
	}

	if len(splitted) > 2 {
		domainExtension = splitted[1]
		domainExtensionSecond = splitted[2]
	}

	reDomain = regexp.MustCompile(fmt.Sprintf(regexMultiDomainPattern, domainName, domainExtension, domainExtensionSecond))
}
