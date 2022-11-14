package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"

	"github.com/rodfer0x80/ATux/atux_api"

	"google.golang.org/grpc"
)

type adminServer struct {
	work, stdout chan *atux_api.Command
	atux_api.UnimplementedAdminServer
}

func MakeAdminServer(work, stdout chan *atux_api.Command) *adminServer {
	s := new(adminServer)
	s.work = work
	s.stdout = stdout
	return s
}

func (s *adminServer) ExecuteCommand(ctx context.Context, cmd *atux_api.Command) (*atux_api.Command, error) {
	var res *atux_api.Command
	go func() {
		s.work <- cmd
	}()
	res = <-s.stdout
	return res, nil
}

type embedServer struct {
	work, stdout chan *atux_api.Command
	atux_api.UnimplementedEmbedServer
}

func MakeEmbedServer(work, stdout chan *atux_api.Command) *embedServer {
	s := new(embedServer)
	s.work = work
	s.stdout = stdout
	return s
}

func (s *embedServer) GetCommand(ctx context.Context, empty *atux_api.Empty) (*atux_api.Command, error) {
	var cmd = new(atux_api.Command)
	select {
	case cmd, ok := <-s.work:
		if ok {
			return cmd, nil
		}
		return cmd, errors.New("Channel error")
	default:
		return cmd, nil
	}
}

func (s *embedServer) SentResult(ctx context.Context, result *atux_api.Command) (*atux_api.Empty, error) {
	s.stdout <- result
	return &atux_api.Empty{}, nil
}

func main() {
	var (
		admin_listener, embed_listener net.Listener
		err                            error
		opts                           []grpc.ServerOption
		work, stdout                   chan *atux_api.Command
	)

	work, stdout = make(chan *atux_api.Command), make(chan *atux_api.Command)
	admin := MakeAdminServer(work, stdout)
	embed := MakeEmbedServer(work, stdout)

	if admin_listener, err = net.Listen("tcp", fmt.Sprintf("localhost:%d", 32001)); err != nil {
		log.Fatal(err)
	}

	if embed_listener, err = net.Listen("tcp", fmt.Sprintf("localhost:%d", 32666)); err != nil {
		log.Fatal(err)
	}

	grpcAdminServer, grpcEmbedServer := grpc.NewServer(opts...), grpc.NewServer(opts...)

	atux_api.RegisterAdminServer(grpcAdminServer, admin)
	atux_api.RegisterEmbedServer(grpcEmbedServer, embed)

	go func() {
		grpcEmbedServer.Serve(embed_listener)
	}()
	grpcAdminServer.Serve(admin_listener)
}
