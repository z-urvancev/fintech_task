package handler

import (
	"fintech/internal/usecase"
	"fintech/pkg/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	useCase usecase.UseCase
}

func NewHandler(useCase usecase.UseCase) *Handler {
	return &Handler{useCase: useCase}
}

func (h *Handler) GetURLByShort(c *gin.Context) {
	input := struct {
		short string
	}{}
	if err := c.BindJSON(&input); err != nil {
		_ = c.Error(errors.ErrBadRequest)
		return
	}
	url, getErr := h.useCase.GetURLByShort(input.short)
	if getErr != nil {
		_ = c.Error(getErr)
		return
	}

	c.JSON(http.StatusOK, gin.H{"url": url})
}

func (h *Handler) GenerateShortURL(c *gin.Context) {
	input := struct {
		url string
	}{}
	if err := c.BindJSON(&input); err != nil {
		_ = c.Error(errors.ErrBadRequest)
		return
	}
	short, generateErr := h.useCase.GenerateShortURL(input.url)
	if generateErr != nil {
		_ = c.Error(generateErr)
		return
	}

	c.JSON(http.StatusOK, gin.H{"short": short})
}
