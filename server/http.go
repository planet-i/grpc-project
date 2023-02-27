package server

import (
	"net/http"
	"os"
	"path"
	"strings"

	assetfs "github.com/elazarl/go-bindata-assetfs"

	"grpc-project/third_party/swagger"
)

func NewHttp() *http.ServeMux {
	mux := http.NewServeMux()
	// 添加swagger-ui
	serveSwaggerUI(mux)
	mux.HandleFunc("/swagger/", serveSwaggerFile)
	return mux
}

func serveSwaggerFile(w http.ResponseWriter, r *http.Request) {
	// 检查请求的 URL 是否以 "swagger.json" 结尾，如果不是则返回 404 Not Found 响应
	if !strings.HasSuffix(r.URL.Path, "swagger.json") {
		http.NotFound(w, r)
		return
	}
	// 获取当前程序的工作目录
	projectPath, err := os.Getwd()
	if err != nil {
		panic("exec path err")
	}
	// 获取 .swagger.json 文件名称
	jsonFile := strings.TrimPrefix(r.URL.Path, "/swagger/")
	// 获取项目中 .swagger.json 文件的绝对路径
	filePath := path.Join(projectPath, "/api", jsonFile)
	// http.ServeFile() 函数向客户端发送 .swagger.json 文件的绝对路径
	http.ServeFile(w, r, filePath)
}

func serveSwaggerUI(mux *http.ServeMux) {
	fileServer := http.FileServer(&assetfs.AssetFS{
		Asset:    swagger.Asset,
		AssetDir: swagger.AssetDir,
		Prefix:   "third_party/swagger-ui",
	})
	prefix := "/swagger-ui/"
	mux.Handle(prefix, http.StripPrefix(prefix, fileServer))
}
