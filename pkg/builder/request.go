package builder

import (
	"io"
	"net/http"
)

func NewRequest(method string, url string, body io.Reader) *http.Request {
	r, _ := http.NewRequest(method, url, body)

	return r
}
