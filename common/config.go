package common

// 此文件中用来获取存配置信息需要用到的变量或常量
// !不要在其他包中用全局变量接收配置，否则拿到的都是空

// Config 全局配置
var Config AllConfig

// AllConfig 所有配置
type AllConfig struct {
	Server Server
	Data   Data
}

// Server 服务信息
type Server struct {
	Http        ConnInfo
	GRPC        ConnInfo
	ServiceName string
}

// ConnInfo 接口信息
type ConnInfo struct {
	Host    string
	Port    int
	Timeout string
}

// Data 数据库配置
type Data struct {
	DataBase DataBase
	Redis    Redis
}

// DataBase 结构化数据库
type DataBase struct {
	Driver string
	Source string
}

// Redis redis数据库
type Redis struct {
	Host string
	Port int
}
