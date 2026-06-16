package middleware

import (
	"time"
	"user-api/internal/logger"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func ZapLogger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		// Continue to the next middleware/handler
		err := c.Next()

		// Grab the Request ID injected by Fiber (we will set this up in main.go)
		reqID := c.Response().Header.Peek("X-Request-Id")

		logger.Log.Info("HTTP Request",
			zap.String("request_id", string(reqID)),
			zap.String("method", c.Method()),
			zap.String("path", c.Path()),
			zap.Int("status", c.Response().StatusCode()),
			zap.Duration("latency", time.Since(start)),
		)

		return err
	}
}
