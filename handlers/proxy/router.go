package proxy

import (
	"github.com/aadillion/personal-http-proxy/services/proxysrv"
	"github.com/go-chi/chi"
	"net/http"
)

func NewRouter(proxySrv proxysrv.Service) http.Handler {
	r := chi.NewRouter()
	h := handler{proxySrv}
	r.With(requestValidator).Post("/proxy", h.request)
	return r
}
