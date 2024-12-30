package rules

import (
	"fmt"

	"github.com/yasastharinda9511/go_gateway_api/message"
)

type HeaderRule struct {
	header      string
	headerValue string
}

func NewHeaderRule(header, headerValue string) *HeaderRule {
	return &HeaderRule{
		header:      header,
		headerValue: headerValue,
	}
}

// Evaluate checks if the request contains the specified header with the specified value
func (r *HeaderRule) Evaluate(request *message.HttpRequestMessage) bool {
	return request.GetHeaders()[r.header] == r.headerValue
}

func (r *HeaderRule) Print() {
	fmt.Printf("	RuleType: HeaderRule\n	Header: %s\n	HeaderValue: %s\n", r.header, r.headerValue)
}
