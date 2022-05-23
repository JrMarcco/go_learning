package route

type trieTreeNode struct {
	path       string
	children   []*trieTreeNode
	handleFunc HandleFunc
}

func (node *trieTreeNode) findChildNode(path string) (*trieTreeNode, bool) {
	for _, childNode := range node.children {
		if childNode.path == path {
			return childNode, true
		}
	}
	return nil, false
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
