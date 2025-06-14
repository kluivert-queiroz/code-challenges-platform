package main

import (
	"context"
	"kluivert-queiroz/api/internal/handlers"
	"kluivert-queiroz/api/pkg/challenge"
	"kluivert-queiroz/api/pkg/runner"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	runnerService := runner.NewRunnerService()

	app.Get("/challenges", handlers.GetChallenges(challenge.NewInMemoryChallengeRepository()))
	app.Post("/submissions", handlers.CreateSubmission(runnerService))

	ctx, _ := signal.NotifyContext(context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGKILL,
		syscall.SIGQUIT,
	)

	go func() {
		if err := app.Listen(":3000"); err != nil {
			log.Panic(err)
		}
	}()

	<-ctx.Done()

	log.Info("Gracefully shutting down...")

	log.Info("Running cleanup tasks...")
	if err := runnerService.RemoveContainers(context.TODO()); err != nil {
		log.Fatalf("failed to remove containers: %v", err)
	}

}
