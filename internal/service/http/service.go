package http

import (
	"net/http"

	"github.com/RSheremeta/web-crawler/internal/config"
	"github.com/sirupsen/logrus"
)

type HttpService struct {
	log    *logrus.Entry
	client *http.Client
}

func NewHttpService(cfg *config.Config, logger *logrus.Entry) *HttpService {
	return &HttpService{
		log: logger.WithField("service", "http"),

		client: &http.Client{
			Timeout: cfg.Http.Timeout,
		},
	}
}
