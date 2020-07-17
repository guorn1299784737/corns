package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
)

func main() {
	var language string

	//实例化cli
	app := cli.NewApp()
	//Name可以设定应用的名字
	app.Name = "corn"
	// Version可以设定应用的版本号
	app.Version = "1.0.0"
	// Commands用于创建命令
	app.Commands = []cli.Command{
		{
			// 命令的名字
			Name:    "start",
			// 命令的缩写，就是不输入language只输入lang也可以调用命令
			Aliases: []string{"corn"},
			// 命令的用法注释，这里会在输入 程序名 -help的时候显示命令的使用方法
			Usage:   "定时任务",
			Flags: []cli.Flag {
				cli.StringFlag {
					Name: "timeout, t",
					Value: "3600",
					Usage: "read from `FILE`",
					Destination: &language,
				},
			},
			// 命令的处理函数
			Action: startCorn,
		},
	}
	// 接受os.Args启动程序
	app.Run(os.Args)
}

func startCorn(c *cli.Context) error {
	language := c.Args().First()
	if language == "chinese"{
		fmt.Println("Language is 中文")
	}else {
		fmt.Println("Language is English")
	}

	return nil
}

func restartCorn()  {
	
}

func stopCorn()  {
	
}
