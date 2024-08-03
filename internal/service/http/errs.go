package http

import "fmt"

var (
	ErrRateLimitExceeded  = fmt.Errorf("rate limit exceeded")
	ErrServiceUnavailable = fmt.Errorf("service is unavailable")
	ErrBrokenLink         = fmt.Errorf("the link is broken")
)
