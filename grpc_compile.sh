cd pluginClient
protoc -I . --proto_path=./proto --go_out=./pb --go-grpc_out=./pb ./proto/*.proto