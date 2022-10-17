package storesrv

import "context"

type Service interface {
	Save(ctx context.Context, key string, value interface{}) (err error)
}
