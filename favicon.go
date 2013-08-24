package stackr

import(
    "io"
    "fmt"
    "io/ioutil"
    "crypto/md5"
)

/*
    The options for the favicon middleware.
*/
type FavOpt struct {
    Path string
    MaxAge int
}

/*
    Favicon:

    By default serves the stackr favicon, or the favicon located by the given `path`.

    Options:

    * `maxAge`  cache-control max-age directive, defaulting to 1 day

    Examples:

        Serve default favicon:

        stackr.CreateServer().Use(stackr.Favicon(stackr.FavOpt{}))

    Serve favicon before logging for brevity:

        app := stackr.CreateServer()
        app.Use(stackr.Favicon(stackr.FavOpt{}))
        app.Use(stackr.Logger(stackr.LogOpt{}))

    Serve custom favicon:
    
        stack.CreateServer().Use(stack.Favicon(stack.FavOpt{path "./public/favicon.ico"}))
 */
func Favicon(opt FavOpt) (func(req *Request, res *Response, next func())) {

    /*
        Create an Icon.
    */

    type Icon struct {
        headers map[string]string
        body []byte
    }

    /*
        Create a new map.
    */

    icon := Icon{
        headers: make(map[string]string),
    }

    /*
        Set the default maxAge.
    */

    maxAge := 86400000

    /*
        If we were given a max age use it.
    */

    if opt.MaxAge > 0 {
        maxAge = opt.MaxAge
    }

    /*
        Set the default path.
    */

    path := "./favicon.ico"

    /*
        If we were given a path use it.
    */

    if len(opt.Path) > 0 {
        path = opt.Path
    }

    /*
        The handler function returned to Use().
    */

    return func(req *Request, res *Response, next func()) {

        /*
            If this is not a fav icon return fast
        */

        if req.OriginalUrl != "/favicon.ico" {
            return
        }

        /*
            If we have the icon cached, serve it.
        */

        if len(icon.body) > 0 {
            res.SetHeaders(icon.headers)
            res.WriteBytes(icon.body)
            res.End("");
            return
        }

        /*
            Otherwise read the icon into cache.
        */

        buf, err := ioutil.ReadFile(path)
        if err != nil {
            fmt.Println(err)
            return
        }

        /*
            Generate an MD5 of the icon to be used in the etag.
        */

        hasher := md5.New()
        io.WriteString(hasher, fmt.Sprint(buf))

        /*
            Create headers for the icon.
        */

        icon.headers["content-type"] = "image/x-icon"
        icon.headers["content-length"] = fmt.Sprint(len(buf))
        icon.headers["etag"] = fmt.Sprint(hasher.Sum(nil))
        icon.headers["cache-control"] = "public, max-age=" + fmt.Sprint(maxAge / 1000)
        icon.body = buf

        /*
            Serve the icon.
        */

        res.SetHeaders(icon.headers)
        res.WriteBytes(icon.body)
        res.End("");
    }
}