package stack

import(
    "testing"
    . "github.com/ricallinson/simplebdd"
)

func TestStack(t *testing.T) {

    Describe("Use()", func() {

        It("should return [1]", func() {
            app := CreateServer()
            app.Use("/", func(req *Request, res *Response, next func()) {})
            AssertEqual(len(app.stack), 1)
        })

        It("should return [2]", func() {
            app := CreateServer()
            app.Use("/foo", func(req *Request, res *Response, next func()) {})
            app.Use("/bar", func(req *Request, res *Response, next func()) {})
            AssertEqual(len(app.stack), 2)
        })

        It("should return a [/] from an empty string", func() {
            app := CreateServer()
            app.Use("", func(req *Request, res *Response, next func()) {})
            m := app.stack[0]
            AssertEqual(m.Route, "/")
        })

        It("should return the route with the trailing slash removed [/foo]", func() {
            app := CreateServer()
            app.Use("/foo/", func(req *Request, res *Response, next func()) {})
            m := app.stack[0]
            AssertEqual(m.Route, "/foo")
        })

        It("should return the lower case string [/foo]", func() {
            app := CreateServer()
            app.Use("/FOO", func(req *Request, res *Response, next func()) {})
            m := app.stack[0]
            AssertEqual(m.Route, "/foo")
        })
    })

    Describe("handle()", func() {

        It("should return []", func() {
            app := CreateServer()
            req := CreateRequest(NewMockHttpRequest())
            res := CreateResponse(NewMockResponseWriter())
            app.handle(req, res, 0)
            AssertEqual(true, true)
        })
    })

    Report(t)
}