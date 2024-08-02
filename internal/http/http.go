package http

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

type HttpClient struct {
	cl  http.Client
	log *logrus.Entry
}

func NewHttpClient(logger *logrus.Entry) *HttpClient {
	return &HttpClient{
		cl:  http.Client{},
		log: logger.WithField("client", "http"),
	}
}
