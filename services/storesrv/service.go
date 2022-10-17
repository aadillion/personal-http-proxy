package storesrv

import (
	"context"
	"fmt"
	"sync"
)

type storeSrv struct {
	m *sync.Map
}

func (s *storeSrv) Save(ctx context.Context, key string, value interface{}) (err error) {
	select {
	default:
	case <-ctx.Done():
		return ctx.Err()
	}

	s.m.Store(key, value)

	i := 0
	s.m.Range(func(key, value interface{}) bool {
		fmt.Printf("\t[%d] key: %v, value: %v\n", i, key, value)
		i++
		return true
	})
	return
}

func NewStoreSrv(m *sync.Map) Service {
	return &storeSrv{m}
}
