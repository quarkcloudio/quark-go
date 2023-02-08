package builder

import (
	"io"
	"net/http"
)

type Response struct {
	statusCode int
	h          http.Header
	w          io.Writer
}

func NewResponse(writer io.Writer) http.ResponseWriter {

	return &Response{w: writer}
}

func (w *Response) StatusCode() int {
	if w.statusCode == 0 {
		return http.StatusOK
	}
	return w.statusCode
}

func (w *Response) Header() http.Header {
	if w.h == nil {
		w.h = make(http.Header)
	}
	return w.h
}

func (w *Response) WriteHeader(statusCode int) {
	w.statusCode = statusCode
}

func (w *Response) Write(p []byte) (int, error) {
	return w.w.Write(p)
}

func (w *Response) Flush() {}
