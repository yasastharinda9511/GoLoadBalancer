package server

import (
	"fmt"
	"net/http"

	"github.com/yasastharinda9511/go_gateway_api/message"
	"github.com/yasastharinda9511/go_gateway_api/pipeline"
)

type Server struct {
	Port     string
	mux      *http.ServeMux
	pipeline *pipeline.ProcessingPipeline
}

func NewServer(port string, pipeline *pipeline.ProcessingPipeline) *Server {
	return &Server{
		Port:     port,
		mux:      http.NewServeMux(),
		pipeline: pipeline,
	}
}

func (s *Server) RegisterRoutes() {
	s.mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		s.handler(w, r)
	})
}

// Start starts the HTTP server.
func (s *Server) Start() error {
	fmt.Printf("Starting server on port %s...\n", s.Port)
	return http.ListenAndServe(":"+s.Port, s.mux)
}

func (s *Server) handler(w http.ResponseWriter, r *http.Request) {
	// Use your message package to extract data from the request
	requestMessage := message.NewHttpRequestMessage(w, r)
	s.pipeline.Execute(requestMessage)
	// Respond back to the client
}
