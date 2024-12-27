package pipeline

import "github.com/yasastharinda9511/go_gateway_api/message"

// ProcessingPipeline defines the interface for the processing pipeline
type ProcessingPipeline interface {
	Execute(req *message.HttpRequestMessage) error
}
