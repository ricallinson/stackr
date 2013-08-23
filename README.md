# Stack

Stack is a middleware framework for Go, shipping with over 1 bundled middleware and a poor selection of 3rd-party middleware.

    package main
    import "github.com/ricallinson/stack"

    func main() {
        app := stack.CreateServer()
        app.Use("", stack.Logger(stack.LogOpt{}))
        app.Use("", stack.Favicon(stack.FavOpt{}))
        app.Use("/", func(req *stack.Request, res *stack.Response, next func()) {
            res.End("Hello world\n")
        })
        app.Listen(3000)
    }

## Middleware

* `Logger` request logger with currently __currently no__ custom format support
* `Favicon` efficient favicon server (that doesn't work yet)

## Notes

This project started out as a clone of the Node.js library [Connect](http://www.senchalabs.org/connect/).