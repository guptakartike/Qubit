package handler

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/guptakartike/qubit/internal/auth"
)

type RegistrationService interface {
	Register(
		ctx context.Context,
		req auth.RegisterRequest,
	) (auth.User, error)
}

type RegistrationHandler struct {
	service RegistrationService
}

func NewRegistrationHandler(
	service RegistrationService,
) *RegistrationHandler {
	return &RegistrationHandler{
		service: service,
	}
}

// RegisterRoutes wires the handler's routes onto the provided router.
// This satisfies the server.RouteRegistrar interface.
func (h *RegistrationHandler) RegisterRoutes(router gin.IRouter) {
	router.POST("/api/v1/auth/register", h.Register)
}

func (h *RegistrationHandler) Register(c *gin.Context) {
	var req auth.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	user, err := h.service.Register(c.Request.Context(), req)
	if err != nil {
		switch {
		case errors.Is(err, auth.ErrInvalidInput):
			var ve *auth.ValidationError
			if errors.As(err, &ve) {
				c.JSON(http.StatusBadRequest, gin.H{
					"error":   "validation_error",
					"field":   ve.Field,
					"message": ve.Message,
				})
			} else {
				c.JSON(http.StatusBadRequest, gin.H{
					"error":   "validation_error",
					"message": err.Error(),
				})
			}

		case errors.Is(err, auth.ErrEmailAlreadyExists):
			c.JSON(http.StatusConflict, gin.H{
				"error": "email already exists",
			})

		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "internal server error",
			})
		}

		return
	}

	response := auth.RegisterResponse{
		ID:     user.ID.String(),
		Name:   user.Name,
		Email:  user.Email,
		Status: user.Status,
	}

	c.JSON(http.StatusCreated, response)
}
