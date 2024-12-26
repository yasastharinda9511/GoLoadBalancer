package rules

import "github.com/yasastharinda9511/go_gateway_api/message"

type Rule interface {
	Evaluate(request *message.HttpRequestMessage) bool
	Print()
}
