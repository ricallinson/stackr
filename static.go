package stack

import(
    "os"
    "fmt"
    "net/http"
)

/*
    Options for the static middleware.

    Note: future options commented out.
*/

type StaticOpt struct {
    Root string
    // MaxAge int64
    // Hidden bool
    // Redirect bool
    // Index bool
}

/*
    Static:

    Static file server with the given `root` path.

    Examples:

        var oneDay = 86400000;

        stack.CreateServer().Use("/", stack.Static(stack.StaticOpt{Root: "./public"}))

    Options (not implemented yet):

        - `maxAge`     Browser cache maxAge in milliseconds. defaults to 0
        - `hidden`     Allow transfer of hidden files. defaults to false
        - `redirect`   Redirect to trailing "/" when the pathname is a dir. defaults to true
        - `index`      Default file name, defaults to 'index.html'
*/

func Static(opt StaticOpt) (func(req *Request, res *Response, next func())) {

    /*
        The default loction of static files.
    */

    root := "./static"

    /*
        If we were given a root use it.
    */

    if len(opt.Root) > 0 {
        root = opt.Root
    }

    /*
        Create a http.FileServer to server the files.
    */

    fileServer := http.FileServer(http.Dir(root))

    /*
        Return the handle function.
    */

    return func(req *Request, res *Response, next func()) {

        /*
            Because http.FileServer serves directories and it's own 404 we 
            want to see if the file is really there before we hand of to it.
            To do that we see if the file exists. If it doesn't, then we return quickly.

            Question: Is this not really expensive?
            Answer: It's not ideal. Writing a custom static server is on the todo list.
        */

        if _, err := os.Stat(root + req.Url); os.IsNotExist(err) {
            return
        }

        /*
            If we have to serve a file strip the matched Url and call ServeHTTP() on the fileServer.
        */

        http.StripPrefix(req.MatchedUrl, fileServer).ServeHTTP(res.Writer, req.Raw)

        /*
            Now call End() to make sure we don't process any more middleware.
        */

        res.End("")
    }
}