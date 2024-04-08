package main

import (
	"git.zx-tech.net/app/ait-go-app/application"
	"git.zx-tech.net/app/ait-go-app/config"
	"git.zx-tech.net/app/ait-go-app/controller"
	"git.zx-tech.net/app/ait-go-app/service/runpodapi"
	"github.com/spf13/pflag"
	"go.uber.org/fx"
)

func main() {
	setFlags()
	app := fx.New(
		config.Module,
		controller.Module,
		runpodapi.Module,
		application.Module,
	)
	app.Run()
}

func setFlags() {
	pflag.StringVarP(&config.CfgPath, "config", "c", "config/config.yml", "配置文件路径")
	pflag.Parse()
}
