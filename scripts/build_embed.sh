#!/bin/sh

cd atux_embed/remote &&\
	protoc -I . embed.proto --go-grpc_out=./ &&\
	mv github.com/rodfer0x80/GoRPC/atux_embed/embed_grpc.pb.go &&\
	rm -rf ./github.com &&\
	cd ..
go build -o ./atux-embed/bin/remote ./atux-embed/run/embed_grpc.pb.go
