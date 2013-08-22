package connect

import(
    "fmt"
    "log"
    "strings"
    "net/http"
)

type ConnectServer struct {
    stack []middleware
}

/*
    Create a new connect server.
*/

func CreateServer() (*ConnectServer) {
    return new(ConnectServer)
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

        var app = connect.CreateServer();
        app.use(connect.favicon.Load());
        app.use(connect.logger.Load());
        app.use(connect.static.Load("./public"));

    If we wanted to prefix static files with _/public_, we could
    "mount" the `static()` middleware:

        app.use("/public", connect.static.Load(__dirname + "/public"));

    This api is chainable, so the following is valid:

        connect.CreateServer().use(connect.favicon.Load()).listen(3000);
*/ 

func (this *ConnectServer) Use(route string, handle func(*Request, *Response, func())) (*ConnectServer) {

    // If the route is empty make it "/"
    if len(route) == 0 {
        route = "/"
    }

    // strip trailing slash
    if size := len(route); size > 1 && route[size-1] == '/' {
        route = route[:size-1]
    }

    this.stack = append(this.stack, middleware{
        Route: strings.ToLower(route),
        Handle: handle,
    })

    return this
}

/*
    Handle server requests, punting them down
    the middleware stack.   
*/

func (this *ConnectServer) handle(req *Request, res *Response, index int) {

    var layer middleware

    // If the response has been closed return.
    if res.Closed == true {
        return
    }

    // Do we have another layer to use?
    if index >= len(this.stack) {
        layer = middleware{}
    } else {
        layer = this.stack[index];
        index++
    }

    // If there are no more layers and no headers have been sent return a 404.
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

    // If there are no more layers and headers were sent then we are done.
    if layer.Handle == nil {
        return
    }

    // Otherwise call the layer handler
    if strings.Contains(strings.ToLower(req.Url), layer.Route) {
        req.UrlMatched = layer.Route
        layer.Handle(req, res, func() {
            this.handle(req, res, index)
        })
    }

    this.handle(req, res, index)
}

/*
    Listen for connections on HTTP.
*/

func (this *ConnectServer) Listen(port int) {

    address := ":" + fmt.Sprint(port)

    log.Fatal(http.ListenAndServe(address, createHttpHandler(this)))
}

/*
    Listen for connections on HTTPS.
*/

func (this *ConnectServer) ListenTLS(port int, certFile string, keyFile string) {

    address := ":" + fmt.Sprint(port)

    log.Fatal(http.ListenAndServeTLS(address, certFile, keyFile, createHttpHandler(this)))
}