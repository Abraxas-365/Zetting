package workRequest_handlers

import (
	ports "zetting/pkg/workRequest/core/ports"
)

import "github.com/gofiber/fiber/v2"

type WorkRequestHandler interface {
	CreateWorkRequest(c *fiber.Ctx) error
}
type workRequestHandler struct {
	userService ports.WorkRequestService
}

func NewWorkRequestHandler(workRequestService ports.WorkRequestService) WorkRequestHandler {
	return &workRequestHandler{
		workRequestService,
	}
}
