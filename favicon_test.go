package stackr

import(
    "testing"
    . "github.com/ricallinson/simplebdd"
)

func TestFavicon(t *testing.T) {

    Describe("Favicon()", func() {

        var app *server
        var req *Request
        var res *Response

        BeforeEach(func() {
            app = CreateServer()
            req = createRequest(NewMockHttpRequest())
            res = createResponse(NewMockResponseWriter(false))
        })

        It("should return [false]", func() {
            app.Use("", Favicon(FavOpt{}))
            app.handle(req, res, 0)
            AssertNotEqual(res.Writer.Header().Get("content-type"), "image/x-icon")
        })

        It("should return [false]", func() {
            req.OriginalUrl = "/favicon.ic"
            app.Use("", Favicon(FavOpt{}))
            app.handle(req, res, 0)
            AssertNotEqual(res.Writer.Header().Get("content-type"), "image/x-icon")
        })

        It("should return [text/plain] from not found", func() {
            req.OriginalUrl = "/favicon.ico"
            app.Use("", Favicon(FavOpt{}))
            app.handle(req, res, 0)
            AssertEqual(res.Writer.Header().Get("content-type"), "text/plain")
        })

        It("should return [image/x-icon]", func() {
            req.OriginalUrl = "/favicon.ico"
            app.Use("", Favicon(FavOpt{Path: "./fixtures/favicon.ico"}))
            app.handle(req, res, 0)
            AssertEqual(res.Writer.Header().Get("content-type"), "image/x-icon")
        })

        It("should return [image/x-icon] from cache", func() { // checked on coverage report
            req.OriginalUrl = "/favicon.ico"
            app.Use("", Favicon(FavOpt{Path: "./fixtures/favicon.ico"}))
            app.handle(req, res, 0)
            res = createResponse(NewMockResponseWriter(false))
            app.handle(req, res, 0)
            AssertEqual(res.Writer.Header().Get("content-type"), "image/x-icon")
        })

        It("should return [public, max-age=1]", func() {
            req.OriginalUrl = "/favicon.ico"
            app.Use("", Favicon(FavOpt{Path: "./fixtures/favicon.ico", MaxAge: 1000}))
            app.handle(req, res, 0)
            AssertEqual(res.Writer.Header().Get("cache-control"), "public, max-age=1")
        })
    })

    Report(t)
}