package handlers

import (
	"encoding/csv"
	"evconn/internal/core/domain/models"
	"evconn/internal/core/ports"
	"evconn/internal/pkg/response"
	"evconn/internal/pkg/utils"
	"io"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService ports.UserService
}

func NewUserHandler(userService ports.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) GetAll(c *gin.Context) {
	// Extract pagination parameters
	page, pageSize := utils.GetPaginationParams(c)

	users, total, err := h.userService.GetAllUsersPaginated(c.Request.Context(), page, pageSize)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.SuccessWithPagination(c, users, total, page, pageSize)
}

func (h *UserHandler) Import(c *gin.Context) {
	// Get file from form
	file, err := c.FormFile("file")
	if err != nil {
		response.Error(c, 400, "File upload error: "+err.Error())
		return
	}

	// Open the file
	src, err := file.Open()
	if err != nil {
		response.Error(c, 500, "Failed to open file: "+err.Error())
		return
	}
	defer src.Close()

	// Read CSV file
	csvReader := csv.NewReader(src)

	// Skip header row
	_, err = csvReader.Read()
	if err != nil {
		response.Error(c, 400, "Failed to read CSV header: "+err.Error())
		return
	}

	// Process user data rows
	var users []*models.User
	var errors []string

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			errors = append(errors, "Failed to read row: "+err.Error())
			continue
		}

		// Validate row has enough columns
		if len(record) < 5 {
			errors = append(errors, "Row has insufficient columns")
			continue
		}

		// Create user from CSV data
		user := &models.User{
			NIM:      record[0],
			Name:     record[1],
			Email:    record[2],
			Password: record[3], // Will be hashed by service
			Role:     record[4],
			Lab:      record[5],
		}

		if len(record) > 6 {
			shift := record[6]
			user.Shift = &shift
		}

		if len(record) > 7 {
			hari := record[7]
			user.Hari = &hari
		}

		if len(record) > 8 {
			kodeAsisten := record[8]
			user.KodeAsisten = &kodeAsisten
		}

		users = append(users, user)
	}

	// Import users
	importErrors, err := h.userService.ImportUsers(c.Request.Context(), users)
	if err != nil {
		response.Error(c, 500, "Import failed: "+err.Error())
		return
	}

	// Combine all errors
	errors = append(errors, importErrors...)

	response.Success(c, gin.H{
		"imported": len(users) - len(errors),
		"total":    len(users),
		"errors":   errors,
	})
}
