package pipeline

import (
	"log"

	"github.com/yasastharinda9511/go_gateway_api/message"
)

type ResponseProcessingPipeline struct{}

// NewProcessingPipeline creates a new instance of ProcessingPipeline
func NewResponseProcessingPipeline() *ResponseProcessingPipeline {
	return &ResponseProcessingPipeline{}
}

func (p *ResponseProcessingPipeline) Execute(msg any) {
	// Try to cast the generic Message to HttpResponseMessage
	responseMessage, ok := msg.(*message.HttpResponseMessage)
	if !ok {
		log.Println("Error: Provided message is not of type HttpResponseMessage")
		return
	}

	// Access data from the HttpResponseMessage
	log.Printf("Processing response with UID: %s\n", responseMessage.GetUID())
	log.Printf("Status Code: %d\n", responseMessage.GetStatusCode())

	// Retrieve and process the body
	body := responseMessage.GetBody()
	log.Printf("Body: %s\n", string(body))

	responseWriter := responseMessage.GetHttpRequestMessage().GetResponseWriter()

	// Example: Write the response to a ResponseWriter (if needed)
	if err := responseMessage.WriteTo(responseWriter); err != nil {
		log.Printf("Error writing response: %v\n", err)
	}
}
