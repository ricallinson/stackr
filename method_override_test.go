package stackr

import (
	. "github.com/ricallinson/simplebdd"
	"testing"
)

func TestMethodOverride(t *testing.T) {

	Describe("MethodOverride()", func() {

		var app *Server
		var req *Request
		var res *Response

		BeforeEach(func() {
			app = CreateServer()
			req = createRequest(NewMockHttpRequest())
			res = createResponse(NewMockResponseWriter(false))
		})

		It("should return [GET]", func() {
			app.Use(MethodOverride())
			req.Method = "GET"
			app.Handle(req, res, 0)
			AssertEqual(req.Method, "GET")
		})

		It("should return [POST]", func() {
			app.Use(MethodOverride())
			req.Method = "GET"
			req.Header.Set("X-HTTP-Method-Override", "POST")
			app.Handle(req, res, 0)
			AssertEqual(req.Method, "POST")
		})

		It("should return [HEAD]", func() {
			app.Use(MethodOverride())
			req.Method = "HEAD"
			req.Header.Set("X-HTTP-Method-Override", "POST")
			app.Handle(req, res, 0)
			AssertEqual(req.Map["OriginalMethod"], "HEAD")
		})

		It("should return [DELETE]", func() {
			app.Use(MethodOverride())
			req.Method = "GET"
			req.Header.Set("x-http-method-override", "delete")
			app.Handle(req, res, 0)
			AssertEqual(req.Method, "DELETE")
		})

		It("should return [POST]", func() {
			app.Use(MethodOverride())
			req.Method = "PUT"
			req.Header.Set("x-http-method-override", "foo")
			app.Handle(req, res, 0)
			AssertEqual(req.Method, "PUT")
		})
	})

	Report(t)
}
