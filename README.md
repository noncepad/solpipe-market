# Protobufs for Common Markets



# Build Libraries

## Go


```bash
protoc --experimental_allow_proto3_optional --proto_path=./txproc/proto --go-grpc_out=. --go_out=. $(ls ./txproc/proto)
protoc --experimental_allow_proto3_optional --proto_path=./geyser/proto --go-grpc_out=. --go_out=. $(ls ./geyser/proto)
protoc --experimental_allow_proto3_optional --proto_path=./solpipe/proto --go-grpc_out=. --go_out=. $(ls ./solpipe/proto)
rm -r ./go
mv github.com/noncepad/solpipe-market/go ./
```


