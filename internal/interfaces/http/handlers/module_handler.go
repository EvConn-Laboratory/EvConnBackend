package handlers

import (
	"evconn/internal/core/domain/models"
	"evconn/internal/core/ports"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ModuleHandler struct {
	moduleService ports.ModuleService
}

func NewModuleHandler(moduleService ports.ModuleService) *ModuleHandler {
	return &ModuleHandler{
		moduleService: moduleService,
	}
}

func (h *ModuleHandler) Create(c *gin.Context) {
	var module models.Module
	if err := c.ShouldBindJSON(&module); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := h.moduleService.CreateModule(c.Request.Context(), &module); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, module)
}

func (h *ModuleHandler) GetByID(c *gin.Context) {
	moduleId, err := strconv.ParseUint(c.Param("moduleId"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid module ID"})
		return
	}

	module, err := h.moduleService.GetModule(c.Request.Context(), uint(moduleId))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, module)
}

func (h *ModuleHandler) AssignMentor(c *gin.Context) {
	moduleID, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	mentorID, _ := strconv.ParseUint(c.Param("mentorId"), 10, 32)

	err := h.moduleService.AssignMentor(c.Request.Context(), uint(moduleID), uint(mentorID))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "mentor assigned successfully"})
}
