package connect

import(
    // "fmt"
)

func faviconHandler(req *Request, res *Response, next func()) {
    
    if req.Url == "/favicon.ico" {
        res.End("")
    }
}

func Favicon() (func(req *Request, res *Response, next func())) {
    return faviconHandler
}