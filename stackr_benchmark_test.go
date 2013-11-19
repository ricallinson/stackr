package stackr

import (
	// "fmt"
	"testing"
)

func BenchmarkMatchOneHandle(b *testing.B) {

	req := createRequest(NewMockHttpRequest())
	res := createResponse(NewMockResponseWriter(false))

	app := CreateServer()
	app.Use("/", func(req *Request, res *Response, next func()) {})

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		app.Handle(req, res, 0)
	}
}

func BenchmarkMatchTenHandles(b *testing.B) {

	req := createRequest(NewMockHttpRequest())
	res := createResponse(NewMockResponseWriter(false))

	app := CreateServer()

	for i := 0; i < 10; i++ {
		app.Use("/", func(req *Request, res *Response, next func()) {})
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		app.Handle(req, res, 0)
	}
}

func BenchmarkMatchFiftyHandles(b *testing.B) {

	req := createRequest(NewMockHttpRequest())
	res := createResponse(NewMockResponseWriter(false))

	app := CreateServer()

	for i := 0; i < 50; i++ {
		app.Use("/", func(req *Request, res *Response, next func()) {})
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		app.Handle(req, res, 0)
	}
}

func BenchmarkMatchThousandHandles(b *testing.B) {

	req := createRequest(NewMockHttpRequest())
	res := createResponse(NewMockResponseWriter(false))

	app := CreateServer()

	for i := 0; i < 1000; i++ {
		app.Use("/", func(req *Request, res *Response, next func()) {})
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		app.Handle(req, res, 0)
	}
}
