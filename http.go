package stackr

import(
    "net/http"
)

/*
    A HTTP Handler for stack.
*/

type handler struct {
    server *server
}

/*
    Create a new handler that implements the http.Handler interface.
*/

func createHttpHandler(server *server) http.Handler {
    return &handler{server: server}
}

/*
    Handles http requests and routes them to stack.server.handle().
*/
func (this *handler) ServeHTTP(res http.ResponseWriter, req *http.Request) {

    /*
        Pass the res and req into there repective create functions.
        The results of these are then passed to stack.server.handle().
    */

    this.server.handle(createRequest(req), createResponse(res), 0)
}