protoc -I.  --go_out=../  protos/*/*/*.proto
protoc -I.  --go_out=../  protos/*/*.proto
protoc -I . --proto_path=./protos --go_out=../ --go-grpc_opt=require_unimplemented_servers=false --go-grpc_out=../ ./protos/plugin/*.proto

