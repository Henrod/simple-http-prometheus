package prometheus

import "net/http"

type ResponseWriter struct {
	httpResponseWriter http.ResponseWriter
	statusCode         int
}

func NewResponseWriter(w http.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{httpResponseWriter: w, statusCode: http.StatusOK}
}

func (w *ResponseWriter) Header() http.Header {
	return w.httpResponseWriter.Header()
}

func (w *ResponseWriter) Write(bytes []byte) (int, error) {
	return w.httpResponseWriter.Write(bytes) //nolint:wrapcheck
}

func (w *ResponseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.httpResponseWriter.WriteHeader(statusCode)
}
