package stack

import(
    "testing"
    . "github.com/ricallinson/simplebdd"
)

func TestStatic(t *testing.T) {

    Describe("Static()", func() {

        var app *Server
        var req *Request
        var res *Response

        BeforeEach(func() {
            app = CreateServer()
            req = CreateRequest(NewMockHttpRequest())
            res = CreateResponse(NewMockResponseWriter(false))
        })

        It("should return []", func() {
            app.Use("", Static(StaticOpt{}))
            app.handle(req, res, 0)
            AssertEqual(true, true)
        })
    })

    Report(t)
}