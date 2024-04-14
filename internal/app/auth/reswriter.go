package apiserver

import "net/http"

type ResWriter struct {
	http.ResponseWriter
	Code int
}

func (w *ResWriter) WriteHeader(statusCode int) {
	w.Code = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}
