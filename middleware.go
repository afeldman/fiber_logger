package fiber_logger

import (
	"fmt"
	"net/http"

	"github.com/afeldman/fiber_logger/logger"
	"github.com/gofiber/fiber/v2"
)

type Config struct {
	Next func(ctx *fiber.Ctx) bool
}

func NewLogger(config ...Config) fiber.Handler {

	return func(c *fiber.Ctx) (err error) {

		code := c.Response().StatusCode()

		msg := "Request"
		if err := c.Next(); err != nil {
			msg = err.Error()
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
		}

		msg = fmt.Sprintf("%v, %s", code, msg)

		if code >= fiber.StatusBadRequest && code < fiber.StatusInternalServerError {
			logger.Warn(msg)
		} else if code >= http.StatusInternalServerError {
			logger.Error(msg)
		} else {
			logger.Info(msg)
		}

		return nil
	}
}
