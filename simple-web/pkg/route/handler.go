package route

import (
	"net/http"
)

type RouterHandler interface {
	http.Handler
	Router
}
