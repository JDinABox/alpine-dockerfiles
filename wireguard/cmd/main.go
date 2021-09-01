package main

import (
	"github.com/allocamelus/allocamelus/pkg/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/jdinabox/alpine-dockerfiles/wireguard"
	await "github.com/jdinabox/go-await"
	jsoniter "github.com/json-iterator/go"
	"k8s.io/klog/v2"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func main() {
	logger.InitKlog(5, "")
	ai := await.NewInterrupt()

	// Add 1 to wait group
	ai.Add(1)
	go wireguard.Wireguard(ai)
	ai.Add(1)
	go server(ai)

	ai.Wait()
}

func server(ai *await.Interrupt) {
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		BodyLimit:   1 * 1024 * 2024, // 1MB
	})
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})
	app.Use(func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"error":     "404",
			"not-found": c.OriginalURL(),
		})
	})

	// Wait for interupt and shutdown
	go func() {
		ai.Await()
		klog.Info("Stopping fiber")
		app.Shutdown()
	}()

	logger.Fatal(app.Listen(":80"))
	ai.Done()
}
