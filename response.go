package stackr

import (
	"fmt"
	"net/http"
)

/*
   Response represents the response from an HTTP request.
*/
type Response struct {

	// The standard http.ResponseWriter interface.
	Writer http.ResponseWriter

	// Ture if headers have been sent.
	HeaderSent bool

	// The HTTP status code to be return.
	StatusCode int

	// True if .End() has been called.
	Closed bool

	// Error. Populated by anything that wants to trigger an error.
	Error error

	// Events
	events map[string][]func()
}

/*
   Returns a new HTTP Response.
*/

func createResponse(writer http.ResponseWriter) *Response {

	/*
	   Create a new Response.
	*/

	this := &Response{writer, false, 200, false, nil, map[string][]func(){}}

	/*
	   Return the finished stack.Response.
	*/

	return this
}

/*
   Register a listener function for an event.
*/
func (this *Response) On(event string, fn func()) {
	this.events[event] = append(this.events[event], fn)
}

/*
   Emit an event calling all registered listeners.
*/
func (this *Response) Emit(event string) {
	e, ok := this.events[event]
	if ok {
		for _, fn := range e {
			fn()
		}
	}
}

/*
   Set a map of headers, calls SetHeader() for each item.
*/
func (this *Response) SetHeaders(headers map[string]string) bool {
	for key, value := range headers {
		if this.SetHeader(key, value) == false {
			return false
		}
	}
	return true
}

/*
   Set a single header.

   Note: all keys are converted to lower case.
*/
func (this *Response) SetHeader(key string, value string) bool {

	/*
	   If the headers have been sent nothing can be done so return false.
	*/

	if this.HeaderSent == true {
		return false
	}

	/*
	   http://www.w3.org/Protocols/rfc2616/rfc2616-sec4.html#sec4.2
	   Message headers are case-insensitive.
	*/

	if len(value) > 0 {
		this.Writer.Header().Set(key, value)
	}

	/*
	   The header was set so return true.
	*/

	return true
}

/*
   Remove the named header.
*/
func (this *Response) RemoveHeader(key string) {
	this.Writer.Header().Del(key)
}

/*
   Write any headers set to the client.
*/

func (this *Response) writeHeaders() {

	/*
	   Fire an event.
	*/

	this.Emit("header")

	/*
	   Set the HeaderSent flag to true.
	*/

	this.HeaderSent = true

	/*
	   Write the headers with the current StatusCode.
	*/

	this.Writer.WriteHeader(this.StatusCode)
}

/*
   Write bytes to the client.
*/
func (this *Response) WriteBytes(data []byte) bool {

	/*
	   If headers have not been sent call writeHeaders().
	*/

	if this.HeaderSent == false {
		this.writeHeaders()
	}

	/*
	   Try and write the byte array to the client.
	*/

	writen, err := this.Writer.Write(data)

	/*
	   If there was an error return false.
	*/

	if err != nil {
		return false
	}

	/*
	   Return true if the number of bytes written matches the data length.
	*/

	return writen == len(data)
}

/*
   Write data to the client.
*/
func (this *Response) Write(data string) bool {

	/*
	   If headers have not been sent call writeHeaders().
	*/

	if this.HeaderSent == false {
		this.writeHeaders()
	}

	/*
		If the string was empty just return.
	*/

	if len(data) == 0 {
		return true
	}

	/*
	   Try and write the string to the client.
	*/

	writen, err := fmt.Fprint(this.Writer, data)

	/*
	   If there was an error return false.
	*/

	if err != nil {
		return false
	}

	/*
	   Return true if the number of bytes written matches the data length.
	*/

	return writen == len(data)
}

/*
   Close the connection to the client.
*/
func (this *Response) End(data string) bool {

	status := true

	/*
	   Write the data to the client.
	*/

	status = this.Write(data)

	/*
	   Set the "Closed" flag to true.
	*/

	this.Closed = true

	/*
	   Return the status of the write operation.
	*/

	return status
}
