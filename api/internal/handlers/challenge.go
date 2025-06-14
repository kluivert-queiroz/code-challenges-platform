package handlers

import (
	"kluivert-queiroz/api/pkg/challenge"

	"github.com/gofiber/fiber/v2"
)

func GetChallenges(repository challenge.ChallengeRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		challenges, err := repository.GetChallenges()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(challenges)
	}
}
