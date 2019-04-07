package main

import (

	"context"
	"crypto/sha256"
	"github.com/micro/go-micro"
	hash "github.com/nskondratev/microservices-frameworks-comparison/go-micro/proto"
)

type Hash struct{}

func (h *Hash) SHA256(ctx context.Context, req *hash.SHA256Request, rsp *hash.SHA256Response) error {
	h := sha256.New()

}

func main() {
	service := micro.NewService(
		micro.Name("hash"),
	)

	service.Init()


}
