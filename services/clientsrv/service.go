package clientsrv

import (
	"context"
	"github.com/sethgrid/pester"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type clientSrv struct {
	httpClient *pester.Client
}

func (s *clientSrv) Do(ctx context.Context, url string, method string, headers map[string]string) (status int, respHeaders map[string]string, contentLength int, err error) {
	req, err := http.NewRequestWithContext(ctx, method, url, nil)
	if err != nil {
		return
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := s.httpClient.Do(req)
	if err != nil {
		return
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)
	b, _ := ioutil.ReadAll(resp.Body)
	contentLength = len(b)
	status = resp.StatusCode
	respHeaders = make(map[string]string)
	for k, v := range resp.Header {
		respHeaders[strings.ToLower(k)] = v[0]
	}
	return
}

func NewClientSrv(httpClient *pester.Client) Service {
	return &clientSrv{httpClient}
}
