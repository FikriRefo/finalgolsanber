package handler

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"socmed/dto"
	"socmed/errorhandle"
	"socmed/helper"
	"socmed/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type postHandler struct {
	service service.PostService
}

func NewPostHandler(s service.PostService) *postHandler {
	return &postHandler{
		service: s,
	}
}

func (h *postHandler) Create(c *gin.Context) {
	var post dto.PostRequest

	// Mengikat data dari form-data
	if err := c.ShouldBind(&post); err != nil {
		errorhandle.HandleError(c, &errorhandle.BadRequestError{Message: err.Error()})
		return
	}

	// Mengambil file gambar dari form-data jika ada
	if post.Picture != nil {
		if err := os.MkdirAll("public/picture", 0755); err != nil {
			errorhandle.HandleError(c, &errorhandle.InternalServerError{Message: err.Error()})
			return
		}

		ext := filepath.Ext(post.Picture.Filename)
		newFileName := uuid.New().String() + ext

		// save image
		dst := filepath.Join("public/picture", filepath.Base(newFileName))
		c.SaveUploadedFile(post.Picture, dst)

		post.Picture.Filename = fmt.Sprintf("%s/public/picture/%s", c.Request.Host, newFileName)
	}

	userID, _ := c.Get("userId")
	post.UserID = userID.(int)

	if err := h.service.Create(&post); err != nil {
		errorhandle.HandleError(c, err)
		return
	}

	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusCreated,
		Message:    "Berhasil Membuat Post",
	})

	c.JSON(http.StatusCreated, res)
}

func (h *postHandler) GetAllByUserID(c *gin.Context) {
	userID, _ := c.Get("userId")
	posts, err := h.service.GetAllByUserID(userID.(int))
	if err != nil {
		errorhandle.HandleError(c, err)
		return
	}

	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "Berhasil mendapatkan data",
		Data:       posts,
	})

	c.JSON(http.StatusOK, res)
}

func (h *postHandler) Update(c *gin.Context) {
	var post dto.PostRequest
	postID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorhandle.HandleError(c, &errorhandle.BadRequestError{Message: "Invalid post ID"})
		return
	}

	if err := c.ShouldBind(&post); err != nil {
		errorhandle.HandleError(c, &errorhandle.BadRequestError{Message: err.Error()})
		return
	}

	if err := h.service.Update(postID, &post); err != nil {

		errorhandle.HandleError(c, err)
		return
	}

	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "Berhasil memperbarui post",
	})

	c.JSON(http.StatusOK, res)
}

func (h *postHandler) Delete(c *gin.Context) {
	postID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorhandle.HandleError(c, &errorhandle.BadRequestError{Message: "Invalid post ID"})
		return
	}

	userID, _ := c.Get("userId")

	if err := h.service.Delete(postID, userID.(int)); err != nil {
		errorhandle.HandleError(c, err)
		return
	}

	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "Berhasil menghapus post",
	})

	c.JSON(http.StatusOK, res)
}
