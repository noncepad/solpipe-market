#!/bin/bash

set -euo pipefail

protoc --experimental_allow_proto3_optional --proto_path=./txproc/proto --go-grpc_out=. --go_out=. $(ls ./txproc/proto)
protoc --experimental_allow_proto3_optional --proto_path=./geyser/proto --go-grpc_out=. --go_out=. $(ls ./geyser/proto)
protoc --experimental_allow_proto3_optional --proto_path=./solpipe --go-grpc_out=. --go_out=. $(ls ./solpipe)
protoc --experimental_allow_proto3_optional --proto_path=./catscope/proto --go-grpc_out=. --go_out=. $(ls ./catscope/proto)
protoc --experimental_allow_proto3_optional --proto_path=./bot/proto --go-grpc_out=. --go_out=. $(ls ./bot/proto)
rm -rf ./go/proto 2>/dev/null || true
rsync -avz github.com/noncepad/solpipe-market/go/ ./go/
rm -r github.com
