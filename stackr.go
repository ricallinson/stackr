/*
    Stackr is an extensible HTTP server framework for Go.

        package main

        import "github.com/ricallinson/stackr"

        func main() {
            app := stackr.CreateServer()
            app.Use("/", stackr.Logger())
            app.Use("/", stackr.Static())
            app.Use("/", func(req *stackr.Request, res *stackr.Response, next func()) {
                res.End("hello world\n")
            })
            app.Listen(3000)
        }
*/
package stackr

import(
    "fmt"
    "log"
    "strings"
    "net/http"
)

/*
    A Stackr Server.
*/
type Server struct {
    stack []middleware
}

/*
    A middleware.
*/

type middleware struct {
    Route string
    Handle func(*Request, *Response, func())
}

/*
    Create a new stackr server.
*/
func CreateServer() (*Server) {
    return new(Server)
}

/*
    Utilize the given middleware `handle` to the given `route`,
    defaulting to _/_. This "route" is the mount-point for the
    middleware, when given a value other than _/_ the middleware
    is only effective when that segment is present in the request's
    pathname.

    For example if we were to mount a function at _/admin_, it would
    be invoked on _/admin_, and _/admin/settings_, however it would
    not be invoked for _/_, or _/posts_.

    Examples:

        var app = stackr.CreateServer();
        app.Use("/", stackr.Favicon(stackr.FavOpt{}))
        app.Use("/", stackr.Logger(stackr.LogOpt{}))
        app.Use("/", stackr.Static(stackr.StaticOpt{}))

    If we wanted to prefix static files with _/public_, we could
    "mount" the `Static()` middleware:

        app.Use("/public", stackr.Static(stackr.StaticOpt{Root: "./static_files"}))

    This api is chainable, so the following is valid:

        stackr.CreateServer().Use(stackr.Favicon(stackr.FavOpt{})).Listen(3000);
*/ 
func (this *Server) Use(route string, handle func(*Request, *Response, func())) (*Server) {

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
        Route: strings.ToLower(route),
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
func (this *Server) handle(req *Request, res *Response, index int) {

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
        layer = this.stack[index]; // yes
        index++ // increment the index by 1
    }

    /*
        If there are no more layers and no headers have been sent return a 404.
    */

    if layer.Handle == nil && res.HeaderSent == false {
        res.StatusCode = 404
        res.SetHeader("Content-Type", "text/plain")
        if req.Raw.Method == "HEAD" {
            res.End("")
            return
        }
        res.End("Cannot " + fmt.Sprint(req.Raw.Method) + " " + fmt.Sprint(req.OriginalUrl));
        return
    }

    /*
        If there are no more layers and headers were sent then we are done so just return.
    */

    if layer.Handle == nil {
        return
    }

    /*
        Otherwise call the layer handler.
    */

    if strings.Contains(strings.ToLower(req.OriginalUrl), layer.Route) {

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

            this.handle(req, res, index)
        })
    }

    /*
        Call this function again, passing the index value.
    */

    this.handle(req, res, index)
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
        Start the server.
    */

    log.Fatal(http.ListenAndServe(address, createHttpHandler(this)))
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
        Start the server.
    */

    log.Fatal(http.ListenAndServeTLS(address, certFile, keyFile, createHttpHandler(this)))
}