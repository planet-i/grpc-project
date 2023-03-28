# grpc-project

```
D:.
├─.idea
├─api
│  └─simple
│      ├─handler   // 1. 
│      │  └─tag
│      └─v1
│          └─proto
├─app
│  └─simple
│      └─cmd
│          └─server
├─common
├─conf
├─logic // 2. 
│  └─tag
├─repo // 4. 
│  └─tag
├─server
├─service // 3. 
│  └─tag
└─third_party
    ├─google
    │  └─api
    ├─swagger
    └─swagger-ui
```

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

// ----生成文件方法----
cd api\simple\v1 && protoc --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. --grpc-gateway_=paths=source_relative:. --grpc-gateway_out=paths=source_relative:. --openapiv2_out=. --proto_path=./proto --proto_path=../../../third_party ./proto/*.proto

// ----生成datafile.go文件----
go-bindata --nocompress -pkg swagger -o third_party/swagger/datafile.go third_party/swagger-ui/...
```