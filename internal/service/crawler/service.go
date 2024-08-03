package crawler

import (
	"context"

	"github.com/sirupsen/logrus"
	"golang.org/x/net/html"

	"github.com/RSheremeta/web-crawler/config"
)

type HttpService interface {
	ParseHTML(ctx context.Context, url string) (*html.Node, error)
}

type CrawlerService struct {
	log        *logrus.Entry
	defaultURL string

	linkMap *LinkMap

	httpService HttpService
}

func NewCrawlerService(
	cfg *config.Config,
	logger *logrus.Entry,
	targetURL string,
	httpService HttpService,
) *CrawlerService {
	return &CrawlerService{
		log: logger.WithField("service", "crawler"),
		defaultURL: func() string {
			if targetURL == "" {
				return cfg.DefaultTargetURL
			}

			return targetURL
		}(),

		linkMap: newLinkMap(),

		httpService: httpService,
	}
}

func (s *CrawlerService) GetProcessedCount() int {
	return len(s.linkMap.storage)
}
