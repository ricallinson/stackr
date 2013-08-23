package stack

import(
    "net/http"
)

type httpHandler struct {
    server *Server
}

func createHttpHandler(server *Server) http.Handler {
    return &httpHandler{server: server}
}

func (this *httpHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
    this.server.handle(CreateRequest(req), CreateResponse(res), 0)
}