package framework

import (
	"errors"
	"strings"
)

type TrieTree struct {
	root *trieTreeNode
}

type trieTreeNode struct {
	isLeafNode bool
	segment    string
	handler    ControllerHandler
	children   []*trieTreeNode
}

func newNode() *trieTreeNode {
	return &trieTreeNode{
		isLeafNode: false,
		children:   []*trieTreeNode{},
	}
}

func NewTrieTree() *TrieTree {
	return &TrieTree{root: newNode()}
}

func (tree *TrieTree) AddRouter(url string, handler ControllerHandler) error {
	url = strings.TrimLeft(url, "/")

	rootNode := tree.root
	if rootNode.findMatchNode(url) != nil {
		return errors.New("router exist: " + url)
	}

	segments := strings.Split(url, "/")
	for index, segment := range segments {
		var objNode *trieTreeNode
		isLast := index == len(segments)-1

		children := rootNode.filterChildren(segment)
		if len(children) > 0 {
			for _, child := range children {
				if child.segment == segment {
					objNode = child
					break
				}
			}
		}

		if objNode == nil {
			child := newNode()
			child.segment = segment
			if isLast {
				child.isLeafNode = true
				child.handler = handler
			}
			rootNode.children = append(rootNode.children, child)
			objNode = child
		}
		rootNode = objNode
	}
	return nil
}

func (tree *TrieTree) FindHandler(url string) ControllerHandler {
	matchNode := tree.root.findMatchNode(url)
	if matchNode != nil {
		return matchNode.handler
	}
	return nil
}

func (node *trieTreeNode) filterChildren(segment string) []*trieTreeNode {
	if len(node.children) == 0 {
		return nil
	}

	if isWildSegment(segment) {
		return node.children
	}

	nodes := make([]*trieTreeNode, 0, len(node.children))
	for _, child := range node.children {
		if isWildSegment(child.segment) || child.segment == segment {
			nodes = append(nodes, child)
		}
	}

	return nodes
}

func (node *trieTreeNode) findMatchNode(url string) *trieTreeNode {
	url = strings.TrimLeft(url, "/")
	segments := strings.SplitN(url, "/", 2)

	children := node.filterChildren(segments[0])
	if children == nil || len(children) == 0 {
		return nil
	}

	if len(segments) == 1 {
		for _, child := range children {
			if child.isLeafNode {
				return child
			}
		}
		return nil
	}

	for _, child := range children {
		matchNode := child.findMatchNode(segments[1])
		if matchNode != nil {
			return matchNode
		}
	}
	return nil
}

func isWildSegment(segment string) bool {
	return strings.HasPrefix(segment, ":")
}
