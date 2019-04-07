package main

import (
	"context"
	"crypto/sha256"
	"github.com/micro/go-micro"
	. "github.com/nskondratev/microservices-frameworks-comparison/go-micro/proto/hash"
	"log"
)

type Hash struct{}

func (h *Hash) SHA256(ctx context.Context, req *SHA256Request, rsp *SHA256Response) error {
	hf := sha256.New()
	hf.Write([]byte(req.Str))
	rsp.Hash = string(hf.Sum(nil))
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("hash"),
	)

	service.Init()

	if err := RegisterHashHandler(service.Server(), new(Hash)); err != nil {
		log.Fatalf("Error while registering Hash service handler: %s", err.Error())
	}

	if err := service.Run(); err != nil {
		log.Fatalf("Error while running Hash service: %s", err.Error())
	}
}
