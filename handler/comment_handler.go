package handler

import (
	"net/http"
	"socmed/dto"
	"socmed/helper"
	"socmed/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentHandler struct {
	service service.CommentService
}

func NewCommentHandler(s service.CommentService) *CommentHandler {
	return &CommentHandler{
		service: s,
	}
}

func (h *CommentHandler) Create(c *gin.Context) {
	var req dto.CommentRequest
	// Bind the incoming JSON request to the req variable
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get user ID from the context (assuming it was set during authentication)
	userID, _ := c.Get("userId")
	req.UserID = userID.(int)

	// Call service layer to create the comment
	if err := h.service.Create(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Create response
	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusCreated,
		Message:    "Berhasil Memberi Komentar",
		Data:       req, // Use the req object which now has UserID set
	})

	c.JSON(http.StatusCreated, res)
}

func (h *CommentHandler) GetAllByPostID(c *gin.Context) {
	postIDStr := c.Param("post_id")        // Get post_id as string
	postID, err := strconv.Atoi(postIDStr) // Convert string to int
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	comments, err := h.service.GetAllByPostID(postID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, comments)
}

func (h *CommentHandler) Update(c *gin.Context) {
	commentIDStr := c.Param("id")                // Get comment ID as string
	commentID, err := strconv.Atoi(commentIDStr) // Convert string to int
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid comment ID"})
		return
	}

	var req dto.CommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.Update(commentID, &req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment updated successfully"})
}

func (h *CommentHandler) Delete(c *gin.Context) {
	commentIDStr := c.Param("id")                // Get comment ID as string
	commentID, err := strconv.Atoi(commentIDStr) // Convert string to int
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid comment ID"})
		return
	}

	if err := h.service.Delete(commentID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment deleted successfully"})
}
