package main

import (
	"context"
	"fmt"
	"github.com/go-leo/leo/v3/authx/jwtx"
	"github.com/go-leo/leo/v3/example/api/helloworld/v1"
	"github.com/go-leo/leo/v3/transportx/httptransportx"
	"github.com/golang-jwt/jwt/v5"
	"log"
)

func main() {
	// success
	// jwt 中间件
	mdw := jwtx.Client([]byte("jwt_key_secret"))
	client := helloworld.NewGreeterHttpClient("localhost:60051", httptransportx.WithMiddleware(mdw))
	// 向ctx中注入jwt信息
	ctx := jwtx.NewContentWithClaims(context.Background(), jwt.MapClaims{"user_id": "123456"})
	r, err := client.SayHello(ctx, &helloworld.HelloRequest{Name: "ubuntu"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())

	// error
	// jwt 中间件
	mdw = jwtx.Client([]byte("wrong_jwt_key_secret"))
	client = helloworld.NewGreeterHttpClient("localhost:60051", httptransportx.WithMiddleware(mdw))
	// 向ctx中注入jwt信息
	ctx = jwtx.NewContentWithClaims(context.Background(), jwt.MapClaims{"user_id": "123456"})
	r, err = client.SayHello(ctx, &helloworld.HelloRequest{Name: "mint"})
	if err == nil {
		panic(err)
	}
	fmt.Println(err)
}
