package stackr

import(
    "reflect"
    "testing"
    . "github.com/ricallinson/simplebdd"
)

func TestHttp(t *testing.T) {

    Describe("createHttpHandler()", func() {

        var app *server

        BeforeEach(func() {
            app = CreateServer()
        })

        It("should return [*stack.Handler]", func() {
            app := CreateServer()
            handler := createHttpHandler(app)
            AssertEqual(reflect.TypeOf(handler).String(), "*stackr.handler")
        })
    })

    Describe("ServeHTTP()", func() {

        var app *server

        BeforeEach(func() {
            app = CreateServer()
        })

        It("should return [true]", func() {
            app := CreateServer()
            test := false
            app.Use("/", func(req *Request, res *Response, next func()) {
                test = true
            })
            handler := createHttpHandler(app)
            req := NewMockHttpRequest()
            res := NewMockResponseWriter(false)
            handler.ServeHTTP(res, req)
            AssertEqual(test, true)
        })
    })

    Report(t)
}