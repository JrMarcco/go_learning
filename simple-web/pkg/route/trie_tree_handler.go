package route

import (
	"fmt"
	"go_learning/simple-web/pkg/hctx"
	"net/http"
	"strings"
)

type TrieTreeHandler struct {
	rootNode *trieTreeNode
}

func (handler *TrieTreeHandler) Route(method string, path string, handleFunc HandleFunc) {

	if err := handler.validateRouterPath(path); err != nil {
		panic(err)
	}

	paths := strings.Split(method+"/"+strings.Trim(path, "/"), "/")
	currentNode := handler.rootNode

	for index, path := range paths {
		if childNode, ok := currentNode.findChildNode(path); ok {
			currentNode = childNode
		} else {
			currentNode.createTrieTree(paths[index:], handleFunc)
			return
		}
	}
	currentNode.handleFunc = handleFunc
}

func (handler *TrieTreeHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	ctx := hctx.BuildHttpContext(writer, request)
	if handleFunc, ok := handler.findRouterHandleFunc(request.Method, request.URL.Path); ok {
		handleFunc(ctx)
		return
	}

	if err := ctx.NotFound(); err != nil {
		fmt.Printf("%-v\n", err)
	}
}

func (handler *TrieTreeHandler) findRouterHandleFunc(method string, path string) (HandleFunc, bool) {
	paths := strings.Split(method+"/"+strings.Trim(path, "/"), "/")

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

func (handler *TrieTreeHandler) validateRouterPath(path string) error {
	wildcardPosition := strings.Index(path, "*")

	if wildcardPosition > 0 {
		if wildcardPosition != len(path)-1 || path[wildcardPosition-1] != '/' {
			return fmt.Errorf("### Invalid router path: %s ###\n", path)
		}
	}

	return nil
}

// 确保 RouterHandler 实现 Handler
var _ RouterHandler = &TrieTreeHandler{}

func NewTrieTreeHandler() RouterHandler {
	return &TrieTreeHandler{
		rootNode: &trieTreeNode{},
	}
}
