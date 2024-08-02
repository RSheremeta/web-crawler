package http

import "regexp"

const (
	anchorTag     = "a"
	linkAttribute = "href"
)

// todo - make dynamic ?
var reDomain = regexp.MustCompile(`^(?:(?:https?)://)?(?:www.)?monzo\.com`)

// todo - invoke in main
func SetDomainRegex(domain string) {
	// reDomain
}
