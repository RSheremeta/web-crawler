package main

import (
	"context"
	"fmt"
	"time"

	"github.com/RSheremeta/web-crawler/internal/http"
	"github.com/RSheremeta/web-crawler/internal/logger"
)

// todo - cfg
var timeout = 60 * time.Second

// todo - default or from input
const targetURL = "https://monzo.com/"

func main() {
	ctx, cncl := context.WithTimeout(context.Background(), timeout)
	defer cncl()

	logger := logger.NewLoggerInstance()

	cl := http.NewHttpClient(logger)

	http.SetDomainRegex(targetURL)

	links, err := cl.ExtractLinks(ctx, targetURL)
	if err != nil {
		logger.Fatalf("cl.ExtractLinks: %s", err)
	}

	fmt.Println()
	for i := range links {
		fmt.Println(links[i])
	}
	fmt.Println()

	logger.Infof("total links len is %d", len(links))
}
