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
	// 在swagger界面输入此模式来加载文件 /swagger/xxx.swagger.json
	mux.HandleFunc("/swagger/", serveSwaggerFile)
	return mux
}

func serveSwaggerFile(w http.ResponseWriter, r *http.Request) {
	// 检查请求的 URL 是否以 "swagger.json" 结尾，如果不是则返回 404 Not Found 响应
	if !strings.HasSuffix(r.URL.Path, "swagger.json") {
		http.NotFound(w, r)
		return
	}
	// 获取项目中 .swagger.json 文件的绝对路径
	projectPath, err := os.Getwd() // 获取当前程序的工作目录
	if err != nil {
		panic("exec path err")
	}
	jsonFile := strings.TrimPrefix(r.URL.Path, "/swagger/") // 获取 .swagger.json 文件名称
	filePath := path.Join(projectPath, "/api", jsonFile)    // 拼接项目中 .swagger.json 文件的路径
	// http.ServeFile() 函数向客户端发送 .swagger.json 文件的绝对路径
	http.ServeFile(w, r, filePath)
}

func serveSwaggerUI(mux *http.ServeMux) {
	fileServer := http.FileServer(&assetfs.AssetFS{
		Asset:    swagger.Asset,
		AssetDir: swagger.AssetDir,
		Prefix:   "third_party/swagger-ui",
	})
	// 通过http://host:port/swagger-ui/ 打开swagger界面
	prefix := "/swagger-ui/"
	mux.Handle(prefix, http.StripPrefix(prefix, fileServer))
}
