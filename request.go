package stackr

import(
    "net/http"
)

/*
    A Request represents an HTTP request received by the server.
*/
type Request struct {

    // The standard http.Request type
    *http.Request

    // The value of .URL.RequestURI() for easy access.
    // Note: this value may be changed by the Stackr.handle() function.
    Url string

    // Set to the vlue of the matched portion of the .URL.RequestURI()
    MatchedUrl string

    // The value of .URL.RequestURI() for easy access.
    // Note: this value should NEVER be changed.
    OriginalUrl string
}

/*
    Returns a new Request.
*/

func createRequest(raw *http.Request) (*Request) {

    /*
        Create a new Request.
    */

    req := &Request{raw, raw.URL.RequestURI(), "", raw.URL.RequestURI()}

    /*
        Return the finished stack.Request.
    */

    return req
}