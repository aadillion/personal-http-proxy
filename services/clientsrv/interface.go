package clientsrv

import "context"

type Service interface {
	Do(ctx context.Context, url string, method string, headers map[string]string) (status int, respHeaders map[string]string, contentLength int, err error)
}
