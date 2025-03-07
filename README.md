# grpc-demo


```shell
protoc --go_out=. ./proto/basic/*.proto

# Provide Go module name at protoc command. 
# This will not create additional folder by name grpc-demo again as we have mentioned grpc-demo here -> option go_package = "grpc-demo/protogen/basic";
# this will skip generating folder by module name
protoc --go_opt=module=grpc-demo --go_out=. ./proto/basic/*.proto

go mod tidy
```