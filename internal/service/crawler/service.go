package crawler

import (
	"context"

	"github.com/RSheremeta/web-crawler/internal/config"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/html"
)

type HttpService interface {
	ParseHTML(ctx context.Context, url string) (*html.Node, error)
}

type CrawlerService struct {
	log              *logrus.Entry
	defaultTargetURL string

	linkMap *LinkMap

	httpService HttpService
}

func NewCrawlerService(
	cfg *config.Config,
	logger *logrus.Entry,
	httpService HttpService,
) *CrawlerService {
	return &CrawlerService{
		log:              logger.WithField("service", "crawler"),
		defaultTargetURL: cfg.DefaultTargetURL,

		linkMap: newLinkMap(),

		httpService: httpService,
	}
}
