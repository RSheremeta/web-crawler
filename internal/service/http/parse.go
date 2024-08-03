package http

import (
	"context"
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

func (s *HttpService) ParseHTML(ctx context.Context, url string) (*html.Node, error) {
	logger := s.log.WithField("url", url)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequestWithContext: %w", err)
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http.client.Do: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {

		switch resp.StatusCode {
		case http.StatusTooManyRequests:
			return nil, ErrRateLimitExceeded

		case http.StatusInternalServerError:
			return nil, ErrServiceUnavailable

		case http.StatusNotFound:
			return nil, ErrBrokenLink
		}

		return nil, fmt.Errorf("response code err: %d", resp.StatusCode)
	}

	parsed, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("html.Parse: %w", err)
	}

	logger.Debugln("html body parsed successfully")

	return parsed, nil
}
