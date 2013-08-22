package connect

import(
    "fmt"
    "time"
)

type LoggerOpt struct {
    format string
    stream string
    buffer int
    immediate bool
    startTime int64
}

func loggerFormat(options LoggerOpt, req *Request, res *Response) (string) {

    // Get the time taken in milliseconds
    totalTime := (time.Now().UnixNano() - options.startTime) / 1000000

    return fmt.Sprint("[" + fmt.Sprint(res.StatusCode) + "] " + req.Raw.Method + " " + req.Url + " " + fmt.Sprint(totalTime) + "ms")
}

func Logger(options LoggerOpt) (func(req *Request, res *Response, next func())) {

    // Output on request instead of response.
    immediate := options.immediate

    // Return the handler function.
    return func(req *Request, res *Response, next func()) {

        // Grab the start time.
        options.startTime = time.Now().UnixNano()

        // Decide if we should log immediately or at the end of the request.
        if immediate {
            line := loggerFormat(options, req, res)
            fmt.Println(line)
        } else {
            next()
            line := loggerFormat(options, req, res)
            fmt.Println(line)
        }
    }
}