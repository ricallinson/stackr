package stackr

import (
	"errors"
	"net/http"
	"net/url"
)

/*
   Create a Mock http.Request for testing.
*/

func NewMockHttpRequest() *http.Request {
	req := &http.Request{
		Header:     http.Header{},
		RequestURI: "/",
		URL:        new(url.URL),
	}
	return req
}

/*
   Create a Mock http.ResponseWriter for testing.
*/

type MockResponseWriter struct {
	error   bool
	headers http.Header
	Written []byte
}

func (this *MockResponseWriter) Header() http.Header {
	return this.headers
}

func (this *MockResponseWriter) Write(data []byte) (int, error) {
	if this.error {
		return 0, errors.New("")
	}
	this.Written = data
	return len(data), nil
}

func (this *MockResponseWriter) WriteHeader(code int) {
	return
}

func NewMockResponseWriter(error bool) *MockResponseWriter {
	return &MockResponseWriter{error, make(http.Header), []byte{}}
}
