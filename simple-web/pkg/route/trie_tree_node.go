package route

type trieTreeNode struct {
	path       string
	children   []*trieTreeNode
	handleFunc HandleFunc
}

func (node *trieTreeNode) findChildNode(path string) (*trieTreeNode, bool) {

	var wildcardNode *trieTreeNode
	for _, childNode := range node.children {
		if childNode.path == path && childNode.path != "*" {
			return childNode, true
		}

		if childNode.path == "*" {
			wildcardNode = childNode
		}
	}
	return wildcardNode, wildcardNode != nil
}

func (node *trieTreeNode) createTrieTree(paths []string, handleFunc HandleFunc) {
	currentNode := node
	for _, path := range paths {
		newNode := &trieTreeNode{
			path:     path,
			children: make([]*trieTreeNode, 0, 4),
		}
		currentNode.children = append(currentNode.children, newNode)
		currentNode = newNode
	}
	currentNode.handleFunc = handleFunc
}
