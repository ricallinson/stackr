package stack

import(
    "strings"
    "net/http"
)

/*
    Options for the static middleware.
*/

type StaticOpt struct {
    Root string
    Listings bool
}

func Static(opt StaticOpt) (func(req *Request, res *Response, next func())) {

    root := "./static"

    if len(opt.Root) > 0 {
        root = opt.Root
    }

    listings := false

    if opt.Listings {
        listings = true
    }

    // Return the handler function.
    return func(req *Request, res *Response, next func()) {
        if listings == false && strings.HasSuffix(req.Url, "/") {
            res.End("")
            return
        }
        http.StripPrefix(req.UrlMatched, http.FileServer(http.Dir(root))).ServeHTTP(res.Writer, req.Raw)
        res.End("")
    }
}