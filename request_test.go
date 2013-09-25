package stackr

import(
    "testing"
    "net/http"
    . "github.com/ricallinson/simplebdd"
)

func TestRequest(t *testing.T) {

    var req *Request

    BeforeEach(func() {
        req = &Request{
            Request: &http.Request{
                Header: http.Header{},
            },
        }
    })

    Describe("Accepted()", func() {

        It("should return [application/xml]", func() {
            req.Header.Set("Accept", "text/html, application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
            a := req.Accepted()
            AssertEqual(a[2], "application/xml")
        })
    })

    Describe("AcceptedLanguages()", func() {

        It("should return [en-us]", func() {
            req.Header.Set("Accept-Language", "zh, en-us; q=0.8, en; q=0.6")
            a := req.AcceptedLanguages()
            AssertEqual(a[1], "en-us")
        })
    })

    Describe("AcceptedCharsets()", func() {

        It("should return [utf-8]", func() {
            req.Header.Set("Accept-Charset", "utf-8")
            a := req.AcceptedCharsets()
            AssertEqual(a[0], "utf-8")
        })
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

        It("should return [application/json]", func() {
            a := req.processAccepted("application/xml,application/xhtml+xml,text/html;q=0.9,text/plain;q=0.8,image/png,*/*;q=0.5")
            AssertEqual(a[2], "image/png")
        })
    })

    Report(t)
}