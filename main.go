package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"github.com/RSheremeta/web-crawler/internal/config"
	"github.com/RSheremeta/web-crawler/internal/logger"
	"github.com/RSheremeta/web-crawler/internal/service/crawler"
	"github.com/RSheremeta/web-crawler/internal/service/http"
)

const ctxTimeout = 60 * time.Second

func main() {
	ctx, cncl := context.WithTimeout(context.Background(), ctxTimeout)
	defer cncl()

	log := logger.NewDefaultLogger()

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("couldn't init cfg %s", err)
	}

	log = logger.NewLogger(cfg)

	targetURL := flag.String("url", "", "a base url to be crawled on")
	flag.Parse()

	httpSvc := http.NewHttpService(cfg, log)
	crawlerSvc := crawler.NewCrawlerService(cfg, log, httpSvc)

	links, err := crawlerSvc.ExtractLinks(ctx, *targetURL)
	if err != nil {
		log.Fatalf("crawlerSvc.ExtractLinks: %s", err)
	}

	fmt.Println()
	for i := range links {
		fmt.Println(links[i])
	}
	fmt.Println()

	fmt.Println("targetURL:", *targetURL)

	log.Infof("total links len is %d", len(links))
}
