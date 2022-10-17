package proxy

import (
	"encoding/json"
	"github.com/aadillion/personal-http-proxy/handlers"
	"github.com/aadillion/personal-http-proxy/services/proxysrv"
	"net/http"
)

type handler struct {
	proxySrv proxysrv.Service
}

func (h *handler) request(w http.ResponseWriter, r *http.Request) {
	reqIn := r.Context().Value(contextItemRequestValidated).(*requestInput)

	pr, err := h.proxySrv.Request(r.Context(), reqIn.Url, reqIn.Method, reqIn.Headers)
	if err != nil {
		handlers.HandleError(w, err)
		return
	}
	reqOut := requestOutput{
		pr.Id,
		pr.Status,
		pr.Headers,
		pr.Length,
	}
	b, _ := json.Marshal(reqOut)
	w.Header().Add("content-type", "application/json")
	_, _ = w.Write(b)
}
