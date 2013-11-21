package stackr

import (
	"strconv"
	"time"
)

/*
   Adds the X-Response-Time header displaying the response duration in milliseconds.

   Example:

       stackr.CreateServer().Use(stackr.ResponseTime())

*/
func ResponseTime() func(*Request, *Response, func()) {
	return func(req *Request, res *Response, next func()) {
		start := time.Now().UnixNano()
		res.On("header", func() {
			duration := time.Now().UnixNano() - start
			res.SetHeader("X-Response-Time", strconv.FormatInt(int64(time.Duration(duration)/time.Millisecond), 10)+"ms")
		})
		next()
	}
}
