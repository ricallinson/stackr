# Stackr

[![Build Status](https://secure.travis-ci.org/ricallinson/stackr.png?branch=master)](http://travis-ci.org/ricallinson/stackr)

Stackr is an extensible HTTP server framework for Go, shipping with over 2 bundled middleware and a poor selection of 3rd-party middleware.

    package main

    import "github.com/ricallinson/stackr"

    func main() {
        app := stackr.CreateServer()
        app.Use("/", stackr.Logger(stackr.LogOpt{}))
        app.Use("/", stackr.Static(stackr.StaticOpt{}))
        app.Use("/", func(req *stackr.Request, res *stackr.Response, next func()) {
            res.End("hello world\n")
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

#### Generate

    gocov test | gocov-html > ./reports/coverage.html

## Notes

This project started out as a clone of the Node.js library [Connect](http://www.senchalabs.org/connect/).