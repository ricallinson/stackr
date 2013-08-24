# Stack

[![Build Status](https://secure.travis-ci.org/ricallinson/stack.png?branch=master)](http://travis-ci.org/ricallinson/stack)

Stack is a middleware framework for Go, shipping with over 2 bundled middleware and a poor selection of 3rd-party middleware.

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

* `Logger` request logger with __no__ custom format support
* `Favicon` efficient favicon server
* `Static` static file server currently based on http.FileServer

## Testing

From the stack directory.

    go test

### Code Coverage

#### Install

    go get github.com/axw/gocov/gocov
    go get -u github.com/matm/gocov-html

#### Run

    gocov test | gocov-html > ./reports/coverage.html

## Notes

This project started out as a clone of the Node.js library [Connect](http://www.senchalabs.org/connect/).