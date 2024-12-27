package message

import "net/http"

type HttpResponseMessage struct {
	*Message
	httpRequestMessage *HttpRequestMessage
	headers            map[string]string
	statusCode         int
	body               []byte
}

func NewHttpResponseMessage(statusCode int, body []byte, httpRequestMessage *HttpRequestMessage) *HttpResponseMessage {
	return &HttpResponseMessage{
		Message:            NewMessage(),
		httpRequestMessage: httpRequestMessage,
		headers:            make(map[string]string),
		statusCode:         statusCode,
		body:               body,
	}
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
