package stackr

import(
    "fmt"
    "strings"
    "net/http"
)

/*
    A HTTP Response.
*/
type Response struct {
    Writer http.ResponseWriter
    HeaderSent bool
    StatusCode int
    Closed bool
}

/*
    Returns a new HTTP Response.
*/

func createResponse(writer http.ResponseWriter) (*Response) {

    /*
        Create a new Response.
    */

    this := new(Response)

    /*
        Set the source http.ResponseWriter so it can be accessed later.
    */

    this.Writer = writer

    /*
        Set a default status code.
    */

    this.StatusCode = 200

    /*
        Return the finished stack.Response.
    */

    return this
}

/*
    Set a map of headers, calls SetHeader() for each item.
*/
func (this *Response) SetHeaders(headers map[string]string) (bool) {
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
func (this *Response) SetHeader(key string, value string) (bool) {

    /*
        If the headers have been sent nothing can be done so return false.
    */

    if this.HeaderSent == true {
        return false
    }

    /*
        http://www.w3.org/Protocols/rfc2616/rfc2616-sec4.html#sec4.2
        Message headers are case-insensitive so they are forced to lower case.
    */

    this.Writer.Header().Set(strings.ToLower(key), value)

    /*
        The header was set so return true.
    */

    return true
}

/*
    Write any headers set to the client.
*/

func (this *Response) writeHeaders() {

    /*
        Set the HeaderSent flag to true.
    */

    this.HeaderSent = true

    /*
        Write the headers with the current StatusCode.
    */

    this.Writer.WriteHeader(this.StatusCode);
}

/*
    Write bytes to the client.
*/
func (this *Response) WriteBytes(data []byte) (bool) {

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
func (this *Response) Write(data string) (bool) {

    /*
        If headers have not been sent call writeHeaders().
    */

    if this.HeaderSent == false {
        this.writeHeaders()
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
func (this *Response) End(data string) (bool) {

    /*
        Write the data to the client.
    */

    status := this.Write(data)

    /*
        Set the "Closed" flag to true.
    */

    this.Closed = true

    /*
        Return the status of the write operation.
    */

    return status
}