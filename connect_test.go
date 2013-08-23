package stack

import(
    "testing"
    . "github.com/ricallinson/simplebdd"
)

func TestStack(t *testing.T) {

    Describe("Use()", func() {

        It("should return 1", func() {

            app := CreateServer()

            app.Use("/", func(req *Request, res *Response, next func()) {
                res.End("Hello world\n")
            })

            AssertEqual(len(app.stack), 1)
        })

        It("should return 2", func() {

            app := CreateServer()

            app.Use("/", func(req *Request, res *Response, next func()) {
                res.End("Hello world\n")
            })

            app.Use("/", func(req *Request, res *Response, next func()) {
                res.End("Hello world\n")
            })

            AssertEqual(len(app.stack), 2)
        })
    })

    Report(t)
}