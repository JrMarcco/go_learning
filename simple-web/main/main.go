package main

import (
	"go_learning/simple-web/pkg/hctx"
	"go_learning/simple-web/pkg/server"
	"log"
	"net/http"
)

type RegRsp struct {
	UserId string `json:"userId"`
}

func TestRoute(ctx *hctx.HttpContext) {
	if err := ctx.Ok(RegRsp{UserId: "testUserId"}); err != nil {
		log.Fatalf("%-v\n", err)
	}
}

func main() {
	simpleHttpServer := server.DefaultHttpServer()

	simpleHttpServer.Route(http.MethodGet, "/*", TestRoute)

	if err := simpleHttpServer.Start(); err != nil {
		panic(err)
	}
}
