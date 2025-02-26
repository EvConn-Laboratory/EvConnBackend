package handlers

import (
	"evconn/internal/core/domain/models"
	"evconn/internal/core/ports"
	"evconn/internal/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type LabHandler struct {
	labService ports.LabService
}

func NewLabHandler(labService ports.LabService) *LabHandler {
	return &LabHandler{labService: labService}
}

func (h *LabHandler) GetAll(c *gin.Context) {
	labs, err := h.labService.GetLabs(c.Request.Context())
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}
	response.Success(c, labs)
}

func (h *LabHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.Error(c, 400, "Invalid ID")
		return
	}

	lab, err := h.labService.GetLabByID(c.Request.Context(), uint(id))
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, lab)
}

func (h *LabHandler) Create(c *gin.Context) {
	var lab models.Lab
	if err := c.ShouldBindJSON(&lab); err != nil {
		response.Error(c, 400, err.Error())
		return
	}

	if err := h.labService.CreateLab(c.Request.Context(), &lab); err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, lab)
}

func (h *LabHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.Error(c, 400, "Invalid ID")
		return
	}

	var lab models.Lab
	if err := c.ShouldBindJSON(&lab); err != nil {
		response.Error(c, 400, err.Error())
		return
	}

	lab.ID = uint(id)
	if err := h.labService.UpdateLab(c.Request.Context(), &lab); err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, lab)
}

func (h *LabHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.Error(c, 400, "Invalid ID")
		return
	}

	if err := h.labService.DeleteLab(c.Request.Context(), uint(id)); err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, nil)
}
