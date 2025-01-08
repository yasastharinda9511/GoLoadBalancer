package rules

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/yasastharinda9511/go_gateway_api/message"
)

func TestHeaderRule_Evaluate(t *testing.T) {
	// Create a new HeaderRule
	rule := NewHeaderRule("Content-Type", "application/json")

	// Create a mock HttpRequestMessage with the expected header
	req := &message.HttpRequestMessage{
		headers: http.Header{
			"Content-Type": []string{"application/json"},
		},
	}

	// Evaluate the rule
	if !rule.Evaluate(req) {
		t.Errorf("Expected HeaderRule to evaluate to true, but got false")
	}

	// Create a mock HttpRequestMessage without the expected header
	req = &message.HttpRequestMessage{
		Headers: http.Header{
			"Content-Type": []string{"text/plain"},
		},
	}

	// Evaluate the rule
	if rule.Evaluate(req) {
		t.Errorf("Expected HeaderRule to evaluate to false, but got true")
	}
}

func TestHeaderRule_Print(t *testing.T) {
	// Create a new HeaderRule
	rule := NewHeaderRule("Content-Type", "application/json")

	// Capture the output of the Print method
	output := captureOutput(func() {
		rule.Print()
	})

	// Expected output
	expectedOutput := "	RuleType: HeaderRule\n	Header: Content-Type\n	HeaderValue: application/json\n"

	// Compare the captured output with the expected output
	if output != expectedOutput {
		t.Errorf("Expected output:\n%s\nBut got:\n%s", expectedOutput, output)
	}
}

// captureOutput captures the output of a function
func captureOutput(f func()) string {
	// Save the original stdout
	old := fmt.Print
	defer func() { fmt.Print = old }()

	// Create a new buffer to capture the output
	var buf strings.Builder
	fmt.Print = func(a ...interface{}) (n int, err error) {
		return buf.WriteString(fmt.Sprint(a...))
	}

	// Call the function
	f()

	// Return the captured output
	return buf.String()
}
