package message

import (
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

func NewHttpRequestMessage() *HttpRequestMessage {
	return &HttpRequestMessage{
		Message:        NewMessage(),
		headers:        make(map[string]string),
		method:         "",
		query:          make(map[string]string),
		responseWriter: nil,
		httpRequest:    nil,
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

func (request *HttpRequestMessage) SetResponseWriter(responseWriter http.ResponseWriter) {
	request.responseWriter = responseWriter
}

func (request *HttpRequestMessage) SetHttpRequest(httpRequest *http.Request) {
	request.httpRequest = httpRequest
}

func (request *HttpRequestMessage) SetHeaders(r *http.Request) {
	for key, values := range r.Header {
		request.headers[key] = values[0] // Take the first value for simplicity
	}
}

func (request *HttpRequestMessage) SetQueryParams(r *http.Request) {
	for key, values := range r.URL.Query() {
		request.query[key] = values[0] // Take the first value for simplicity
	}
}

func (request *HttpRequestMessage) SetMethod(r *http.Request) {
	request.method = r.Method
}

func (request *HttpRequestMessage) Clear() {
	for k := range request.headers {
		delete(request.headers, k)
	}

	for k := range request.query {
		delete(request.query, k)
	}

	request.responseWriter = nil
	request.httpRequest = nil
}
