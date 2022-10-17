package proxy

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aadillion/personal-http-proxy/handlers"
	"golang.org/x/exp/slices"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const contextItemRequestValidated contextItem = "context_request_validated"

var allowedHTTPMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE"}

func requestValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			handlers.HandleError(w, err)
			return
		}
		var req requestInput

		if err = json.Unmarshal(body, &req); err != nil {
			handlers.HandleError(w, handlers.ErrBadJsonError)
			return
		}

		if len(req.Url) == 0 {
			handlers.HandleError(w, handlers.BadRequestError{Message: "field `url` is empty"})
			return
		} else if _, err = url.ParseRequestURI(req.Url); err != nil {
			handlers.HandleError(
				w, handlers.BadRequestError{Message: fmt.Sprintf("field `url` is invalid; %s", err)})
			return
		}

		if len(req.Method) == 0 {
			handlers.HandleError(w, handlers.BadRequestError{Message: "field `method` is empty"})
			return
		}
		req.Method = strings.ToUpper(req.Method)
		if !slices.Contains(allowedHTTPMethods, req.Method) {
			handlers.HandleError(w,
				handlers.BadRequestError{Message: fmt.Sprintf("field `method` is invalid; allowed methods: %s",
					strings.Join(allowedHTTPMethods, ", "))})
			return
		}

		ctx := context.WithValue(r.Context(), contextItemRequestValidated, &req)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
