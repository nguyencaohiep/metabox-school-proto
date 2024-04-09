./protoc/bin/protoc --plugin=./protoc/protoc-gen-go-grpc --plugin=./protoc/protoc-gen-go  ./proto/*.proto --proto_path=./proto --go-grpc_out=./golang --go_out=./golang
