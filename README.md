# grpc-project

```
go get google.golang.org/grpc
go get google.golang.org/protobuf
go get github.com/grpc-ecosystem/grpc-gateway/v2 
> protoc --version
libprotoc 3.21.9
> protoc-gen-go --version
protoc-gen-go.exe v1.28.1
> protoc-gen-go-grpc --version
protoc-gen-go-grpc 1.2.0

// ----生成文件方法一
option go_package = "./api;api";
// 生成在api目录
protoc --go_out=. --go-grpc_out=. --grpc-gateway_out=. --openapiv2_out=./api --proto_path=./api/proto --proto_path=./third_party ./api/proto/*.proto
// 生成在当前目录(根目录)
protoc --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. --grpc-gateway_out=paths=source_relative:. --openapiv2_out=:. --proto_path=./api/proto --proto_path=./third_party ./api/proto/*.proto
// ----生成文件方法二
option go_package = "./;api";
// 生成在api目录
cd api && protoc --go_out=. --go-grpc_out=. --grpc-gateway_out=. --openapiv2_out=. --proto_path=./proto --proto_path=../third_party ./proto/*.proto
// 生成在当前目录(api目录)
cd api && protoc --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. --grpc-gateway_out=paths=source_relative:. --openapiv2_out=. --proto_path=./proto --proto_path=../third_party ./proto/*.proto
```