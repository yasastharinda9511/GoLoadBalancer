package pathtrie

import (
	"testing"

	"github.com/yasastharinda9511/go_gateway_api/message"
)

// MockRule is a mock implementation of the Rule interface for testing
type MockRule struct {
	id string
}

func (r *MockRule) Evaluate(request *message.HttpRequestMessage) bool {
	return true
}

func (r *MockRule) Print() {
	println("MockRule:", r.id)
}

func TestPathTrie_InsertAndMatch(t *testing.T) {
	trie := NewPathTrie()

	// Insert rules into the trie
	trie.Insert("user/customer/*", &MockRule{id: "rule1"})
	trie.Insert("user/*", &MockRule{id: "rule2"})
	trie.Insert("customer/*", &MockRule{id: "rule3"})
	trie.Insert("user/customer/id", &MockRule{id: "rule4"})

	// Test matching paths
	tests := []struct {
		path     string
		expected []string
	}{
		{"user/customer/id", []string{"rule2", "rule1", "rule4"}},
		{"user/customer/123", []string{"rule2", "rule1"}},
		{"user/123", []string{"rule2"}},
		{"customer/123", []string{"rule3"}},
	}

	for _, test := range tests {
		matchedRules := trie.MatchAllPaths(test.path)
		var matchedRuleIDs []string
		for _, rule := range matchedRules {
			if mockRule, ok := rule.(*MockRule); ok {
				matchedRuleIDs = append(matchedRuleIDs, mockRule.id)
			}
		}

		if len(matchedRuleIDs) != len(test.expected) {
			t.Errorf("For path %s, expected %v, but got %v", test.path, test.expected, matchedRuleIDs)
			continue
		}

		for i, expectedID := range test.expected {
			if matchedRuleIDs[i] != expectedID {
				t.Errorf("For path %s, expected %v, but got %v", test.path, test.expected, matchedRuleIDs)
				break
			}
		}
	}
}

// Test Exact Path Matching
func TestPathTrie_MatchExactPaths(t *testing.T) {

	trie := NewPathTrie()

	// Insert rules into the trie
	trie.Insert("user/customer/*", &MockRule{id: "rule1"})
	trie.Insert("user/*", &MockRule{id: "rule2"})
	trie.Insert("customer/*", &MockRule{id: "rule3"})
	trie.Insert("user/customer/id", &MockRule{id: "rule4"})
	trie.Insert("user/customer/123", &MockRule{id: "rule5"})
	trie.Insert("user/123", &MockRule{id: "rule6"})

	// Test matching paths
	tests := []struct {
		path     string
		expected []string
	}{
		{"user/customer/id", []string{"rule4"}},
		{"user/customer/123", []string{"rule5"}},
		{"user/123", []string{"rule6"}},
	}

	for _, test := range tests {
		matchedRules := trie.MatchExactPaths(test.path)
		var matchedRuleIDs []string
		for _, rule := range matchedRules {
			if mockRule, ok := rule.(*MockRule); ok {
				matchedRuleIDs = append(matchedRuleIDs, mockRule.id)
			}
		}

		if len(matchedRuleIDs) != len(test.expected) {
			t.Errorf("For path %s, expected %v, but got %v", test.path, test.expected, matchedRuleIDs)
			continue
		}

		for i, expectedID := range test.expected {
			if matchedRuleIDs[i] != expectedID {
				t.Errorf("For path %s, expected %v, but got %v", test.path, test.expected, matchedRuleIDs)
				break
			}
		}
	}

}
