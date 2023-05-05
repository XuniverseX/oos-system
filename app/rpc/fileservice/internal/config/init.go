package config

import "os"

func Init(config Config) {
	//writeToPath := config.WriteToPath
	//如果目录不存在 就创建
	//if _, err := os.Stat(writeToPath); os.IsNotExist(err) {
	//	os.MkdirAll(writeToPath, os.ModePerm)
	//}
	tempPath := config.TempPath
	//如果目录不存在 就创建
	if _, err := os.Stat(tempPath); os.IsNotExist(err) {
		os.MkdirAll(tempPath, os.ModePerm)
	}
}
