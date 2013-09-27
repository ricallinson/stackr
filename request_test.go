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

    Report(t)
}