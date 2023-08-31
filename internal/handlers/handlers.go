package handlers

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
)

type repository interface {
	Get(id int) ([]byte, error)
}

type Handlers struct {
	repository
}

func New(repository repository) *Handlers {
	return &Handlers{repository}
}
func (h *Handlers) Get(c *fiber.Ctx) error {

	// Read the param noteId
	id := c.Params("Id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.SendStatus(http.StatusNotFound)
	}
	res, err := h.repository.Get(idInt)

	return c.Send(res)
	// Find the note with the given id
	// Return the note with the id
}
