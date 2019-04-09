package main

import (
	"context"
	"github.com/micro/go-micro"
	"go-micro-benchmark/hash/hasher"
	. "go-micro-benchmark/proto/hash"
	"log"
)

var hf = hasher.New()

type Hash struct{}

func (h *Hash) SHA256(ctx context.Context, req *SHA256Request, rsp *SHA256Response) error {
	hash, err := hf.Sum(&req.Str)
	if err != nil {
		return err
	}
	rsp.Hash = hash
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
