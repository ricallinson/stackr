package connect

import(
    "fmt"
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
    Set a header.
*/

func (this *Response) SetHeader(key string, value string) (bool) {
    if this.HeaderSent == true {
        return false
    }
    this.Writer.Header().Set(key, value)
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