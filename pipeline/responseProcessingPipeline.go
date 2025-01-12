package pipeline

import (
	"log"

	"github.com/yasastharinda9511/go_gateway_api/message"
)

type ResponseProcessingPipeline struct {
	requestMessagePool  *message.Pool[*message.HttpRequestMessage]
	responseMessagePool *message.Pool[*message.HttpResponseMessage]
}

// NewProcessingPipeline creates a new instance of ProcessingPipeline
func NewResponseProcessingPipeline(reqMessagePool *message.Pool[*message.HttpRequestMessage], resMessagePool *message.Pool[*message.HttpResponseMessage]) *ResponseProcessingPipeline {
	return &ResponseProcessingPipeline{
		requestMessagePool:  reqMessagePool,
		responseMessagePool: resMessagePool,
	}
}

func (p *ResponseProcessingPipeline) Execute(msg any) {
	// Try to cast the generic Message to HttpResponseMessage
	responseMessage, ok := msg.(*message.HttpResponseMessage)
	if !ok {
		log.Println("Error: Provided message is not of type HttpResponseMessage")
		return
	}

	// Retrieve and process the body

	responseWriter := responseMessage.GetHttpRequestMessage().GetResponseWriter()

	// Example: Write the response to a ResponseWriter (if needed)
	if err := responseMessage.WriteTo(responseWriter); err != nil {
		log.Printf("Error writing response: %v\n", err)
	}

	requestMessage := responseMessage.GetHttpRequestMessage()
	requestMessage.Clear()
	responseMessage.Clear()

	p.requestMessagePool.Put(requestMessage)
	p.responseMessagePool.Put(responseMessage)

}
