package main

import (
	"os"
	"strconv"

	"github.com/allocamelus/allocamelus/pkg/logger"
	"github.com/jdinabox/alpine-dockerfiles/wireguard"
	await "github.com/jdinabox/go-await"
	toolserver "github.com/jdinabox/tool-server"
)

func main() {
	logV, _ := strconv.ParseInt(os.Getenv("LOG_V"), 10, 8)
	logger.InitKlog(int8(logV), false, "")
	ai := await.NewInterrupt()

	// Add 1 to wait group
	ai.Add(1)
	go wireguard.Wireguard(ai)
	ai.Add(1)
	go toolserver.StartAwaitInterupt(&toolserver.Config{Listen: ":80"}, ai)

	ai.Wait()
}
