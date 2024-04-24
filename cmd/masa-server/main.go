package main

import (
	"context"
	"log"
	"net/http"

	"connectrpc.com/connect"
	"github.com/machshev/masa/config"
	thoughtspb "github.com/machshev/masa/pb/thoughts"
	"github.com/machshev/masa/pb/thoughts/thoughtsconnect"
	"github.com/machshev/masa/thoughts"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type MasaServer struct {
	pad *thoughts.ThoughtPad
}

func (s *MasaServer) Add(
	ctx context.Context,
	req *connect.Request[thoughtspb.AddRequest],
) (*connect.Response[thoughtspb.AddResponse], error) {
	log.Println("Request headers: ", req.Header())

	id, err := s.pad.Add(req.Msg.Thought)
	if err != nil {
		return nil, err
	}

	id_bytes, err := id.MarshalBinary()
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&thoughtspb.AddResponse{
		Uuid: id_bytes,
	})

	return res, nil
}

func (s *MasaServer) List(
	ctx context.Context,
	req *connect.Request[thoughtspb.ListRequest],
) (*connect.Response[thoughtspb.ListResponse], error) {
	log.Println("Request headers: ", req.Header())

	thoughts := s.pad.GetAll()
	pb_thoughts := make([]*thoughtspb.Thought, len(thoughts))

	for i, t := range thoughts {
		pb_t, err := t.AsPB()
		if err != nil {
			return nil, err
		}
		pb_thoughts[i] = pb_t
	}

	res := connect.NewResponse(&thoughtspb.ListResponse{Thoughts: pb_thoughts})

	return res, nil
}

func main() {
	tp, err := thoughts.GetThoughtPad()
	if err != nil {
		log.Fatalf("Failed to open Thought Pad: %v", err)
	}

	server := &MasaServer{pad: tp}
	path, handler := thoughtsconnect.NewThoughtServiceHandler(server)

	mux := http.NewServeMux()
	mux.Handle(path, handler)

	tp.Load()

	http.ListenAndServe(
		config.GetMasaServerURL(),
		h2c.NewHandler(mux, &http2.Server{}),
	)

	tp.Save()
}
