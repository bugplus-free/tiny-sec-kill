package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	DbIp       string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径", err)
	}
	LoadDb(file)
}

func LoadDb(file *ini.File) {
	DbIp = file.Section("mysql").Key("DbHost").MustString("127.0.0.1")
	DbPort = file.Section("mysql").Key("DbPort").MustString("3306")
	DbUser = file.Section("mysql").Key("DbUser").MustString("root")
	DbPassword = file.Section("mysql").Key("DbPassWord").MustString("root")
	DbName = file.Section("mysql").Key("DbName").MustString("ginblogm")
}
