package stack

import(
    "net/http"
)

/*
    Create a Mock http.Request for testing.
*/

func NewMockHttpRequest() (*http.Request) {
    req := &http.Request{}
    return req
}

/*
    Create a Mock http.ResponseWriter for testing.
*/

type MockResponseWriter struct {}

func (this *MockResponseWriter) Header() (http.Header) {
    return make(http.Header)
}

func (this *MockResponseWriter) Write([]byte) (int, error) {
    return 0, nil
}

func (this *MockResponseWriter) WriteHeader(code int) {
    return
}

func NewMockResponseWriter() (*MockResponseWriter) {
    return new(MockResponseWriter)
}