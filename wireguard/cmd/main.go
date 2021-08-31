package main

import (
	"github.com/allocamelus/allocamelus/pkg/logger"
	"github.com/jdinabox/alpine-dockerfiles/wireguard"
	await "github.com/jdinabox/go-await"
)

func main() {
	logger.InitKlog(5, "")
	ai := await.NewInterrupt()

	// Add 1 to wait group
	ai.Add(1)
	go wireguard.Wireguard(ai)

	ai.Wait()
}
