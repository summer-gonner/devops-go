package framework

import (
	"devops-go/basicdata/common/global"
	"devops-go/basicdata/config"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

const ApplicationYamlFilePath = "resources/application.yaml"

func InitApplicationYaml() {
	dataBytes, err := os.ReadFile(ApplicationYamlFilePath)
	if err != nil {
		fmt.Println("读取文件失败：", err)
		return
	}
	//fmt.Println("yaml 文件的内容: \n", string(dataBytes))
	c := config.Application{}
	err = yaml.Unmarshal(dataBytes, &c)
	if err != nil {
		fmt.Println("解析 yaml 文件失败：", err)
		return
	}
	//fmt.Printf("c → %+v\n", c) // c → {Mysql:{Url:127.0.0.1 Port:3306} Redis:{Host:127.0.0.1 Port:6379}}
	global.Application = &c // 存入全局变量
}
