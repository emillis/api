package api

type HttpResponder interface {
	Register(hs *HttpServer)
}

type HttpResponse struct {
	//Path defines url path to which this struct will respond
	Path string
}

func (hr *HttpResponse) Register(hs *HttpServer) {

}
