package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	"kluivert-queiroz/api/pkg/challenge"
	"kluivert-queiroz/api/pkg/runner"
	"kluivert-queiroz/api/pkg/submission"
)

var validate = validator.New()

func CreateSubmission(runner runner.Runner) fiber.Handler {
	return func(c *fiber.Ctx) error {
		submission := submission.Submission{}
		if err := c.BodyParser(&submission); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		if err := validate.Struct(submission); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		challengeRepository := challenge.NewInMemoryChallengeRepository()
		challenge, err := challengeRepository.GetChallenge(submission.ChallengeID)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		output, err := runner.Assert(submission.Code, submission.Language, challenge)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(output)
	}
}
