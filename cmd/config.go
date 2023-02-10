package main

import (
	"os"

	"grpc-project/common"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func initConfig() error {
	// 定义配置文件路径选项
	var configFile string
	pflag.StringVar(&configFile, "configFile", "conf/config.yaml", "server config file")
	pflag.Parse()

	if configFile != "" {
		// 获取配置文件的路径，运行时未传入参数就使用默认的
		viper.SetConfigFile(configFile)
		// 读取配置信息
		err := viper.ReadInConfig()
		if err != nil && !os.IsNotExist(err) {
			// 文件不存在会报此错误：open conf/config.yaml: The system cannot find the file specified.
			return err
		}
	}

	// 绑定cmd参数
	_ = viper.BindPFlags(pflag.CommandLine)

	// 将配置信息存储到全局变量中
	err := viper.Unmarshal(&common.Config)
	if err != nil {
		return err
	}
	return nil
}
