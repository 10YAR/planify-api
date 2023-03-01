package utils

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func SetLocal[T any](c *fiber.Ctx, key string, value T) {
	c.Locals(key, value)
}
func GetLocal[T any](c *fiber.Ctx, key string) T {
	fmt.Println(c.Locals(key).(T))
	return c.Locals(key).(T)
}
