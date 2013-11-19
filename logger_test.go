package stackr

import (
	"errors"
	. "github.com/ricallinson/simplebdd"
	"testing"
	"time"
)

func TestLogger(t *testing.T) {

	Describe("Logger()", func() {

		var app *Server
		var req *Request
		var res *Response

		BeforeEach(func() {
			app = CreateServer()
			req = createRequest(NewMockHttpRequest())
			res = createResponse(NewMockResponseWriter(false))
		})

		It("should return [GET / 200 0ms]", func() {
			req.Method = "GET"
			test := false
			app.Use("", Logger())
			app.Use("", func(req *Request, res *Response, next func()) {
				res.End("")
				test = true
			})
			app.Handle(req, res, 0)
			AssertEqual(test, true)
		})

		It("should return [GET / 200 0ms] with config object", func() {
			req.Method = "GET"
			test := ""
			writer := func(a ...interface{}) (int, error) {
				test = a[0].(string)
				return 0, errors.New("")
			}
			app.Use("", Logger(map[string]string{"format": "dev"}, writer))
			app.Use("", func(req *Request, res *Response, next func()) {
				res.End("")
			})
			app.Handle(req, res, 0)
			AssertEqual(test, "\x1b[90mGET / \x1b[32m200 \x1b[90m0ms\x1b[0m")
		})

		It("should return [GET / 300 0ms]", func() {
			req.Method = "GET"
			test := ""
			writer := func(a ...interface{}) (int, error) {
				test = a[0].(string)
				return 0, errors.New("")
			}
			app.Use("", Logger(map[string]string{"format": "dev"}, writer))
			app.Use("", func(req *Request, res *Response, next func()) {
				res.StatusCode = 300
				res.End("")
			})
			app.Handle(req, res, 0)
			AssertEqual(test, "\x1b[90mGET / \x1b[36m300 \x1b[90m0ms\x1b[0m")
		})

		It("should return [GET / 400 0ms]", func() {
			req.Method = "GET"
			test := ""
			writer := func(a ...interface{}) (int, error) {
				test = a[0].(string)
				return 0, errors.New("")
			}
			app.Use("", Logger(map[string]string{"format": "dev"}, writer))
			app.Use("", func(req *Request, res *Response, next func()) {
				res.StatusCode = 400
				res.End("")
			})
			app.Handle(req, res, 0)
			AssertEqual(test, "\x1b[90mGET / \x1b[33m400 \x1b[90m0ms\x1b[0m")
		})

		It("should return [GET / 500 0ms]", func() {
			req.Method = "GET"
			test := ""
			writer := func(a ...interface{}) (int, error) {
				test = a[0].(string)
				return 0, errors.New("")
			}
			app.Use("", Logger(map[string]string{"format": "dev"}, writer))
			app.Use("", func(req *Request, res *Response, next func()) {
				res.StatusCode = 500
				res.End("")
			})
			app.Handle(req, res, 0)
			AssertEqual(test, "\x1b[90mGET / \x1b[31m500 \x1b[90m0ms\x1b[0m")
		})

		It("should return [GET / 500 0ms - 100]", func() {
			req.Method = "GET"
			test := ""
			writer := func(a ...interface{}) (int, error) {
				test = a[0].(string)
				return 0, errors.New("")
			}
			app.Use("", Logger(map[string]string{"format": "dev"}, writer))
			app.Use("", func(req *Request, res *Response, next func()) {
				res.StatusCode = 500
				res.SetHeader("content-length", "100")
				res.End("")
			})
			app.Handle(req, res, 0)
			AssertEqual(test, "\x1b[90mGET / \x1b[31m500 \x1b[90m0ms - 100\x1b[0m")
		})

		It("should return [GET - 200]", func() {
			req.Method = "GET"
			test := ""
			writer := func(a ...interface{}) (int, error) {
				test = a[0].(string)
				return 0, errors.New("")
			}
			app.Use("", Logger(map[string]string{"format": ":method - :status - :res[content-length]"}, writer))
			app.Use("", func(req *Request, res *Response, next func()) {
				res.StatusCode = 500
				res.SetHeader("content-length", "100")
				res.End("")
			})
			app.Handle(req, res, 0)
			AssertEqual(test, "GET - 500 - 100")
		})
	})

	Describe("loggerFormatFunctions[match]()", func() {

		var opt *loggerOpt
		var req *Request
		var res *Response

		BeforeEach(func() {
			opt = &loggerOpt{}
			req = createRequest(NewMockHttpRequest())
			res = createResponse(NewMockResponseWriter(false))
		})

		It("should return [0] from :res[content-length]", func() {
			result := loggerFormatFunctions[":res[content-length]"](opt, req, res)
			AssertEqual(result, "0")
		})

		It("should return [100] from :res[content-length]", func() {
			res.SetHeader("content-length", "100")
			result := loggerFormatFunctions[":res[content-length]"](opt, req, res)
			AssertEqual(result, "100")
		})

		It("should return [HTTP/1.1] from :http-version]", func() {
			req.Proto = "HTTP/1.1"
			result := loggerFormatFunctions[":http-version"](opt, req, res)
			AssertEqual(result, "HTTP/1.1")
		})

		It("should return [1] from :response-time]", func() {
			/*
			   This test is time based and will come back to bite me.
			*/
			opt.startTime = time.Now().UnixNano() - 1000000
			result := loggerFormatFunctions[":response-time"](opt, req, res)
			AssertEqual(result, "1")
		})

		It("should return [foo] from :remote-addr]", func() {
			req.RemoteAddr = "foo"
			result := loggerFormatFunctions[":remote-addr"](opt, req, res)
			AssertEqual(result, "foo")
		})

		It("should return [25] from :date]", func() {
			result := loggerFormatFunctions[":date"](opt, req, res)
			AssertEqual(len(result), 31)
		})

		It("should return [GET] from :method]", func() {
			req.Method = "GET"
			result := loggerFormatFunctions[":method"](opt, req, res)
			AssertEqual(result, "GET")
		})

		It("should return [/] from :url]", func() {
			req.OriginalUrl = "/"
			result := loggerFormatFunctions[":url"](opt, req, res)
			AssertEqual(result, "/")
		})

		It("should return [] from :referrer]", func() {
			result := loggerFormatFunctions[":referrer"](opt, req, res)
			AssertEqual(result, "")
		})

		It("should return [] from :user-agent]", func() {
			result := loggerFormatFunctions[":user-agent"](opt, req, res)
			AssertEqual(result, "")
		})

		It("should return [200] from :status]", func() {
			result := loggerFormatFunctions[":status"](opt, req, res)
			AssertEqual(result, "200")
		})
	})

	Report(t)
}
