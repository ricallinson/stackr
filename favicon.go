package stack

import(
    "io"
    "fmt"
    "io/ioutil"
    "crypto/md5"
)

type FavOpt struct {
    Path string
    MaxAge int
}

/*
    Favicon:

    By default serves the stack favicon, or the favicon located by the given `path`.

    Options:

        - `maxAge`  cache-control max-age directive, defaulting to 1 day

    Examples:

        Serve default favicon:

        stack.CreateServer().Use(stack.Favicon(stack.FavOpt{}))

    Serve favicon before logging for brevity:

        app := stack.CreateServer()
        app.use(stack.Favicon(stack.FavOpt{}))
        app.use(stack.Logger(stack.LogOpt{}))

    Serve custom favicon:
    
        stack.CreateServer().Use(stack.Favicon(stack.FavOpt{path "public/favicon.ico"}))
 */

func Favicon(opt FavOpt) (func(req *Request, res *Response, next func())) {

    type Icon struct {
        headers map[string]string
        body []byte
    }

    icon := Icon{
        headers: make(map[string]string),
    }

    maxAge := 86400000

    if opt.MaxAge > 0 {
        maxAge = opt.MaxAge
    }

    path := "./favicon.ico"

    if len(opt.Path) > 0 {
        path = opt.Path
    }

    return func(req *Request, res *Response, next func()) {

        // If this is not a fav icon return fast
        if req.Url != "/favicon.ico" {
            return
        }

        // If we have the icon cached, serve it.
        if len(icon.body) > 0 {
            res.SetHeaders(icon.headers)
            res.WriteBytes(icon.body)
            return
        }

        // Otherwise read the icon into cache.
        buf, err := ioutil.ReadFile(path)
        if err != nil {
            return
        }

        hasher := md5.New()
        io.WriteString(hasher, fmt.Sprint(buf))

        icon.headers["content-type"] = "image/x-icon"
        icon.headers["content-length"] = fmt.Sprint(len(buf))
        icon.headers["etag"] = fmt.Sprint(hasher.Sum(nil))
        icon.headers["cache-control"] = "public, max-age=" + fmt.Sprint(maxAge / 1000)
        icon.body = buf

        res.SetHeaders(icon.headers)
        res.WriteBytes(icon.body)
    }
}