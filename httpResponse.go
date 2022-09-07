package api

type HttpResponder interface {
	Register(hs *HttpServer)
}

type HttpResponse struct {
}

func (hr *HttpResponse) Register(hs *HttpServer) {

}
