package stackr

import (
	"github.com/ricallinson/httputils"
	"strings"
)

func supportsMethod(method string) bool {
	method = strings.ToLower(method)
	for _, m := range httphelp.Methods {
		if method == m {
			return true
		}
	}
	return false
}

/*
	Provides faux HTTP method support.

	Pass an optional key to use when checking for a method override, otherwise defaults to _method.
	The original method is available via req.Map["OriginalMethod"].

	Example:

		stackr.CreateServer().Use(stackr.MethodOverride())

*/
func MethodOverride() func(*Request, *Response, func()) {
	return func(req *Request, res *Response, next func()) {
		// If there's a Method override header.
		if method := req.Header.Get("X-HTTP-Method-Override"); method != "" {
			// Copy the original method into the req.Map
			if _, ok := req.Map["OriginalMethod"]; ok == false {
				req.Map["OriginalMethod"] = req.Method
			}
			// See if the method is supported.
			if supportsMethod(method) {
				// If it is then swap it in.
				req.Method = strings.ToUpper(method)
			}
		}
		// Call the next middleware.
		next()
	}
}
