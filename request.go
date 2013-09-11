package stackr

import(
    "net/http"
)

/*
    A HTTP Request.

    Access to the following http.Request attributes is possible via the Http attribute.

    .Http.Method string
    .Http.Proto string
    .Http.ProtoMajor int
    .Http.ProtoMinor int
    .Http.Header http.Header
    .Http.Body io.ReadCloser
    .Http.ContentLength int64
    .Http.TransferEncoding []string
    .Http.Close bool
    .Http.Host string
    .Http.Form url.Values
    .Http.PostForm url.Values
    .Http.MultipartForm *multipart.Form
    .Http.Trailer http.Header
    .Http.RemoteAddr string
    .Http.RequestURI string
    .Http.TLS *tls.ConnectionState
*/
type Request struct {
    Http *http.Request
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

    req := new(Request)

    /*
        Set the source http.Request so it can be accessed later.
    */

    req.Http = raw

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