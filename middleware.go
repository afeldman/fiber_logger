package fiber_logger

import (
	"fmt"
	"net/http"

	"github.com/afeldman/fiber_logger/logger"
	"github.com/gofiber/fiber/v2"
)

// middleware configuration
type Config struct {
	Next func(ctx *fiber.Ctx) bool
}

// create a new logger middleware object for fiber framework
func NewLogger(config ...Config) fiber.Handler {

	return func(c *fiber.Ctx) (err error) {

		// get the errorcode object from fiber context
		code := c.Response().StatusCode()

		// set a default string for the logger
		msg := "Request"
		if err := c.Next(); err != nil {
			// translate the error in a nice messqage
			msg = fmt.Sprintf("Error: %s\t", err.Error())
			if e, ok := err.(*fiber.Error); ok {
				// translate the error code into a http error code based on fiber lib
				code = e.Code
			}
		}

		// build a message
		msg = fmt.Sprintf("%v, %s", code, msg)

		// set the logging output based in the error code
		if code >= fiber.StatusBadRequest && code < fiber.StatusInternalServerError {
			logger.Warn(msg)
		} else if code >= http.StatusInternalServerError {
			logger.Error(msg)
		} else {
			logger.Info(msg)
		}

		// if no error continue with rest
		return nil
	}
}
