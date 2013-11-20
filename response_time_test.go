package stackr

import (
	. "github.com/ricallinson/simplebdd"
	"testing"
)

func TestResponseTime(t *testing.T) {

	Describe("ResponseTime()", func() {

		var app *Server
		var req *Request
		var res *Response

		BeforeEach(func() {
			app = CreateServer()
			req = createRequest(NewMockHttpRequest())
			res = createResponse(NewMockResponseWriter(false))
		})

		It("should return [3]", func() {
			app.Use(ResponseTime())
			app.Handle(req, res, 0)
			AssertEqual(len(res.Writer.Header().Get("X-Response-Time")), 3)
		})

		It("should return [3]", func() {
			app.Use(ResponseTime())
			app.Use(func(req *Request, res *Response, next func()) {
				res.Write("")
			})
			app.Handle(req, res, 0)
			AssertEqual(len(res.Writer.Header().Get("X-Response-Time")), 3)
		})

		It("should return [3]", func() {
			app.Use(ResponseTime())
			app.Use(func(req *Request, res *Response, next func()) {
				res.End("")
			})
			app.Handle(req, res, 0)
			AssertEqual(len(res.Writer.Header().Get("X-Response-Time")), 3)
		})

		It("should return [3, true]", func() {
			app.Use(ResponseTime())
			app.Use(func(req *Request, res *Response, next func()) {
				res.SetHeader("X-Set", "true")
				res.End("")
			})
			app.Handle(req, res, 0)
			AssertEqual(res.Writer.Header().Get("X-Set"), "true")
		})
	})

	Report(t)
}
