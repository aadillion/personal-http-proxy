package services

import (
	"github.com/sethgrid/pester"
	"net/http"
	"time"
)

func NewHttpClient() *pester.Client {
	client := pester.New()
	client.Transport = &http.Transport{
		Proxy:                 http.ProxyFromEnvironment,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          7,
		MaxIdleConnsPerHost:   10,
		IdleConnTimeout:       10 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	client.MaxRetries = 5
	client.Backoff = pester.ExponentialBackoff
	return client
}
