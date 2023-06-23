package handlers

import (
	"learn-clean-arch/models"
	"learn-clean-arch/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	bookService services.BookService
}

func NewBookHandler(bookService services.BookService) *BookHandler {
	return &BookHandler{bookService}
}

func (h *BookHandler) FindAll(ctx *gin.Context) {
	books, err := h.bookService.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": books,
	})
}

func (h *BookHandler) FindByID(ctx *gin.Context) {
	bookIDParam := ctx.Param("id")
	bookID, err := strconv.Atoi(bookIDParam)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	book, err := h.bookService.FindByID(bookID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}

func (h *BookHandler) Create(ctx *gin.Context) {
	var bookInput models.BookInput

	err := ctx.ShouldBindJSON(&bookInput)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	createdBook, err := h.bookService.Create(bookInput)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"data": createdBook,
	})
}

func (h *BookHandler) Update(ctx *gin.Context) {
	var bookInput models.BookInput

	bookID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctx.ShouldBindJSON(&bookInput)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	updatedBook, err := h.bookService.Update(bookInput, bookID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": updatedBook,
	})
}

func (h *BookHandler) Delete(ctx *gin.Context) {
	bookID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	deletedBook, err := h.bookService.Delete(bookID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": deletedBook,
	})
}
