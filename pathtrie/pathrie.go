package pathtrie

import (
	"fmt"
	"strings"
)

// TrieNode represents a node in the path trie
type TrieNode struct {
	children map[string]*TrieNode
	isEnd    bool
	rule_ids []string
}

// PathTrie represents the trie structure
type PathTrie struct {
	root *TrieNode
}

// NewPathTrie creates a new PathTrie
func NewPathTrie() *PathTrie {
	return &PathTrie{
		root: &TrieNode{children: make(map[string]*TrieNode)},
	}
}

// Insert inserts a path and associated rules into the trie
func (t *PathTrie) Insert(path string, rule_id string) {
	parts := strings.Split(path, "/")
	node := t.root
	for _, part := range parts {
		if _, exists := node.children[part]; !exists {
			node.children[part] = &TrieNode{children: make(map[string]*TrieNode)}
		}
		node = node.children[part]
	}
	node.isEnd = true
	node.rule_ids = append(node.rule_ids, rule_id)

}

// Match matches a request path against the trie and returns the associated rules
func (t *PathTrie) MatchAllPaths(path string) []string {
	parts := strings.Split(path, "/")
	node := t.root
	var matchedRules []string

	for _, part := range parts {

		if child, exists := node.children["*"]; exists && child.isEnd {
			matchedRules = append(matchedRules, child.rule_ids...)
		}
		if child, exists := node.children[part]; exists {
			node = child
		}
		if node.isEnd {
			matchedRules = append(matchedRules, node.rule_ids...)
		}
	}
	return matchedRules
}

func (t *PathTrie) MatchExactPaths(path string) []string {
	parts := strings.Split(path, "/")
	node := t.root
	var matchedRules []string

	for _, part := range parts {
		if child, exists := node.children[part]; exists {
			node = child
		}
		if node.isEnd {
			matchedRules = append(matchedRules, node.rule_ids...)
		}
	}
	return matchedRules
}

func (t *PathTrie) MatchPrefixPaths(path string) []string {
	parts := strings.Split(path, "/")
	node := t.root
	var matchedRules []string

	for _, part := range parts {

		if child, exists := node.children["*"]; exists && child.isEnd {
			matchedRules = append(matchedRules, child.rule_ids...)
		}
		if child, exists := node.children[part]; exists {
			node = child
		}
	}
	return matchedRules
}

// RuleStore holds a map of rule lists indexed by an ID
type RuleStore struct {
	trie *PathTrie
}

// NewRuleStore creates a new RuleStore
func NewRuleStore() *RuleStore {
	return &RuleStore{
		trie: NewPathTrie(),
	}
}

// AddRule adds a rule to the store under the given path
func (rs *RuleStore) AddRule(path string, rule_id string) {
	rs.trie.Insert(path, rule_id)
}

// GetRules retrieves the list of rules for the given path
func (rs *RuleStore) GetRules(path string) []string {
	return rs.trie.MatchAllPaths(path)
}

// PrintAllRules prints all rules in the store
func (rs *RuleStore) PrintAllRules() {
	// Implement a method to traverse the trie and print all rules
	var traverse func(node *TrieNode, path string)
	traverse = func(node *TrieNode, path string) {
		if node.isEnd {
			println("Path:", path)
			for _, rule := range node.rule_ids {
				fmt.Printf("	Rule: %s\n", rule)
			}
		}
		for part, child := range node.children {
			traverse(child, path+"/"+part)
		}
	}
	traverse(rs.trie.root, "")
}
