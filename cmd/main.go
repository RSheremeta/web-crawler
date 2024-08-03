package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/RSheremeta/web-crawler/config"
	"github.com/RSheremeta/web-crawler/internal/logger"
	"github.com/RSheremeta/web-crawler/internal/service/crawler"
	"github.com/RSheremeta/web-crawler/internal/service/http"
	"github.com/RSheremeta/web-crawler/internal/service/printer"
)

const (
	ctxTimeout = 60 * time.Second

	flagName  = "url"
	flagDescr = "the base url to be crawled on"
)

func main() {
	ctx, cncl := context.WithTimeout(context.Background(), ctxTimeout)
	defer cncl()

	log := logger.NewDefaultLogger()

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("couldn't init cfg %s", err)
	}

	log = logger.NewLogger(cfg)

	targetURL := flag.String(flagName, "", flagDescr)
	flag.Parse()

	httpSvc := http.NewHttpService(cfg, log)
	crawlerSvc := crawler.NewCrawlerService(cfg, log, *targetURL, httpSvc)
	printerSvc := printer.NewPrinterService(cfg, log, crawlerSvc)

	stopChan := make(chan os.Signal, 1)
	signal.Notify(
		stopChan,
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	go func() {
		sign := <-stopChan
		log.Infof("Captured exit signal %v, stopping...", sign.String())

		cncl()

		os.Exit(0)
	}()

	printerSvc.PrintAllLinks(ctx, *targetURL)
}
