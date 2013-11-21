package stackr

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
)

/*
   The options for the favicon middleware.
*/
type faviconOpt struct {
	Path   string
	MaxAge int
}

/*
	By default serves the stackr favicon, or the favicon located by the given `path`.

	Options:

		* `MaxAge` cache-control max-age directive, defaulting to 1 day

	Examples:

	Serve default favicon:

		stackr.CreateServer().Use(stackr.Favicon())
		stackr.CreateServer().Use(stackr.Favicon(map[string]string{"maxage": "1000"}))

	Serve favicon before logging for brevity:

		app := stackr.CreateServer()
		app.Use(stackr.Favicon())
		app.Use(stackr.Logger())

	Serve custom favicon:

		stackr.CreateServer().Use(stackr.Favicon(map[string]string{"path": "./public/favicon.ico"}))
*/
func Favicon(o ...map[string]string) func(*Request, *Response, func()) {

	/*
	   If we got options use them.
	*/

	opt := faviconOpt{}

	if len(o) == 1 {
		val := o[0]
		opt.Path = val["path"]
		opt.MaxAge, _ = strconv.Atoi(val["maxage"])
	}

	/*
	   Create an Icon.
	*/

	type Icon struct {
		headers map[string]string
		body    []byte
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
	   The Handler function returned to Use().
	*/

	return func(req *Request, res *Response, next func()) {

		/*
		   If this is not a fav icon return fast
		*/

		if req.OriginalUrl != "/favicon.ico" {
			next()
			return
		}

		/*
		   If we have the icon cached, serve it.
		*/

		if len(icon.body) > 0 {
			res.SetHeaders(icon.headers)
			res.WriteBytes(icon.body)
			res.End("")
			next()
			return
		}

		/*
		   Otherwise read the icon into cache.
		*/

		buf, err := ioutil.ReadFile(path)
		if err != nil {
			fmt.Println(err)
			next()
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
		icon.headers["cache-control"] = "public, max-age=" + fmt.Sprint(maxAge/1000)
		icon.body = buf

		/*
		   Serve the icon.
		*/

		res.SetHeaders(icon.headers)
		res.WriteBytes(icon.body)
		res.End("")
	}
}
