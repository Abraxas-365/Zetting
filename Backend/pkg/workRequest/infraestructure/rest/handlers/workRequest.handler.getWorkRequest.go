package workRequest_handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h *workRequestHandler) GetWorkRequestsByProject(c *fiber.Ctx) error {

	// userTokenData, err := auth.ExtractTokenMetadata(c)
	// if err != nil {
	// 	return c.Status(500).SendString(err.Error())
	// }
	page, _ := strconv.Atoi(c.Params("page"))
	projectId := c.Params("project_id")
	projects, err := h.workRequestService.GetWorkRequests(projectId, page, "project_id")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(projects)
}

func (h *workRequestHandler) GetWorkRequestsByWorker(c *fiber.Ctx) error {

	// userTokenData, err := auth.ExtractTokenMetadata(c)
	// if err != nil {
	// 	return c.Status(500).SendString(err.Error())
	// }
	page, _ := strconv.Atoi(c.Params("page"))
	workerId := c.Params("worker_id")
	projects, err := h.workRequestService.GetWorkRequests(workerId, page, "worker_id")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(projects)
}
