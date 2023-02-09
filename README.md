# grpc-project

```
go get go get google.golang.org/grpc
go get google.golang.org/protobuf  
> protoc --version
libprotoc 3.21.9
> protoc-gen-go --version
protoc-gen-go.exe v1.28.1
> protoc-gen-go-grpc --version
protoc-gen-go-grpc 1.2.0

protoc --go_out=:. --go-grpc_out=:. ./api/proto/*.proto
```