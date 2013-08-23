package stack

import(
    "net/http"
)

/*
    A HTTP Handler for stack.
*/

type Handler struct {
    server *Server
}

/*
    Create a new stack.Handler that implements the http.Handler interface.
*/

func createHttpHandler(server *Server) http.Handler {
    return &Handler{server: server}
}

/*
    Handles http requests and routes them to stack.server.handle().
*/

func (this *Handler) ServeHTTP(res http.ResponseWriter, req *http.Request) {

    /*
        Pass the res and req into there repective create functions.
        The results of these are then passed to stack.server.handle().
    */

    this.server.handle(CreateRequest(req), CreateResponse(res), 0)
}