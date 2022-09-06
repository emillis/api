package api

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
)

//===========[STATIC/CACHE]====================================================================================================

//This will be used as a default HttpServer
var defaultHttpServer = HttpServer{
	Port:               80,
	UseSecure:          false,
	SSLCertificatePath: "",
	PrivateKeyPath:     "",
	Handler:            httprouter.New(),
}

//===========[TYPES]====================================================================================================

//HttpServer defines HttpServer type
type HttpServer struct {
	//Port on which the http server will be listening. Default 80 for http and 443 for https
	Port int

	//UseSecure defines whether the server should be https or http
	UseSecure bool

	//Path to the SSL certificate file. Only needed if UseSecure is set to true
	SSLCertificatePath string

	//Path to the Private Key file. Only used if UseSecure is set to true
	PrivateKeyPath string

	//Defines the handler that this HttpServer will use. In not specified, default http handler is used
	Handler http.Handler

	//You can add request handlers through this
	router *httprouter.Router
}

//Start starts serving requests on the port provided
func (hs *HttpServer) Start() error {
	port := fmt.Sprintf(":%d", hs.Port)

	if hs.UseSecure {
		return http.ListenAndServeTLS(port, hs.SSLCertificatePath, hs.PrivateKeyPath, hs.Handler)
	}

	return http.ListenAndServe(port, hs.Handler)
}

//NewHandler
func (hs *HttpServer) NewResponse() {
}

//===========[FUNCTIONALITY]====================================================================================================

//makeHttpServerSane checks all the value provided in the HttpServer and makes sure that there are no contradictions
func makeHttpServerSane(server *HttpServer) HttpServer {
	if server == nil {
		d := defaultHttpServer
		server = &d
	}

	if server.Port < 0 || server.Port > 65535 {
		log.Fatalf("port specified is out of range. Available 0-65535, got %d", server.Port)
	}

	if server.UseSecure {
		if _, err := os.Stat(server.SSLCertificatePath); err != nil {
			log.Fatalf("could not access ssl certification file in location \"%s\"\n%e", server.SSLCertificatePath, err)
		}
		if _, err := os.Stat(server.PrivateKeyPath); err != nil {
			log.Fatalf("could not access private key file in location \"%s\"\n%e", server.PrivateKeyPath, err)
		}
	}

	if server.Handler == nil {
		server.router = httprouter.New()
		server.Handler = server.router
	}

	return *server
}

//New initiates and returns new HttpServer
func New(s *HttpServer) HttpServer {
	return makeHttpServerSane(s)
}
