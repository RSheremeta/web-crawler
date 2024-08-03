package crawler

import (
	"fmt"
	"net/url"
	"strings"
)

func (s *CrawlerService) parseURL(rawURL string) (*url.URL, error) {
	res, err := url.ParseRequestURI(rawURL)
	if err != nil {
		s.log.Errorf("passed url %s is invalid", rawURL)
		return nil, fmt.Errorf("url.ParseRequestURI: %w", err)
	}

	return res, nil
}

func (s *CrawlerService) isLinkSuitable(link string) bool {
	link = s.prettifyLink(link)

	r := reDomain
	_ = r

	b := reDomain.MatchString(strings.ToLower(link))
	return b
}

func (s *CrawlerService) prettifyLink(link string) string {
	if strings.HasPrefix(link, "/") && link != emptyPath {
		link = s.defaultURL + link
	}

	return link
}
