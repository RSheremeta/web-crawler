package crawler

import (
	"context"
	"fmt"

	"golang.org/x/net/html"
)

func (s *CrawlerService) ExtractLinks(ctx context.Context, url string) ([]string, error) {
	if url == "" {
		s.log.Infof("url param is empty, setting the default value of %q", s.defaultTargetURL)
		url = s.defaultTargetURL
	}

	logger := s.log.WithField("url", url)

	parsed, err := s.parseURL(url)
	if err != nil {
		return nil, fmt.Errorf("s.parseURL: %w", err)
	}

	setDomainRegex(parsed)

	parsedHTML, err := s.httpService.ParseHTML(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("httpService.ParseHTML: %w", err)
	}

	if parsedHTML == nil {
		logger.Warnf("parsedHTML is nil")
		return []string{}, nil
	}

	links := s.extractLinksFromHTML(parsedHTML, url)

	logger.Debugf("extracted %d links", len(links))

	// go func()

	return links, nil
}

func (s *CrawlerService) extractLinksFromHTML(doc *html.Node, targetURL string) []string {
	var res []string

	if doc.Type == html.ElementNode && doc.Data == anchorTag {
		for _, attr := range doc.Attr {
			if attr.Key == linkAttribute {
				if !s.isLinkSuitable(targetURL, attr.Val) {
					continue
				}

				res = append(res, attr.Val)
			}
		}
	}

	for child := doc.FirstChild; child != nil; child = child.NextSibling {
		res = append(res, s.extractLinksFromHTML(child, targetURL)...)
	}

	return res
}
