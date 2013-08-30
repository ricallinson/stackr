package stackr

import(
    "testing"
    . "github.com/ricallinson/simplebdd"
)

func TestStack(t *testing.T) {

    Describe("Use()", func() {

        var app *Server

        BeforeEach(func() {
            app = CreateServer()
        })

        It("should return [1]", func() {
            app.Use("/", func(req *Request, res *Response, next func()) {})
            AssertEqual(len(app.stack), 1)
        })

        It("should return [2]", func() {
            app.Use("/foo", func(req *Request, res *Response, next func()) {})
            app.Use("/bar", func(req *Request, res *Response, next func()) {})
            AssertEqual(len(app.stack), 2)
        })

        It("should return a [/] from an empty string", func() {
            app.Use("", func(req *Request, res *Response, next func()) {})
            m := app.stack[0]
            AssertEqual(m.Route, "/")
        })

        It("should return the route with the trailing slash removed [/foo]", func() {
            app.Use("/foo/", func(req *Request, res *Response, next func()) {})
            m := app.stack[0]
            AssertEqual(m.Route, "/foo")
        })

        It("should return the lower case string [/foo]", func() {
            app.Use("/FOO", func(req *Request, res *Response, next func()) {})
            m := app.stack[0]
            AssertEqual(m.Route, "/foo")
        })
    })

    Describe("handle()", func() {

        var app *Server
        var req *Request
        var res *Response

        BeforeEach(func() {
            app = CreateServer()
            req = createRequest(NewMockHttpRequest())
            res = createResponse(NewMockResponseWriter(false))
        })

        It("should return [true] after default function is called", func() {
            test := false
            app.Use("/", func(req *Request, res *Response, next func()) {
                test = true
            })
            app.handle(req, res, 0)
            AssertEqual(test, true)
        })

        It("should return [false] from url /foo", func() {
            test := false
            app.Use("/foo", func(req *Request, res *Response, next func()) {
                test = true
            })
            app.handle(req, res, 0)
            AssertEqual(test, false)
        })

        It("should return [true] from url /foo", func() {
            test := false
            req.OriginalUrl = "/foo"
            app.Use("/foo", func(req *Request, res *Response, next func()) {
                test = true
            })
            app.handle(req, res, 0)
            AssertEqual(test, true)
        })

        It("should return [true] from url /foo/bar", func() {
            test := false
            req.OriginalUrl = "/foo/bar"
            app.Use("/foo", func(req *Request, res *Response, next func()) {
                test = true
            })
            app.handle(req, res, 0)
            AssertEqual(test, true)
        })

        It("should return [true] from url /foo/bar/ with trialling slash", func() {
            test := false
            req.OriginalUrl = "/foo/bar/"
            app.Use("/foo/bar", func(req *Request, res *Response, next func()) {
                test = true
            })
            app.handle(req, res, 0)
            AssertEqual(test, true)
        })

        It("should return [true] from url /foo/bar with trialling slash on route", func() {
            test := false
            req.OriginalUrl = "/foo/bar"
            app.Use("/foo/bar/", func(req *Request, res *Response, next func()) {
                test = true
            })
            app.handle(req, res, 0)
            AssertEqual(test, true)
        })

        It("should return [false] from url /foo/bar with double trialling slash on route", func() {
            test := false
            req.OriginalUrl = "/foo/bar"
            app.Use("/foo/bar//", func(req *Request, res *Response, next func()) {
                test = true
            })
            app.handle(req, res, 0)
            AssertEqual(test, false)
        })

        It("should return [2] from url /foo", func() {
            test := 0
            req.OriginalUrl = "/foo/bar"
            app.Use("/foo", func(req *Request, res *Response, next func()) {
                test++
            })
            app.Use("/foo", func(req *Request, res *Response, next func()) {
                test++
            })
            app.handle(req, res, 0)
            AssertEqual(test, 2)
        })

        It("should return [1] from url /foo/bar", func() {
            test := 0
            req.OriginalUrl = "/foo/bar"
            app.Use("/foo", func(req *Request, res *Response, next func()) {
                test++
                res.End("")
            })
            app.Use("/foo", func(req *Request, res *Response, next func()) {
                test++
            })
            app.handle(req, res, 0)
            AssertEqual(test, 1)
        })

        It("should return [404] as nothing is matched", func() {
            req.Http.Method = "HEAD"
            app.handle(req, res, 0)
            AssertEqual(res.StatusCode, 404)
        })

        It("should return [true] from route without a call to res.End()", func() {
            test := false
            app.Use("/", func(req *Request, res *Response, next func()) {
                res.Write("foo")
                test = true
            })
            app.handle(req, res, 0)
            AssertEqual(test, true)
        })

        It("should return [firstsecond] after calling next() in route handler", func() {
            test := ""
            app.Use("/", func(req *Request, res *Response, next func()) {
                next()
                test += "second"
            })
            app.Use("/", func(req *Request, res *Response, next func()) {
                test += "first"
            })
            app.handle(req, res, 0)
            AssertEqual(test, "firstsecond")
        })

        It("should return [false] as the writer throws an error", func() {
            test := true
            res = createResponse(NewMockResponseWriter(true))
            app.Use("/", func(req *Request, res *Response, next func()) {
                test = res.Write("foo")
            })
            app.handle(req, res, 0)
            AssertEqual(test, false)
        })

        It("should return [false] after trying to set a header after data is sent", func() {
            test := true
            app.Use("/", func(req *Request, res *Response, next func()) {
                res.Write("foo")
            })
            app.Use("/", func(req *Request, res *Response, next func()) {
                test = res.SetHeader("key", "val")
            })
            app.handle(req, res, 0)
            AssertEqual(test, false)
        })
    })

    Report(t)
}