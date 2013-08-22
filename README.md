# Connect

Connect is a middleware framework for Go, shipping with over 1 bundled middleware and a poor selection of 3rd-party middleware.

    package main
    import "github.com/ricallinson/connect"

    func main() {
        app := connect.CreateServer()
        app.Use("", connect.Logger(connect.LogOpt{}))
        app.Use("", connect.Favicon(connect.FavOpt{}))
        app.Use("/", func(req *connect.Request, res *connect.Response, next func()) {
            res.End("Hello world\n")
        })
        app.Listen(3000)
    }

## Middleware

* `Logger` request logger with currently __currently no__ custom format support
* `Favicon` efficient favicon server (that doesn't work yet)

## Notes

In case you don't know this is clone of the Node.js library by the same name http://www.senchalabs.org/connect/.