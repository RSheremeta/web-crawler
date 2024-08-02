package http

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func (c *HttpClient) ExtractLinks(ctx context.Context, url string) ([]string, error) {
	c.log = c.log.WithField("url", url)

	if !c.isValidURL(url) {
		return nil, fmt.Errorf("invalid url")
	}

	parsedHTML, err := c.fetchParsedHTML(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("fetchParsedHTML: %w", err)
	}

	if parsedHTML == nil {
		c.log.Warnf("parsedHTML is nil")
		return []string{}, nil
	}

	links := c.extractLinksFromHTML(parsedHTML, url)

	c.log.Infof("total: %d", len(links))

	return links, nil
}

func (c *HttpClient) isValidURL(rawURL string) bool {
	_, err := url.Parse(rawURL)
	if err != nil {
		c.log.Errorf("passed url %s is invalid", rawURL)
		return false
	}

	return true
}

func (c *HttpClient) fetchParsedHTML(ctx context.Context, url string) (*html.Node, error) {
	req, err := http.NewRequest(http.MethodGet, url, http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequest: %w", err)
	}

	req = req.WithContext(ctx)

	resp, err := c.cl.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http.client.Do() %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("response code err: %d", resp.StatusCode)
	}

	parsed, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("html.Parse: %w", err)
	}

	return parsed, nil
}

func (c *HttpClient) extractLinksFromHTML(doc *html.Node, targetURL string) []string {
	var res []string

	if doc.Type == html.ElementNode && doc.Data == anchorTag {
		for _, attr := range doc.Attr {
			if attr.Key == linkAttribute {
				if !c.isLinkSuitable(targetURL, attr.Val) {
					continue
				}

				res = append(res, attr.Val)
			}
		}
	}

	for child := doc.FirstChild; child != nil; child = child.NextSibling {
		res = append(res, c.extractLinksFromHTML(child, targetURL)...)
	}

	return res
}

func (c *HttpClient) isLinkSuitable(url, link string) bool {
	// meaning it's an inner route with no domain
	if strings.HasSuffix(link, "/") && link != "/" {
		link = url + link // todo - prettify url parsing
	}

	// we're interested in links of the target domain only, and not the external ones
	if !reDomain.MatchString(strings.ToLower(link)) {
		return false
	}

	return true
}
