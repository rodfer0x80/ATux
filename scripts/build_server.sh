#!/bin/sh

cd atux_api &&\
	protoc -I . api.proto --go_out=. --go-grpc_out=. &&\
	mv github.com/rodfer0x80/ATux/atux_api/api_grpc.pb.go github.com/rodfer0x80/GoRPC/atux_api/api.pb.go . &&\
	rm -rf ./github.com &&\
	cd ..

go build -o ./atux_server/bin/main ./atux_server/main.go
