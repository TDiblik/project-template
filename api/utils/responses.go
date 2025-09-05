package utils

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
)

func InvalidRequestResponse(c fiber.Ctx, e error) error {
	LogErr(e)
	if EnvData.Debug {
		return c.Status(400).JSON(fiber.Map{"status": "error", "msg": "msg.general.invalid_request", "message": "Review your input", "error_info": fmt.Sprint(e)})
	}
	return c.Status(400).JSON(fiber.Map{"status": "error", "msg": "msg.general.invalid_request"})
}

func InternalServerErrorResponse(c fiber.Ctx, e error) error {
	LogErr(e)
	if EnvData.Debug {
		return c.Status(500).JSON(fiber.Map{"status": "error", "msg": "msg.general.internal_server_error", "message": "Internal Server Error", "error_info": fmt.Sprint(e)})
	}
	return c.Status(500).JSON(fiber.Map{"status": "error", "msg": "msg.general.internal_server_error", "message": "Internal Server Error"})
}

func UnauthentizatedResponse(c fiber.Ctx, e error) error {
	LogErr(e)
	if EnvData.Debug {
		return c.Status(401).JSON(fiber.Map{"status": "unauthenticated", "msg": "msg.general.invalid_token", "message": "valid token is required", "error_info": fmt.Sprint(e)})
	}
	return c.Status(401).JSON(fiber.Map{"status": "unauthenticated", "msg": "msg.general.invalid_token"})
}

func UnauthorizedResponse(c fiber.Ctx, e error) error {
	LogErr(e)
	if EnvData.Debug {
		return c.Status(403).JSON(fiber.Map{"status": "unauthorized", "msg": "msg.general.unatuhorized", "message": "you cannot access this endpoint", "error_info": fmt.Sprint(e)})
	}
	return c.Status(403).JSON(fiber.Map{"status": "unauthorized", "msg": "msg.general.unatuhorized"})
}

func ConflictResponse(c fiber.Ctx, msg_reason string) error {
	Log("conflict: ", msg_reason)
	return c.Status(409).JSON(fiber.Map{"status": "conflict", "msg": msg_reason})
}

func NotFoundResponse(c fiber.Ctx, msg_reason string) error {
	return c.Status(404).JSON(fiber.Map{"status": "not found", "msg": msg_reason})
}

func OkResponse(c fiber.Ctx, data fiber.Map) error {
	data["status"] = "ok"
	return c.Status(200).JSON(data)
}
