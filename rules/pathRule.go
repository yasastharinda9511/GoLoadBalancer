package rules

import (
	"fmt"

	"github.com/yasastharinda9511/go_gateway_api/message"
)

type PathRule struct {
	path string
}

func NewPathRule(path string) *PathRule {
	return &PathRule{
		path: path,
	}
}

// Evaluate checks if the request contains the specified header with the specified value
func (r *PathRule) Evaluate(request *message.HttpRequestMessage) bool {
	return request.GetURL() == r.path
}

func (r *PathRule) Print() {
	fmt.Printf("	RuleType: PathRule\n	Path: %s\n", r.path)
}

func (r *PathRule) GetPath() string {
	return r.path
}
