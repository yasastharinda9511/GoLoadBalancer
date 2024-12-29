package message

import (
	"log"
	"net/http"
)

type HttpRequestMessage struct {
	*Message
	headers        map[string]string
	method         string
	query          map[string]string
	responseWriter http.ResponseWriter
	httpRequest    *http.Request
}

func NewHttpRequestMessage(w http.ResponseWriter, r *http.Request) *HttpRequestMessage {
	// Extract headers into a map

	log.Println("Initializing HttpRequestMessage...")

	log.Printf("Method %s \n", r.Method)

	headers := make(map[string]string)
	for key, values := range r.Header {
		log.Printf("Header - Key: %s, Value: %s\n", key, values[0])
		headers[key] = values[0] // Take the first value for simplicity
	}

	// Extract query parameters into a map
	queryParams := make(map[string]string)
	for key, values := range r.URL.Query() {
		log.Printf("Query Parameter - Key: %s, Value: %s\n", key, values[0])
		queryParams[key] = values[0] // Take the first value for simplicity
	}

	return &HttpRequestMessage{
		Message:        NewMessage(),
		headers:        headers,
		method:         r.Method,
		query:          queryParams,
		responseWriter: w,
		httpRequest:    r,
	}
}

func (request *HttpRequestMessage) GetResponseWriter() http.ResponseWriter {
	return request.responseWriter
}

func (request *HttpRequestMessage) GetHeaders() map[string]string {
	return request.headers
}

func (request *HttpRequestMessage) GetQueryParams() map[string]string {
	return request.query
}
func (request *HttpRequestMessage) GetURL() string {
	return request.httpRequest.URL.String()
}
