package connect

import(
    "fmt"
    "strings"
    "net/http"
)

type Response struct {
    Writer http.ResponseWriter
    HeaderSent bool
    StatusCode int
    Closed bool
}

func CreateResponse(writer http.ResponseWriter) (*Response) {
    this := new(Response)
    this.Writer = writer
    this.StatusCode = 200
    return this
}

/*
    Set map of headers
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
    Set a header.
*/

func (this *Response) SetHeader(key string, value string) (bool) {
    if this.HeaderSent == true {
        return false
    }
    /*
        http://www.w3.org/Protocols/rfc2616/rfc2616-sec4.html#sec4.2
        Message headers are case-insensitive so they are forced to lower case.
    */
    this.Writer.Header().Set(strings.ToLower(key), value)
    return true
}

/*
    Write headers.
*/

func (this *Response) writeHeaders() {
    this.HeaderSent = true
    this.Writer.WriteHeader(this.StatusCode);
}

/*
    Write bytes to the client.
*/

func (this *Response) WriteBytes(data []byte) (bool) {
    if this.HeaderSent == false {
        this.writeHeaders()
    }
    writen, err := this.Writer.Write(data)
    if err != nil {
        return false
    }
    return writen == len(data)
}

/*
    Write data to the client.
*/

func (this *Response) Write(data string) (bool) {
    if this.HeaderSent == false {
        this.writeHeaders()
    }
    writen, err := fmt.Fprint(this.Writer, data)
    if err != nil {
        return false
    }
    return writen == len(data)
}

/*
    Closes the connection to the client.
*/

func (this *Response) End(data string) (bool) {
    status := this.Write(data)
    this.Closed = true
    return status
}