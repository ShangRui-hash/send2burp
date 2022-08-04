package main

import (
	"bufio"
	"os"

	"github.com/ShangRui-hash/send2burp/burpapi"
	"github.com/ShangRui-hash/send2burp/config"
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
			&cli.StringFlag{
				Name:        "configuration_name",
				Aliases:     []string{"cn"},
				Usage:       "configuration name",
				Value:       "Crawl and Audit - Lightweight",
				Destination: &config.CurrentConf.ConfigurationName,
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
	//从标准输入中读
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		//获取输入的内容
		url := scanner.Text()
		//发送到burp
		if err := burpapi.Scan(url, config.CurrentConf.ConfigurationName); err != nil {
			logrus.Error("burpapi.Scan failed,err:", err)
		}
	}
	return nil
}
