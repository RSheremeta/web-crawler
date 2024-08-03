package http

import (
	"context"
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

var (
	ErrRateLimitExceeded  = fmt.Errorf("rate limit exceeded")
	ErrServiceUnavailable = fmt.Errorf("service is unavailable")
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
		if resp.StatusCode == http.StatusTooManyRequests {
			return nil, ErrRateLimitExceeded
		}
		if resp.StatusCode == http.StatusInternalServerError {
			return nil, ErrServiceUnavailable
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
