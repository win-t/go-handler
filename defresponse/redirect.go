package defresponse

import (
	"net/http"

	handler "github.com/payfazz/go-handler"
)

// Redirect as Response.
func Redirect(status int, url string) handler.Response {
	return handler.Response{
		Status: status,
		Executor: func(resp handler.Response, w http.ResponseWriter, r *http.Request) error {
			handler.MergeHeader(w.Header(), resp.Header)
			http.Redirect(w, r, url, resp.Status)
			return nil
		},
	}
}
