package stack

import(
    "fmt"
    "time"
    "strconv"
)

/*
    Options for the logger middleware.
*/

type LogOpt struct {
    // format string
    // stream string
    // buffer int
    immediate bool
    startTime int64
}

/*
    Format the log for "dev" mode.
*/

func loggerFormatDev(opt LogOpt, req *Request, res *Response) (string) {

    // Get the time taken in milliseconds.
    totalTime := (time.Now().UnixNano() - opt.startTime) / 1000000

    // Get the status code for the request.
    status := res.StatusCode

    // Get the length of the data sent.
    length, _ := strconv.Atoi(res.Writer.Header().Get("content-length"))

    // The length as a string.
    strLen := ""

    if length > 0 {
        strLen = " - " + fmt.Sprint(length);
    }

    // Set the default color for the log.
    color := 32

    // Pick a color for the log.
    switch {
    case status >= 500:
        color = 31
    case status >= 400:
        color = 33
    case status >= 300:
        color = 36
    }

    // Build the log line.
    log := "\x1b[90m" + req.Raw.Method
    log += " " + req.OriginalUrl + " "
    log += "\x1b[" + fmt.Sprint(color) + "m" + fmt.Sprint(status)
    log += " \x1b[90m"
    log += fmt.Sprint(totalTime)
    log += "ms" + strLen
    log += "\x1b[0m"

    return log
}

func Logger(opt LogOpt) (func(req *Request, res *Response, next func())) {

    // Output on request instead of response.
    immediate := opt.immediate

    // Return the handler function.
    return func(req *Request, res *Response, next func()) {

        // Grab the start time.
        opt.startTime = time.Now().UnixNano()

        // Decide if we should log immediately or at the end of the request.
        if immediate {
            line := loggerFormatDev(opt, req, res)
            fmt.Println(line)
        } else {
            next()
            line := loggerFormatDev(opt, req, res)
            fmt.Println(line)
        }
    }
}