package api

//===========[STATIC/CACHE]====================================================================================================

var defaultHttpServer = HttpServer{
	Port:      80,
	UseSecure: false,
}

//===========[TYPES]====================================================================================================

//HttpServer defines HttpServer type
type HttpServer struct {
	//Port on which the http server will be listening. Default 80 for http and 443 for https
	Port int

	//UseSecure defines whether the server should be https or http
	UseSecure bool
}

//PRIVATE

//copy makes an identical copy of the HttpServer
func (hs HttpServer) copy() HttpServer {
	return hs
}

//PUBLIC

func (hs HttpServer) Start() {
	if hs.UseSecure {
		//TODO add https listening
	}

}

//===========[FUNCTIONALITY]====================================================================================================

//makeHttpServerSane checks all the value provided in the HttpServer and makes sure that there are no contradictions
func makeHttpServerSane(server *HttpServer) HttpServer {
	if server == nil {
		return defaultHttpServer.copy()
	}

	if server.Port < 0 || server.Port > 65535 {
		server.Port = defaultHttpServer.Port

		//TODO add log message about this
	}

	return server.copy()
}
