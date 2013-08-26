package stackr

import(
    "errors"
    "testing"
    . "github.com/ricallinson/simplebdd"
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
            req.Raw.Method = "GET"
            test := false
            app.Use("", Logger())
            app.Use("", func(req *Request, res *Response, next func()) {
                res.End("")
                test = true
            })
            app.handle(req, res, 0)
            AssertEqual(test, true)
        })

        It("should return [GET / 200 0ms] with config object", func() {
            req.Raw.Method = "GET"
            test := ""
            writer := func(a ...interface{}) (int, error) {
                test = a[0].(string)
                return 0, errors.New("")
            }
            app.Use("", Logger(OptLog{Writer: writer}))
            app.Use("", func(req *Request, res *Response, next func()) {
                res.End("")
            })
            app.handle(req, res, 0)
            AssertEqual(test, "\x1b[90mGET / \x1b[32m200 \x1b[90m0ms\x1b[0m")
        })

        It("should return [GET / 300 0ms]", func() {
            req.Raw.Method = "GET"
            test := ""
            writer := func(a ...interface{}) (int, error) {
                test = a[0].(string)
                return 0, errors.New("")
            }
            app.Use("", Logger(OptLog{Writer: writer}))
            app.Use("", func(req *Request, res *Response, next func()) {
                res.StatusCode = 300
                res.End("")
            })
            app.handle(req, res, 0)
            AssertEqual(test, "\x1b[90mGET / \x1b[36m300 \x1b[90m0ms\x1b[0m")
        })

        It("should return [GET / 400 0ms]", func() {
            req.Raw.Method = "GET"
            test := ""
            writer := func(a ...interface{}) (int, error) {
                test = a[0].(string)
                return 0, errors.New("")
            }
            app.Use("", Logger(OptLog{Writer: writer}))
            app.Use("", func(req *Request, res *Response, next func()) {
                res.StatusCode = 400
                res.End("")
            })
            app.handle(req, res, 0)
            AssertEqual(test, "\x1b[90mGET / \x1b[33m400 \x1b[90m0ms\x1b[0m")
        })

        It("should return [GET / 500 0ms]", func() {
            req.Raw.Method = "GET"
            test := ""
            writer := func(a ...interface{}) (int, error) {
                test = a[0].(string)
                return 0, errors.New("")
            }
            app.Use("", Logger(OptLog{Writer: writer}))
            app.Use("", func(req *Request, res *Response, next func()) {
                res.StatusCode = 500
                res.End("")
            })
            app.handle(req, res, 0)
            AssertEqual(test, "\x1b[90mGET / \x1b[31m500 \x1b[90m0ms\x1b[0m")
        })

        It("should return [GET / 500 0ms - 100]", func() {
            req.Raw.Method = "GET"
            test := ""
            writer := func(a ...interface{}) (int, error) {
                test = a[0].(string)
                return 0, errors.New("")
            }
            app.Use("", Logger(OptLog{Writer: writer}))
            app.Use("", func(req *Request, res *Response, next func()) {
                res.StatusCode = 500
                res.SetHeader("content-length", "100")
                res.End("")
            })
            app.handle(req, res, 0)
            AssertEqual(test, "\x1b[90mGET / \x1b[31m500 \x1b[90m0ms - 100\x1b[0m")
        })
    })

    Report(t)
}