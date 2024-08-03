package crawler

import (
	"fmt"
	"regexp"
)

const (
	www           = "www"
	hostSeparator = "."
)

const (
	anchorTag     = "a"
	linkAttribute = "href"
)

const emptyPath = "/"

const regexSimpleDomainPattern = `^(?:(?:https?)://)?(?:www\.)?%s\.%s\/.*$`
const regexMultiDomainPattern = `^(?:(?:https?)://)?(?:www\.)?%s\.%s\.%s\/.*$`

var reDomainDefault = regexp.MustCompile(`^(?:(?:https?)://)?(?:www\.)?monzo\.com(/.*)?$`)

var (
	ErrLinkAlreadyProcessed = fmt.Errorf("link already processed")
	ErrNilParsedBody        = fmt.Errorf("parsed html body is nil")
)
