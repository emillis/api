package api

import (
	"fmt"
	"log"
	"net/http"
)

//===========[STATIC/CACHE]====================================================================================================

//This will be used as a default HttpServer
var defaultHttpServer = HttpServer{
	Port:               80,
	UseSecure:          false,
	SSLCerticifatePath: "",
	PrivateKeyPath:     "",
}

//===========[TYPES]====================================================================================================

//HttpServer defines HttpServer type
type HttpServer struct {
	//Port on which the http server will be listening. Default 80 for http and 443 for https
	Port int

	//UseSecure defines whether the server should be https or http
	UseSecure bool

	//Path to the SSL certificate file. Only needed if UseSecure is set to true
	SSLCerticifatePath string

	//Path to the Private Key file. Only used if UseSecure is set to true
	PrivateKeyPath string

	//Defines the handler that this HttpServer will use. In not specified, default http handler is used
	Handler http.Handler
}

//PRIVATE

//copy makes an identical copy of the HttpServer
func (hs HttpServer) copy() HttpServer {
	return hs
}

//PUBLIC

//Start starts serving requests on the port provided
func (hs HttpServer) Start() error {
	port := fmt.Sprintf(":%d", hs.Port)

	if hs.UseSecure {
		return http.ListenAndServeTLS(port, hs.SSLCerticifatePath, hs.PrivateKeyPath, nil)
	}

	return http.ListenAndServe(port, nil)
}

//NewHandler
func (hs *HttpServer) NewHandler() {

}

//===========[FUNCTIONALITY]====================================================================================================

//makeHttpServerSane checks all the value provided in the HttpServer and makes sure that there are no contradictions
func makeHttpServerSane(server *HttpServer) HttpServer {
	if server == nil {
		return defaultHttpServer.copy()
	}

	if server.Port < 0 || server.Port > 65535 {
		log.Fatalf("port specified is out of range. Available 0-65535, got %d", server.Port)
	}

	return server.copy()
}

//New initiates and returns new HttpServer
func New(s *HttpServer) HttpServer {
	if s == nil {
		return defaultHttpServer.copy()
	}

	return makeHttpServerSane(s)
}
