package connect

type middleware struct {
    Route string
    Handle func(*Request, *Response, func())
}