package printer

import (
	"context"
	"sync"
	"time"

	"github.com/RSheremeta/web-crawler/config"
	"github.com/sirupsen/logrus"
)

type CrawlerService interface {
	ExtractLinks(
		ctx context.Context,
		url string,
		dataChan chan<- string,
		errChan chan<- error,
		wg *sync.WaitGroup,
	)
}

type PrinterService struct {
	log        *logrus.Entry
	ctxTimeout time.Duration

	crawlerService CrawlerService
}

func NewPrinterService(
	cfg *config.Config,
	logger *logrus.Entry,

	crawlerService CrawlerService,
) *PrinterService {
	return &PrinterService{
		log:        logger.WithField("service", "printer"),
		ctxTimeout: cfg.Printer.ContextTimeout,

		crawlerService: crawlerService,
	}
}
