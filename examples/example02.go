package main

import (
	"time"

	"github.com/lei006/zlog"
)

func main() {
	//zlog.SetSaveFile("logs.log", true)
	zlog.SetSaveFile("logs.log", true)

	for {
		zlog.Debug("hello world")

		time.Sleep(1 * time.Second)
	}
}
