package connect

import(
    // "fmt"
)

type FavOpt struct {}

func faviconHandler(req *Request, res *Response, next func()) {
    
    if req.Url == "/favicon.ico" {
        res.End("")
    }
}

func Favicon(opt FavOpt) (func(req *Request, res *Response, next func())) {
    return faviconHandler
}