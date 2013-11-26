/*
   Stackr is an extensible HTTP server framework for Go.

	   package main

	   import "github.com/ricallinson/stackr"

	   func main() {
		   app := stackr.CreateServer()
		   app.Use(stackr.Logger())
		   app.Use(stackr.Static())
		   app.Use("/", func(req *stackr.Request, res *stackr.Response, next func()) {
			   res.End("hello world\n")
		   })
		   app.Listen(3000)
	   }
*/
package stackr

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

/*
   A Stackr Server.
*/
type Server struct {
	Env   string
	stack []middleware
}

/*
   A middleware.
*/

type middleware struct {
	Route  string
	Handle func(*Request, *Response, func())
}

/*
   Create a new stackr server.
*/
func CreateServer() *Server {

	this := &Server{}

	this.Env = os.Getenv("GO_ENV")

	if this.Env == "" {
		this.Env = "development"
	}

	return this
}

/*
   Utilize the given middleware `Handle` to the given `route`,
   defaulting to _/_. This "route" is the mount-point for the
   middleware, when given a value other than _/_ the middleware
   is only effective when that segment is present in the request's
   pathname.

   For example if we were to mount a function at _/admin_, it would
   be invoked on _/admin_, and _/admin/settings_, however it would
   not be invoked for _/_, or _/posts_.

   Examples:

	   var app = stackr.CreateServer();
	   app.Use(stackr.Favicon())
	   app.Use(stackr.Logger())
	   app.Use("/public", stackr.Static())

   If we wanted to prefix static files with _/public_, we could
   "mount" the `Static()` middleware:

	   app.Use("/public", stackr.Static(stackr.OptStatic{Root: "./static_files"}))

   This api is chainable, so the following is valid:

	   stackr.CreateServer().Use(stackr.Favicon()).Listen(3000);
*/
func (this *Server) Use(in ...interface{}) *Server {

	var route string
	var handle func(*Request, *Response, func())

	for _, i := range in {
		switch i.(type) {
		case string:
			route = i.(string)
		case func(*Request, *Response, func()):
			handle = i.(func(*Request, *Response, func()))
		default:
			panic("stackr: Go home handler, you're drunk!")
		}
	}

	/*
	   If the route is empty make it "/".
	*/

	if len(route) == 0 {
		route = "/"
	}

	/*
	   Strip trailing slash
	*/

	if size := len(route); size > 1 && route[size-1] == '/' {
		route = route[:size-1]
	}

	/*
	   Add the middleware to the stack.
	*/

	this.stack = append(this.stack, middleware{
		Route:  strings.ToLower(route),
		Handle: handle,
	})

	/*
	   Return the Server so calls can be chained.
	*/

	return this
}

/*
   Handle server requests, punting them down
   the middleware stack.

   Note: this is a recursive function.
*/
func (this *Server) Handle(req *Request, res *Response, index int) {

	/*
		For each call to Handle we want to catch anything that panics unless in development mode.
	*/

	defer func() {
		if this.Env != "development" {
			err := recover()
			if err != nil {
				res.Error = errors.New(fmt.Sprint(err))
			}
		}
	}()

	/*
	   If the response has been closed return.
	*/

	if res.Closed == true {
		return
	}

	/*
	   Create a var for the middleware.
	*/

	var layer middleware

	/*
	   Do we have another layer to use?
	*/

	if index >= len(this.stack) {
		layer = middleware{} // no
	} else {
		layer = this.stack[index] // yes
		index++                   // increment the index by 1
	}

	/*
	   If there are no more layers and no headers have been sent return a 404.
	*/

	if layer.Handle == nil && res.HeaderSent == false {
		res.StatusCode = 404
		res.SetHeader("Content-Type", "text/plain")
		if req.Method == "HEAD" {
			res.End("")
			return
		}
		res.End("Cannot " + fmt.Sprint(req.Method) + " " + fmt.Sprint(req.OriginalUrl))
		return
	}

	/*
	   If there are no more layers and headers were sent then we are done so just return.
	*/

	if layer.Handle == nil {
		return
	}

	/*
	   Otherwise call the layer Handler.
	*/

	if strings.Contains(req.OriginalUrl, layer.Route) {

		/*
		   Set the value of Url to the portion after the matched layer.Route
		*/

		req.Url = strings.TrimPrefix(req.OriginalUrl, layer.Route)

		/*
		   Set the matched portion of the Url.
		*/

		req.MatchedUrl = layer.Route

		/*
		   Call the middleware function.
		*/

		layer.Handle(req, res, func() {

			/*
			   The value of next is a function that calls this function again, passing the index value.
			*/

			this.Handle(req, res, index)
		})

	} else {

		/*
		   Call this function, passing the index value.
		*/

		this.Handle(req, res, index)
	}
}

/*
   ServeHTTP calls .Handle(req, res).
*/
func (this *Server) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	/*
	   Pass the res and req into there repective create functions.
	   The results of these are then passed to stack.server.Handle().
	*/

	this.Handle(createRequest(req), createResponse(res), 0)
}

/*
   Listen for connections on HTTP.
*/
func (this *Server) Listen(port int) {

	/*
	   Set the address to run on.
	*/

	address := ":" + fmt.Sprint(port)

	/*
	   Start the server by mapping all URLs into this Server.
	*/

	log.Fatal(http.ListenAndServe(address, this))
}

/*
   Listen for connections on HTTPS.
*/
func (this *Server) ListenTLS(port int, certFile string, keyFile string) {

	/*
	   Set the address to run on.
	*/

	address := ":" + fmt.Sprint(port)

	/*
	   Start the server by mapping all URLs into this Server.
	*/

	log.Fatal(http.ListenAndServeTLS(address, certFile, keyFile, this))
}
