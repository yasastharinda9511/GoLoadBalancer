package message

import "net/http"

type HttpResponseMessage struct {
	*Message
	httpRequestMessage *HttpRequestMessage
	headers            map[string]string
	statusCode         int
	body               []byte
}

// (statusCode int, body []byte, httpRequestMessage *HttpRequestMessage
func NewHttpResponseMessage() *HttpResponseMessage {
	return &HttpResponseMessage{
		Message:            NewMessage(),
		httpRequestMessage: nil,
		headers:            make(map[string]string),
		statusCode:         200,
		body:               nil,
	}
}

func (response *HttpResponseMessage) SetStatusCode(statusCode int) {
	response.statusCode = statusCode
}

func (response *HttpResponseMessage) SetBody(body []byte) {
	response.body = body
}

func (response *HttpResponseMessage) SetHttpRequestMessage(httpRequestMessage *HttpRequestMessage) {
	response.httpRequestMessage = httpRequestMessage
}

func (response *HttpResponseMessage) GetHeaders() map[string]string {
	return response.headers
}

func (response *HttpResponseMessage) GetStatusCode() int {
	return response.statusCode
}

func (response *HttpResponseMessage) GetBody() []byte {
	return response.body
}

func (response *HttpResponseMessage) GetHttpRequestMessage() *HttpRequestMessage {
	return response.httpRequestMessage
}

func (response *HttpResponseMessage) WriteTo(w http.ResponseWriter) error {
	for key, value := range response.headers {
		w.Header().Set(key, value)
	}
	w.WriteHeader(response.statusCode)
	_, err := w.Write(response.body)
	return err
}

func (request *HttpResponseMessage) Clear() {
	for k := range request.headers {
		delete(request.headers, k)
	}
	request.httpRequestMessage = nil
	request.body = nil
}
