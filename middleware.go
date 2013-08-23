package stack

type middleware struct {
    Route string
    Handle func(*Request, *Response, func())
}