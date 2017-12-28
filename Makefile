plugins = grpc
target = go
protoc_location = rpc

.PHONY: rpc
rpc:
	protoc -I $(protoc_location) -I/usr/local/include -I /Users/radu/go/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --$(target)_out=plugins=$(plugins):$(protoc_location) $(protoc_location)/*.proto --grpc-gateway_out=logtostderr=true:rpc/