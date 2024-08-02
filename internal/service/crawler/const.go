package crawler

import "regexp"

const emptyPath = "/"

const (
	anchorTag     = "a"
	linkAttribute = "href"
)

const hostSeparator = "."

const regexSimpleDomainPattern = `^https:\/\/%s\.%s\/.*$`
const regexMultiDomainPattern = `^https:\/\/%s\.%s\/%s\/.*$`

var reDomainDefault = regexp.MustCompile(`^(?:(?:https?)://)?(?:www\.)?monzo\.com(/.*)?$`)
