#!/bin/sh

cd atux_embed/remote &&\
	protoc -I . embed.proto --go-grpc_out=./ &&\
	mv github.com/rodfer0x80/ATux/atux_embed/embed_grpc.pb.go &&\
	rm -rf ./github.com &&\
	cd ..
go build -o ./atux_embed/bin/remote ./atux_embed/run/embed_grpc.pb.go
