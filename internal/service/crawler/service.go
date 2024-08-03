package crawler

import (
	"context"
	"time"

	"github.com/RSheremeta/web-crawler/config"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/html"
)

type HttpService interface {
	ParseHTML(ctx context.Context, url string) (*html.Node, error)
}

type CrawlerService struct {
	log        *logrus.Entry
	defaultURL string

	linkMap *LinkMap

	httpService HttpService

	ctxTimeout time.Duration
}

func NewCrawlerService(
	cfg *config.Config,
	logger *logrus.Entry,
	defaultURL string,
	httpService HttpService,
) *CrawlerService {
	return &CrawlerService{
		log: logger.WithField("service", "crawler"),
		defaultURL: func() string {
			if defaultURL == "" {
				return cfg.DefaultTargetURL
			}

			return defaultURL
		}(),

		linkMap: newLinkMap(),

		httpService: httpService,

		ctxTimeout: cfg.Crawler.ContextTimeout,
	}
}
