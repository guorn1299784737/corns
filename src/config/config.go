package config

import (
	"fmt"
	"github.com/Unknwon/goconfig"
	log "github.com/sirupsen/logrus"
)

func main() {
	cfg, err := goconfig.LoadConfigFile("./conf.ini")
	if err != nil {
		log.Fatalf("无法加载配置文件：%s", err)
	}
	userListKey, err := cfg.GetValue("", "USER_LIST")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(userListKey)
}
