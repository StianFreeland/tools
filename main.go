package main

import (
	"tools/kernel"
	"tools/services/config"
	"tools/services/zlog"
)

func main() {
	if !config.Init() {
		return
	}
	zlog.Start()
	defer zlog.Stop()
	zlog.Warn("program", "start ...")
	defer zlog.Warn("program", "stop ...")
	kernel.Run()
}
