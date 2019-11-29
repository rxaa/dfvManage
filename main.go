package main

import (
	"dfv/model"
	"dfv/utils"
	"fmt"
	"os"
)

func rocoverErr() {
	if err := recover(); err != nil {
		fmt.Println(err)
	}
}

func main() {

	if len(os.Args) > 1 {
		parseArg()
		return;
	}

	fmt.Println("list (show process list)")
	fmt.Println("start xx (start process)")
	fmt.Println("stop xx (stop process)")
	fmt.Println("restart xx (restart process)")
	fmt.Println("install (install to startup)")
	fmt.Println("start_watch (start dfv watch process)")
	fmt.Println("stop_watch (stop dfv watch process)")
	fmt.Println("restart_watch (restart dfv watch process)")
}

func parseArg() {
	switch os.Args[1] {
	case "install":
		install()
	case "start":
		start()
	default:
		fmt.Println("unknown command")
	}

}

const cfgFile = "/usr/local/dfv.cfg"

func start() {
	if len(os.Args) < 3 {
		fmt.Println("Please input start menu")
		return;
	}

	//menu := os.Args[2]
	//dir := path.Dir(menu)
	//name := path.Base(menu)
	var cfg model.Cfg
	utils.JsonDecode(utils.ReadFile(cfgFile), &cfg);

}

func systemctlMenu() string {
	//centos
	if utils.Exists("/usr/lib/systemd/system") {
		return "/lib/systemd/system"
	}

	//debian
	if utils.Exists("/lib/systemd/system") {
		return "/lib/systemd/system"
	}

	panic("unsupport systemd")
}

const dfvService = `[Unit]
Description=dfvService

# 描述服务类别，表示本服务需要在network服务启动后在启动
After=network.target

[Service]

Type=forking

Restart=no

ExecStart=dfv start_watch
ExecReload=dfv restart_watch
ExecStop=dfv stop_watch

PrivateTmp=true


[Install]
WantedBy=multi-user.target
`
const serviceFile = "/dfv.service"

func install() {
	if utils.Exists(systemctlMenu() + serviceFile) {
		fmt.Println("dfv already installed")
		return
	}

	utils.WriteFile(systemctlMenu()+serviceFile, []byte(dfvService), 0644)

	utils.Exec("systemctl start dfv.service")
	utils.Exec("systemctl enable softether.service")
	utils.Exec("systemctl status softether.service")
}
