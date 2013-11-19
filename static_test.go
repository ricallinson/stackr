package stackr

import (
	. "github.com/ricallinson/simplebdd"
	"testing"
)

func TestStatic(t *testing.T) {

	Describe("Static()", func() {

		var app *Server
		var req *Request
		var res *Response

		BeforeEach(func() {
			app = CreateServer()
			req = createRequest(NewMockHttpRequest())
			res = createResponse(NewMockResponseWriter(false))
		})

		It("should return [404]", func() {
			app.Use("", Static())
			app.Handle(req, res, 0)
			AssertEqual(res.StatusCode, 404)
		})

		It("should return [404]", func() {
			app.Use("", Static(map[string]string{}))
			app.Handle(req, res, 0)
			AssertEqual(res.StatusCode, 404)
		})

		It("should return [404] from a directory", func() {
			req.OriginalUrl = "/directory/"
			app.Use("", Static(map[string]string{"root": "./fixtures/"}))
			app.Handle(req, res, 0)
			AssertEqual(res.StatusCode, 404)
		})

		It("should return [404] from a directory without a trailing slash", func() {
			req.OriginalUrl = "/directory"
			app.Use("", Static(map[string]string{"root": "./fixtures/"}))
			app.Handle(req, res, 0)
			AssertEqual(res.StatusCode, 404)
		})

		It("should return [200] from a file", func() {
			req.OriginalUrl = "/text.txt"
			app.Use("", Static(map[string]string{"root": "./fixtures/"}))
			app.Handle(req, res, 0)
			AssertEqual(res.StatusCode, 200)
		})

		It("should return [404] from a directory on /public path", func() {
			req.OriginalUrl = "/public/directory/"
			app.Use("/public", Static(map[string]string{"root": "./fixtures/"}))
			app.Handle(req, res, 0)
			AssertEqual(res.StatusCode, 404)
		})

		It("should return [200] from a file on /public path", func() {
			req.OriginalUrl = "/public/text.txt"
			app.Use("/public", Static(map[string]string{"root": "./fixtures/"}))
			app.Handle(req, res, 0)
			AssertEqual(res.StatusCode, 200)
		})

		It("should return [404] from a directory", func() {
			req.OriginalUrl = "/directory/"
			app.Use("", Static(map[string]string{"root": "./fixtures"}))
			app.Handle(req, res, 0)
			AssertEqual(res.StatusCode, 404)
		})

		It("should return [200] from a opt.Root with no trailing slash", func() {
			req.OriginalUrl = "/text.txt"
			app.Use("", Static(map[string]string{"root": "./fixtures"}))
			app.Handle(req, res, 0)
			AssertEqual(res.StatusCode, 200)
		})

		It("should return [200] from a cached file", func() {
			req.OriginalUrl = "/text.txt"
			app.Use("", Static(map[string]string{"root": "./fixtures"}))
			app.Handle(req, res, 0)
			app.Handle(req, res, 0)
			app.Handle(req, res, 0)
			AssertEqual(res.StatusCode, 200)
		})
	})

	Report(t)
}
