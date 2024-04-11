package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string

	DbIp       string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string

	RdsIp       string
	RdsPort     string
	RdsNo       string
	RdsUser     string
	RdsPassword string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径", err)
	}
	LoadServer(file)
	LoadDb(file)
	LoadCache(file)
}
func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3001")
}

func LoadDb(file *ini.File) {
	DbIp = file.Section("mysql").Key("DbHost").MustString("127.0.0.1")
	DbPort = file.Section("mysql").Key("DbPort").MustString("3306")
	DbUser = file.Section("mysql").Key("DbUser").MustString("root")
	DbPassword = file.Section("mysql").Key("DbPassWord").MustString("root")
	DbName = file.Section("mysql").Key("DbName").MustString("ginblogm")
}

func LoadCache(file *ini.File) {
	RdsIp = file.Section("redis").Key("RdsIp").String()
	RdsPort = file.Section("redis").Key("RdsPort").String()
	RdsNo = file.Section("redis").Key("RdsNo").String()
	RdsUser = file.Section("redis").Key("RdsUser").String()
	RdsPassword = file.Section("redis").Key("RdsPassword").String()
}
