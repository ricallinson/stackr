package stackr

import (
	"bytes"
	. "github.com/ricallinson/simplebdd"
	"strings"
	"testing"
)

func TestErrorHandler(t *testing.T) {

	Describe("ErrorHandler()", func() {

		var app *Server
		var req *Request
		var mock *MockResponseWriter
		var res *Response

		BeforeEach(func() {
			app = CreateServer()
			req = createRequest(NewMockHttpRequest())
			mock = NewMockResponseWriter(false)
			res = createResponse(mock)
		})

		It("should return [nil]", func() {
			app.Use(ErrorHandler())
			app.Use(func(req *Request, res *Response, next func()) {
				res.End("nil")
			})
			app.Handle(req, res, 0)
			w := bytes.NewBuffer(mock.Written).String()
			AssertEqual(w, "nil")
		})

		It("should return [panic]", func() {
			app.Env = "prod"
			app.Use(ErrorHandler())
			app.Use(func(req *Request, res *Response, next func()) {
				panic("panic")
			})
			app.Handle(req, res, 0)
			w := bytes.NewBuffer(mock.Written).String()
			AssertEqual(w, "panic")
		})

		It("should return [34]", func() {
			app.Env = "prod"
			app.Use(ErrorHandler())
			app.Use(func(req *Request, res *Response, next func()) {
				panic("panic")
			})
			req.Header.Set("Accept", "text/html")
			app.Handle(req, res, 0)
			w := bytes.NewBuffer(mock.Written).String()
			AssertEqual(strings.Index(w, "<title>panic</title>"), 34)
		})

		It("should return [67]", func() {
			app.Env = "prod"
			app.Use(ErrorHandler("Title"))
			app.Use(func(req *Request, res *Response, next func()) {
				panic("panic")
			})
			req.Header.Set("Accept", "text/html")
			app.Handle(req, res, 0)
			w := bytes.NewBuffer(mock.Written).String()
			AssertEqual(strings.Index(w, "<h1>Title</h1>"), 67)
		})

		It("should return [{\"code\":\"500\",\"error\":\"panic\"}]", func() {
			app.Env = "prod"
			app.Use(ErrorHandler())
			app.Use(func(req *Request, res *Response, next func()) {
				panic("panic")
			})
			req.Header.Set("Accept", "application/json")
			app.Handle(req, res, 0)
			w := bytes.NewBuffer(mock.Written).String()
			AssertEqual(w, "{\"code\":\"500\",\"error\":\"panic\"}")
		})
	})

	Report(t)
}
