package stackr

import(
    "fmt"
    "time"
    "strings"
    "strconv"
)

/*
    Options for the logger middleware. _Note: future options commented out._
*/
type OptLog struct {
    Format string
    Writer func(...interface {}) (int, error)
    // Buffer int
    Immediate bool
    startTime int64
}

/*
    Logger output format options.
*/

var loggerFormatOptions map[string]string = map[string]string{
    "default": ":remote-addr - - [:date] \":method :url :http-version\" :status :res[content-length] \":referrer\" \":user-agent\"",
    "short": ":remote-addr - :method :url :http-version :status :res[content-length] - :response-time ms",
    "tiny": ":method :url :status :res[content-length] - :response-time ms",
    "dev": "",
}

/*
    Logger format functions.
*/

var loggerFormatFunctions map[string]func(*OptLog, *Request, *Response)string = map[string]func(*OptLog, *Request, *Response)string{

    /*
        Response header content-length.
    */

    ":res[content-length]": func(opt *OptLog, req *Request, res *Response) string {
        length := res.Writer.Header().Get("content-length")
        if len(length) == 0 {
            length = "0"
        }
        return length
    },

    /*
        HTTP version.
    */

    ":http-version": func(opt *OptLog, req *Request, res *Response) string {
        return req.Raw.Proto
    },

    /*
        Response time in milliseconds.
    */

    ":response-time": func(opt *OptLog, req *Request, res *Response) string {
        return fmt.Sprint((time.Now().UnixNano() - opt.startTime) / 1000000)
    },

    /*
        Remote address.
    */

    ":remote-addr": func(opt *OptLog, req *Request, res *Response) string {
        return req.Raw.RemoteAddr
    },

    /*
        UTC date.
    */

    ":date": func(opt *OptLog, req *Request, res *Response) string {
        return time.Now().Format(time.RFC3339)
    },

    /*
        Request method.
    */

    ":method": func(opt *OptLog, req *Request, res *Response) string {
        return req.Raw.Method
    },

    /*
        Request url.
    */

    ":url": func(opt *OptLog, req *Request, res *Response) string {
        return req.OriginalUrl
    },

    /*
        Normalized referrer.
    */

    ":referrer": func(opt *OptLog, req *Request, res *Response) string {
        return req.Raw.Referer()
    },

    /*
        UA string.
    */

    ":user-agent": func(opt *OptLog, req *Request, res *Response) string {
        return req.Raw.UserAgent()
    },

    /*
        Response status code.
    */

    ":status": func(opt *OptLog, req *Request, res *Response) string {
        return fmt.Sprint(res.StatusCode)
    },
}

/*
    Logger:

    Log requests with the given `options` or a `format` string.

    __Options:__

        * `format`  Format string, see below for tokens
        * `writer`  Output writer, defaults to _fmt.Println_
        * `buffer`  (not implemented yet) Buffer duration, defaults to 1000ms when _true_
        * `immediate`  Write log line on request instead of response (for response times)

    Tokens:

        * (not implemented) `:req[header]` ex: `:req[Accept]`
        * `:res[Content-Length]`
        * `:http-version`
        * `:response-time`
        * `:remote-addr`
        * `:date`
        * `:method`
        * `:url`
        * `:referrer`
        * `:user-agent`
        * `:status`

    Formats:

    Pre-defined formats that ship with connect:

        * `default` ':remote-addr - - [:date] ":method :url HTTP/:http-version" :status :res[content-length] ":referrer" ":user-agent"'
        * `short` ':remote-addr - :method :url HTTP/:http-version :status :res[content-length] - :response-time ms'
        * `tiny`  ':method :url :status :res[content-length] - :response-time ms'
        * `dev` concise output colored by response status for development use

    Examples:

        app.Use("/", stackr.Logger()) // default
        app.Use("/", stackr.Logger(stackr.OptLog{Format: "short"}))
        app.Use("/", stackr.Logger(stackr.OptLog{Format: "tiny"}))
        app.Use("/", stackr.Logger(stackr.OptLog{Immediate: true, Format: "dev"})
        app.Use("/", stackr.Logger(stackr.OptLog{Format: ":method :url - :referrer"})
        app.Use("/", stackr.Logger(stackr.OptLog{Format: ":req[content-type] -> :res[content-type]"})
*/
func Logger(o ...OptLog) (func(req *Request, res *Response, next func())) {

    /*
        If we got an OptLog use it.
    */

    var opt OptLog

    if len(o) == 1 {
        opt = o[0]
    } else {
        opt = OptLog{}
    }

    /*
        If we were not given a format use "default".
    */

    if opt.Format == "" {
        opt.Format = "default"
    }

    /*
        Set the default stream.
    */

    writer := fmt.Println

    /*
        If we were given a different stream use that.
    */

    if opt.Writer != nil {
        writer = opt.Writer
    }

    /*
        Output on request instead of response.
    */

    immediate := opt.Immediate

    /*
        Return the handler function.
    */

    return func(req *Request, res *Response, next func()) {

        /*
            Grab the start time.
        */

        opt.startTime = time.Now().UnixNano()

        /*
            If we are to log at the end of the request call next() now.
        */

        if immediate == false {

            /*
                Once all the other middleware has run, execution
                will come back to this point and continue.
            */

            next()
        }

        /*
            Format the log string requested and write it to the stream function.
        */

        writer(loggerFormat(opt, req, res, opt.Format))
    }
}

/*
    Format the log with the given format string.
*/

func loggerFormat(opt OptLog, req *Request, res *Response, format string) (string) {

    /*
        If format is "dev" we are done.
    */

    if format == "dev" {
        return loggerFormatDev(opt, req, res)
    }

    /*
        See if "format" is a key in loggerFormats.
    */

    log := loggerFormatOptions[format]

    /*
        If there is no format matched then use the format string as the log template.
    */

    if len(log) == 0 {
        log = format
    }

    /*
        Replace tokens in the log template string.
    */

    for match := range loggerFormatFunctions {
        log = strings.Replace(log, match, loggerFormatFunctions[match](&opt, req, res), -1)
    }

    /*
        Return the final string.
    */

    return log
}

/*
    Format the log for "dev" mode.
*/

func loggerFormatDev(opt OptLog, req *Request, res *Response) (string) {

    /*
        Get the length of the data sent.
    */

    length, _ := strconv.Atoi(loggerFormatFunctions[":res[content-length]"](&opt, req, res))

    /*
        The length as a string.
    */

    strLen := ""

    if length > 0 {
        strLen = " - " + fmt.Sprint(length);
    }

    /*
        Set the default color for the log.
    */

    color := 32

    /*
        Get the status code for the request.
    */

    status := res.StatusCode

    /*
        Pick a color for the log.
    */

    switch {
    case status >= 500:
        color = 31
    case status >= 400:
        color = 33
    case status >= 300:
        color = 36
    }

    /*
        Build the log line.
    */

    log := "\x1b[90m" + loggerFormatFunctions[":method"](&opt, req, res)
    log += " " + loggerFormatFunctions[":url"](&opt, req, res) + " "
    log += "\x1b[" + fmt.Sprint(color) + "m" + fmt.Sprint(status)
    log += " \x1b[90m"
    log += loggerFormatFunctions[":response-time"](&opt, req, res)
    log += "ms" + strLen
    log += "\x1b[0m"

    /*
        Return the log string.
    */

    return log
}