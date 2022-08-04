package main

import (
	"fmt"
	"os"

	"github.com/ShangRui-hashy/send-to-burp/config"
	"github.com/sirupsen/logrus"
	cli "github.com/urfave/cli/v2"
)

var configFile string

func main() {
	author := cli.Author{
		Name:  "无在无不在",
		Email: "2227627947@qq.com",
	}
	app := &cli.App{
		Name:      "Send-To-Burp",
		Usage:     "A tool to send url to BurpScanner",
		UsageText: "Send-To-Burp",
		Version:   "v0.1",
		Authors:   []*cli.Author{&author},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "debug",
				Aliases:     []string{"d"},
				Usage:       "debug mode",
				Value:       false,
				Destination: &config.CurrentConf.Debug,
			},
		},
		Action: run,
	}
	//启动app
	if err := app.Run(os.Args); err != nil {
		logrus.Error(err)
	}
}

func run(c *cli.Context) (err error) {
	fmt.Println("Send-To-Burp")
	return nil
}
