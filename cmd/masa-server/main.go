package main

import (
	"context"
	"log"
	"net/http"

	"connectrpc.com/connect"
	thoughtspb "github.com/machshev/masa/pb/thoughts"
	"github.com/machshev/masa/pb/thoughts/thoughtsconnect"
	"github.com/machshev/masa/thoughts"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type MasaServer struct{}

func (s *MasaServer) Add(
	ctx context.Context,
	req *connect.Request[thoughtspb.AddRequest],
) (*connect.Response[thoughtspb.AddResponse], error) {
	log.Println("Request headers: ", req.Header())
	res := connect.NewResponse(&thoughtspb.AddResponse{
		Id: 0,
	})

	return res, nil
}

func (s *MasaServer) List(
	ctx context.Context,
	req *connect.Request[thoughtspb.ListRequest],
) (*connect.Response[thoughtspb.ListResponse], error) {
	log.Println("Request headers: ", req.Header())
	res := connect.NewResponse(&thoughtspb.ListResponse{})

	return res, nil
}

func main() {
	server := &MasaServer{}
	path, handler := thoughtsconnect.NewThoughtServiceHandler(server)

	mux := http.NewServeMux()
	mux.Handle(path, handler)

	tp, err := thoughts.GetThoughtPad()
	if err != nil {
		log.Fatalf("Failed to open Thought Pad: %v", err)
	}

	tp.Load()

	http.ListenAndServe(
		"localhost:8080",
		h2c.NewHandler(mux, &http2.Server{}),
	)

	tp.Save()
}
