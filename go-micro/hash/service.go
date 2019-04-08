package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"github.com/micro/go-micro"
	. "go-micro-benchmark/proto/hash"
	"log"
)

type Hash struct{}

func (h *Hash) SHA256(ctx context.Context, req *SHA256Request, rsp *SHA256Response) error {
	hf := sha256.New()
	log.Printf("SHA256 method call. Input string: %s", req.Str)
	hf.Write([]byte(req.Str))
	rsp.Hash = hex.EncodeToString(hf.Sum(nil))
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
