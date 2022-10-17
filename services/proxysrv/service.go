package proxysrv

import (
	"context"
	"crypto/sha256"
	"fmt"
	"github.com/aadillion/personal-http-proxy/services/clientsrv"
	"github.com/aadillion/personal-http-proxy/services/storesrv"
	"github.com/google/uuid"
	"sort"
	"strings"
)

type proxySrv struct {
	clientSrv clientsrv.Service
	storeSrv  storesrv.Service
}

func (s *proxySrv) Request(ctx context.Context, url string, method string, headers map[string]string) (pr ProcessedRequest, err error) {
	pr.Id = uuid.NewString()
	if pr.Status, pr.Headers, pr.Length, err = s.clientSrv.Do(ctx, url, method, headers); err != nil {
		return
	}
	keyHash := s.getHashByFields(url, method, headers)
	if err = s.storeSrv.Save(ctx, keyHash, pr); err != nil {
		return
	}

	return
}

func (s *proxySrv) getHashByFields(url string, method string, headers map[string]string) (hash string) {
	sb := strings.Builder{}
	sb.WriteString(url)
	sb.WriteString(method)
	if len(headers) > 0 {
		keys := make([]string, 0, len(headers))
		for k := range headers {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		for _, k := range keys {
			sb.WriteString(k)
			sb.WriteString(headers[k])
		}

	}

	return fmt.Sprintf("%x", sha256.Sum256([]byte(sb.String())))
}

func NewProxySrv(clientSrv clientsrv.Service, storeSrv storesrv.Service) Service {
	return &proxySrv{clientSrv, storeSrv}
}
