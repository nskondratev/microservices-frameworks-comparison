package main

import (
	"context"
	"github.com/facebookgo/grace/gracehttp"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-web"
	. "go-micro-benchmark/proto/hash"
	"log"
	"net/http"
)

type ApiService struct{}

type ErrorJSONResponse struct {
	Error string `json:"error"`
}

type SHA256JSONResponse struct {
	Hash string `json:"hash"`
}

var (
	cl HashService
)

func (as *ApiService) SHA256(c echo.Context) error {

	str := c.QueryParam("string")

	response, err := cl.SHA256(context.TODO(), &SHA256Request{
		Str: str,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, &ErrorJSONResponse{
			Error: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, &SHA256JSONResponse{
		Hash: response.Hash,
	})
}

func main() {
	service := web.NewService(
		web.Name("api"),
	)

	if err := service.Init(); err != nil {
		log.Fatalf("Error while init API service: %s", err.Error())
	}

	cl = NewHashService("hash", client.DefaultClient)

	apiService := new(ApiService)

	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Recover())

	hg := e.Group("/hash")
	hg.GET("/sha256", apiService.SHA256)

	e.Server.Addr = ":8080"
	log.Fatal(gracehttp.Serve(e.Server))
}
