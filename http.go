package connect

import(
    "net/http"
)

type httpHandler struct {
    server *ConnectServer
}

func createHttpHandler(server *ConnectServer) http.Handler {
    return &httpHandler{server: server}
}

func (this *httpHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
    this.server.handle(CreateRequest(req), CreateResponse(res), 0)
}