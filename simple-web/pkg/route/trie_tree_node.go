package route

const (
	nodeTypeRoot      = iota // 根节点
	nodeTypeWildcard         // 通配符匹配
	nodeTypePathParam        // 路径参数匹配
	nodeTypeReg              // 正则匹配
	nodeTypeComplete         // 完全匹配
)

type trieTreeNode struct {
	children   []*trieTreeNode
	handleFunc HandleFunc

	path string
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
