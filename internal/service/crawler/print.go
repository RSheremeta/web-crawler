package crawler

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/RSheremeta/web-crawler/internal/service/http"
)

func (s *CrawlerService) PrintAllLinks(ctx context.Context, url string) {
	dataChan := make(chan string)
	errChan := make(chan error, 1)

	var wg sync.WaitGroup

	wg.Add(1)
	go s.ExtractLinks(ctx, url, dataChan, errChan, &wg)

	go func() {
		wg.Wait()
		close(dataChan)
		close(errChan)
	}()

	for {
		select {
		case <-time.After(s.ctxTimeout):
			s.log.Infof("Timeout reached, no more data received.")
			return

		case <-ctx.Done():
			s.log.Infof("Context done.")
			return

		case err, ok := <-errChan:
			if ok && err != nil {
				switch err {
				case ErrLinkAlreadyProcessed, ErrNilParsedBody:
					// do nothing
				case http.ErrRateLimitExceeded, http.ErrServiceUnavailable:
					s.log.Infof("Target website doesn't respond, so tearing down")
					return
				default:
					s.log.Fatalf("error: %s", err)
					return
				}
			}
			return
		case data, ok := <-dataChan:
			if !ok {
				s.log.Infof("Successfully printed %d links", len(s.linkMap.storage))
				return
			}

			fmt.Println(data)
		}
	}
}
