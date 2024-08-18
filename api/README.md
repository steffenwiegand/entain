
**package proto**
When new services get added to the api, ensure to include the new .proto file within the proto/api..go file

//go:generate protoc -I . --go_out . --go_opt paths=source_relative --go-grpc_out . --go-grpc_opt paths=source_relative --grpc-gateway_out . --grpc-gateway_opt paths=source_relative **racing/racing.proto sports/sports.proto** --experimental_allow_proto3_optional

Currently the 'go generate ./...' command will only look for sports.proto and racing.proto file.