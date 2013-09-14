package stackr

import(
    "net/http"
)

/*
    A HTTP Request.

    Access to the http.Request attributes is possible.

    .Method string
    .Proto string
    .ProtoMajor int
    .ProtoMinor int
    .Header http.Header
    .Body io.ReadCloser
    .ContentLength int64
    .TransferEncoding []string
    .Close bool
    .Host string
    .Form url.Values
    .PostForm url.Values
    .MultipartForm *multipart.Form
    .Trailer http.Header
    .RemoteAddr string
    .RequestURI string
    .TLS *tls.ConnectionState
*/
type Request struct {
    *http.Request
    Url string
    MatchedUrl string
    OriginalUrl string
}

/*
    Returns a new Request.
*/

func createRequest(raw *http.Request) (*Request) {

    /*
        Create a new Request.
    */

    req := &Request{raw, "", "", ""}

    /*
        Set the source http.Request so it can be accessed later.
    */

    // req. = raw

    /*
        Set the Url for easy access.

        Note: this value may be changed by the stack.handle() function.
    */

    req.Url = raw.URL.RequestURI()

    /*
        Set the Url for easy access.

        Note: this value should never change over the life time of the request.
    */

    req.OriginalUrl = req.Url

    /*
        Return the finished stack.Request.
    */

    return req
}