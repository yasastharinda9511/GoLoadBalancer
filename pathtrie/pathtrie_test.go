package pathtrie

import (
	"testing"
)

func TestPathTrie_InsertAndMatch(t *testing.T) {
	trie := NewPathTrie()

	// Insert rules into the trie
	trie.Insert("user/customer/*", "rule1")
	trie.Insert("user/*", "rule2")
	trie.Insert("customer/*", "rule3")
	trie.Insert("user/customer/id", "rule4")

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
		matchedRuleIDs := append([]string{}, matchedRules...)

		if len(matchedRuleIDs) != len(test.expected) {
			t.Errorf("For path %s, expected %v, but got %v", test.path, test.expected, matchedRuleIDs)
			break
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
	trie.Insert("user/customer/123", "rule1")
	trie.Insert("user/*", "rule2")
	trie.Insert("customer/*", "rule3")
	trie.Insert("user/customer/id", "rule4")
	trie.Insert("user/123", "rule5")
	trie.Insert("customer/123", "rule6")

	// Test matching paths
	tests := []struct {
		path     string
		expected []string
	}{
		{"user/customer/id", []string{"rule4"}},
		{"user/customer/123", []string{"rule1"}},
		{"user/123", []string{"rule5"}},
		{"customer/123", []string{"rule6"}},
	}

	for _, test := range tests {
		matchedRules := trie.MatchExactPaths(test.path)
		matchedRuleIDs := append([]string{}, matchedRules...)

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

func TestPathTrie_MatchPrefixPaths(t *testing.T) {

	trie := NewPathTrie()

	// Insert rules into the trie
	trie.Insert("/user/customer/*", "rule1")
	trie.Insert("/user/*", "rule2")
	trie.Insert("/api/*", "rule3")

	// Test matching paths
	tests := []struct {
		path     string
		expected []string
	}{
		{"/user/customer/id", []string{"rule2", "rule1"}},
		{"/api/123", []string{"rule3"}},
	}

	for _, test := range tests {
		matchedRules := trie.MatchPrefixPaths(test.path)
		matchedRuleIDs := append([]string{}, matchedRules...)

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
