package route

import (
	"go_learning/simple-web/pkg/context"
	"log"
	"net/http"
	"strings"
)

type TrieTreeHandler struct {
	rootNode *trieTreeNode
}

func (handler *TrieTreeHandler) Route(method string, path string, handleFunc HandleFunc) {

	paths := strings.Split(strings.Trim(path, "/"), "/")
	currentNode := handler.rootNode

	for index, path := range paths {
		if childNode, ok := currentNode.findChildNode(path); ok {
			currentNode = childNode
		} else {
			currentNode.createTrieTree(paths[index:], handleFunc)
			return
		}
	}
}

func (handler *TrieTreeHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	ctx := context.BuildHttpContext(writer, request)
	if handleFunc, ok := handler.findRouterHandleFunc(request.URL.Path); ok {
		handleFunc(ctx)
		return
	}

	if err := ctx.NotFound(); err != nil {
		log.Fatalf("%-v\n", err)
	}
}

func (handler *TrieTreeHandler) findRouterHandleFunc(path string) (HandleFunc, bool) {
	paths := strings.Split(strings.Trim(path, "/"), "/")

	currentNode := handler.rootNode

	for _, path := range paths {
		if childNode, ok := currentNode.findChildNode(path); ok {
			currentNode = childNode
		}
	}

	if currentNode.handleFunc != nil {
		return currentNode.handleFunc, true
	}

	return nil, false
}

// 确保 RouterHandler 实现 Handler
var _ RouterHandler = &TrieTreeHandler{}

func NewTrieTreeHandler() RouterHandler {
	return &TrieTreeHandler{
		rootNode: &trieTreeNode{},
	}
}
