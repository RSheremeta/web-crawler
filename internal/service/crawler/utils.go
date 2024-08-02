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

func (s *CrawlerService) isLinkSuitable(url, link string) bool {
	if strings.HasSuffix(link, "/") && link != emptyPath {
		link = url + link // todo - prettify url parsing
	}

	if !reDomain.MatchString(strings.ToLower(link)) {
		return false
	}

	return true
}
