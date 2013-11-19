package stackr

import (
	. "github.com/ricallinson/simplebdd"
	"net/http"
	"testing"
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
