package stackr

import (
	"encoding/json"
	"fmt"
	"strings"
)

/*
   Error handler, providing res.Error message responses for requests accepting text, html, or json.

   Example:

       stackr.CreateServer().Use(stackr.ErrorHandler())
       stackr.CreateServer().Use(stackr.ErrorHandler("Error Title"))

*/
func ErrorHandler(t ...string) func(*Request, *Response, func()) {

	title := "Forgery"

	if len(t) == 1 {
		title = t[0]
	}

	return func(req *Request, res *Response, next func()) {
		// Call next as we're going to wait for an error.
		next()
		// If there is no error then call next.
		if res.Error == nil {
			next()
			return
		}
		// Otherwise...
		// Reset the status code to a general error.
		if res.StatusCode < 400 {
			res.StatusCode = 500
		}
		// html
		if strings.Index(req.Header.Get("Accept"), "html") > 0 {
			res.SetHeader("Content-Type", "text/html; charset=utf-8")
			res.End("<html><head><meta charset=\"utf-8\"><title>" +
				fmt.Sprint(res.Error) + "</title></head><body><h1>" + title + "</h1><h2><em>" +
				fmt.Sprint(res.StatusCode) + "</em> " +
				fmt.Sprint(res.Error) + "</h2></body></html></h1>")
			return
		}
		// json
		if strings.Index(req.Header.Get("Accept"), "json") > 0 {
			res.SetHeader("Content-Type", "application/json")
			j, _ := json.Marshal(map[string]string{
				"code":  fmt.Sprint(res.StatusCode),
				"error": fmt.Sprint(res.Error),
			})
			res.WriteBytes(j)
			return
		}
		// plain text
		res.SetHeader("Content-Type", "text/plain")
		res.End(fmt.Sprint(res.Error))
	}
}
