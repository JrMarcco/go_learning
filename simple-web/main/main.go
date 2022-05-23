package main

import (
	"go_learning/simple-web/pkg/context"
	"go_learning/simple-web/pkg/server"
	"log"
)

type RegRsp struct {
	UserId string `json:"userId"`
}

func TestRoute(ctx *context.HttpContext) {
	if err := ctx.Ok(RegRsp{UserId: "testUserId"}); err != nil {
		log.Fatalf("%-v", err)
	}
}

func main() {
	simpleHttpServer := server.DefaultHttpServer()

	simpleHttpServer.Route("/reg", TestRoute)

	if err := simpleHttpServer.Start(); err != nil {
		log.Fatalf("%-v", err)
	}
}
