package crawler

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/RSheremeta/web-crawler/internal/service/http"
	"golang.org/x/net/html"
)

func (s *CrawlerService) ExtractLinks(
	ctx context.Context,
	url string,
	dataChan chan<- string,
	errChan chan<- error,
	wg *sync.WaitGroup,
) {
	defer wg.Done()

	if url == "" {
		s.log.Infof("url param is empty, setting the default value of %q", s.defaultURL)
		url = s.defaultURL
	}

	if !s.linkMap.storeIfNotExists(url) {
		return
	}

	logger := s.log.WithField("url", url)

	parsed, err := s.parseURL(url)
	if err != nil {
		errChan <- fmt.Errorf("s.parseURL: %w", err)
	}

	setDomainRegex(parsed)

	parsedHTML, err := s.httpService.ParseHTML(ctx, url)
	if err != nil {
		if errors.Is(err, http.ErrRateLimitExceeded) || errors.Is(err, http.ErrServiceUnavailable) {
			errChan <- err
		}
		errChan <- fmt.Errorf("httpService.ParseHTML: %w", err)
	}

	if parsedHTML == nil {
		logger.Warnf(ErrNilParsedBody.Error())

		errChan <- ErrNilParsedBody
	}

	// s.linkMap.store(url)

	links := s.extractLinksFromHTML(parsedHTML)

	for i := range links {
		dataChan <- links[i]

		wg.Add(1)
		go s.ExtractLinks(ctx, links[i], dataChan, errChan, wg)
	}

	logger.Debugf("extracted %d links", len(links))
}

func (s *CrawlerService) extractLinksFromHTML(doc *html.Node) []string {
	var res []string

	if doc == nil {
		return res
	}

	if doc.Type == html.ElementNode && (doc.Data == anchorTag || doc.Data == linkTag) {
		for _, attr := range doc.Attr {
			if attr.Key == linkAttribute {
				if !s.isLinkSuitable(attr.Val) {
					continue
				}

				item := s.prettifyLink(attr.Val)

				if s.linkMap.exists(item) {
					continue
				}

				res = append(res, item)
			}
		}
	}

	for child := doc.FirstChild; child != nil; child = child.NextSibling {
		res = append(res, s.extractLinksFromHTML(child)...)
	}

	return res
}
