package stackr

import(
    "testing"
    . "github.com/ricallinson/simplebdd"
)

func TestRequest(t *testing.T) {

    var req *Request

    BeforeEach(func() {
        req = &Request{}
    })

    Describe("processAccepted()", func() {

        It("should return [text/plain]", func() {
            a := req.processAccepted("text/plain")
            AssertEqual(a[0], "text/plain")
        })

        It("should return [application/json]", func() {
            a := req.processAccepted("text/plain, text/html,application/json , image/png")
            AssertEqual(a[2], "application/json")
        })
    })

    Report(t)
}