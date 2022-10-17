package proxysrv

import "context"

type Service interface {
	Request(ctx context.Context, url string, method string, headers map[string]string) (pr ProcessedRequest, err error)
}
