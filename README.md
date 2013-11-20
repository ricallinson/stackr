# Stackr

[![Build Status](https://secure.travis-ci.org/ricallinson/stackr.png?branch=master)](http://travis-ci.org/ricallinson/stackr)

__WARNING: UNSTABLE API__

Stackr is an extensible HTTP server framework for Go, shipping with over 5 bundled middleware and a poor selection of 3rd-party middleware.

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

## Middleware

* `ErrorHandler` flexible error handler
* `Favicon` efficient favicon server
* `Logger` request logger with custom format support
* `ResponseTime` calculates response-time and exposes via X-Response-Time
* `MethodOverride` faux HTTP method support
* `Static` static file server currently based on http.FileServer

## Testing

The following should all be executed from the `stackr` directory $GOPATH/src/github.com/ricallinson/stackr/.

#### Install

    go get github.com/ricallinson/simplebdd

#### Run

    go test

### Code Coverage

#### Install

    go get github.com/axw/gocov/gocov
    go get -u github.com/matm/gocov-html

#### Generate

    gocov test | gocov-html > ./reports/coverage.html

## Notes

This project started out as a clone of the superb Node.js library [Connect](http://www.senchalabs.org/connect/).
