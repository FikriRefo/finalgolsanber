package handler

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"socmed/dto"
	"socmed/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ProfileHandler struct {
	service service.ProfileService
}

func NewProfileHandler(s service.ProfileService) *ProfileHandler {
	return &ProfileHandler{service: s}
}

func (h *ProfileHandler) Create(c *gin.Context) {
	var req dto.ProfileRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if Avatar is provided
	if req.Avatar != nil {
		// Create the profile directory if it doesn't exist
		if err := os.MkdirAll("public/profile", 0755); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create directory: " + err.Error()})
			return
		}

		// Get the file extension and generate a new file name
		ext := filepath.Ext(req.Avatar.Filename)
		newFileName := uuid.New().String() + ext

		// Define the destination path for saving the file
		dst := filepath.Join("public/profile", filepath.Base(newFileName))

		// Save the uploaded file
		if err := c.SaveUploadedFile(req.Avatar, dst); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file: " + err.Error()})
			return
		}

		// Set the Avatar URL in the request object
		req.AvatarUrl = fmt.Sprintf("%s/public/profile/%s", c.Request.Host, newFileName)
	}

	// Get the user ID from the context and assign it to the request
	userID, _ := c.Get("userId")
	req.UserID = userID.(int)

	// Call the service to create the profile
	if err := h.service.Create(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Profile created"})
}

func (h *ProfileHandler) GetByUserID(c *gin.Context) {
	// Get the user ID from the context
	userID, _ := c.Get("userId")

	// Fetch the profile using the service layer
	profile, err := h.service.GetByUserID(userID.(int))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Profile not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"profile": profile})
}

func (h *ProfileHandler) Update(c *gin.Context) {
	var req dto.ProfileRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get the profile ID from the URL parameter
	profileID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid profile ID"})
		return
	}

	// Check if Avatar is provided
	if req.Avatar != nil {
		// Create the profile directory if it doesn't exist
		if err := os.MkdirAll("public/profile", 0755); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create directory: " + err.Error()})
			return
		}

		// Get the file extension and generate a new file name
		ext := filepath.Ext(req.Avatar.Filename)
		newFileName := uuid.New().String() + ext

		// Define the destination path for saving the file
		dst := filepath.Join("public/profile", filepath.Base(newFileName))

		// Save the uploaded file
		if err := c.SaveUploadedFile(req.Avatar, dst); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file: " + err.Error()})
			return
		}

		// Set the Avatar URL in the request object
		req.AvatarUrl = fmt.Sprintf("%s/public/profile/%s", c.Request.Host, newFileName)
	}

	// Get the user ID from the context
	userID, _ := c.Get("userId")
	req.UserID = userID.(int)

	// Call the service to update the profile
	if err := h.service.Update(profileID, &req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile updated"})
}

func (h *ProfileHandler) Delete(c *gin.Context) {
	// Get the profile ID from the URL parameter
	profileID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid profile ID"})
		return
	}

	// Call the service to delete the profile
	if err := h.service.Delete(profileID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile deleted"})
}
