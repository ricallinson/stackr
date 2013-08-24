package stack

import(
    "errors"
    "net/http"
)

/*
    Create a Mock http.Request for testing.
*/

func NewMockHttpRequest() (*http.Request) {
    req := &http.Request{RequestURI: "/"}
    return req
}

/*
    Create a Mock http.ResponseWriter for testing.
*/

type MockResponseWriter struct {
    error bool
    headers http.Header
}

func (this *MockResponseWriter) Header() (http.Header) {
    return this.headers
}

func (this *MockResponseWriter) Write(data []byte) (int, error) {
    if this.error {
        return 0, errors.New("")
    }
    return len(data), nil
}

func (this *MockResponseWriter) WriteHeader(code int) {
    return
}

func NewMockResponseWriter(error bool) (*MockResponseWriter) {
    return &MockResponseWriter{error, make(http.Header)}
}