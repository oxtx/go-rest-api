package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/oxtx/go-rest-api/internal/dto"
	"github.com/oxtx/go-rest-api/internal/service"
	"github.com/oxtx/go-rest-api/pkg/response"
)

type UserHandler struct {
	svc       service.UserService
	validator *validator.Validate
}

func NewUserHandler(s service.UserService) *UserHandler {
	return &UserHandler{svc: s, validator: validator.New()}
}

// @Summary Create user
// @Tags users
// @Accept json
// @Produce json
// @Param body body dto.CreateUserRequest true "Create User"
// @Success 201 {object} dto.UserResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /api/v1/users [post]
func (h *UserHandler) Create(c *gin.Context) {
	var req dto.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, http.StatusBadRequest, "invalid_body", err.Error())
		return
	}
	if err := h.validator.Struct(req); err != nil {
		response.Fail(c, http.StatusBadRequest, "validation_failed", err.Error())
		return
	}
	u, err := h.svc.CreateUser(c.Request.Context(), req.Email, req.Name)
	if err != nil {
		response.Fail(c, http.StatusInternalServerError, "create_failed", err.Error())
		return
	}
	response.JSON(c, http.StatusCreated, dto.MapUserToResponse(u))
}

// @Summary Get user
// @Tags users
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} dto.UserResponse
// @Failure 404 {object} response.ErrorResponse
// @Router /api/v1/users/{id} [get]
func (h *UserHandler) Get(c *gin.Context) {
	id := c.Param("id")
	u, err := h.svc.GetUser(c.Request.Context(), id)
	if err != nil {
		response.Fail(c, http.StatusNotFound, "not_found", err.Error())
		return
	}
	response.JSON(c, http.StatusOK, dto.MapUserToResponse(u))
}
