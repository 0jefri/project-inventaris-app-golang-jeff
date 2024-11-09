package controller

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/project-app-inventaris/internal/app/service"
	"github.com/project-app-inventaris/internal/model"
	"github.com/project-app-inventaris/internal/model/dto"
	"github.com/project-app-inventaris/utils/common"
	"github.com/project-app-inventaris/utils/exception"
)

type UserController struct {
	service service.UserService
	auth    service.AuthService
}

func NewUserController(service service.UserService, auth service.AuthService) *UserController {
	return &UserController{
		service: service,
		auth:    auth,
	}
}

func (ctr *UserController) Registration(c *gin.Context) {
	payload := model.User{}

	payload.ID = common.GenerateUUID()
	payload.RegistrationDate = time.Now()
	payload.LastLogin = time.Now()

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"status":  exception.StatusBadRequest,
			"message": exception.FieldErrors(err),
		})
		return
	}

	_, err := ctr.service.RegisterNewUser(&payload)

	if err != nil {
		if errors.Is(err, exception.ErrFailedCreate) {
			c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Status:  exception.StatusInternalServer,
				Message: exception.ErrFailedCreate.Error(),
			})
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Status:  exception.StatusInternalServer,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, dto.Response{
		Code:    http.StatusCreated,
		Status:  exception.StatusSuccess,
		Message: "Register Successful",
	})
}

func (ctr *UserController) Login(c *gin.Context) {
	payload := dto.Auth{}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"status":  exception.StatusBadRequest,
			"message": exception.FieldErrors(err),
		})
		return
	}

	data, err := ctr.auth.Login(payload.Username, payload.Password)

	if err != nil {
		if errors.Is(err, exception.ErrInvalidParseToken) {
			c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Status:  exception.StatusInternalServer,
				Message: exception.ErrInvalidParseToken.Error(),
			})
			return
		}

		if errors.Is(err, exception.ErrInvalidTokenStringMethod) {
			c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Status:  exception.StatusInternalServer,
				Message: exception.ErrInvalidTokenStringMethod.Error(),
			})
			return
		}

		if errors.Is(err, exception.ErrInvalidTokenMapclaims) {
			c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Status:  exception.StatusInternalServer,
				Message: exception.ErrInvalidTokenMapclaims.Error(),
			})
			return
		}

		if errors.Is(err, exception.ErrFailedCreateToken) {
			c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Status:  exception.StatusInternalServer,
				Message: exception.ErrFailedCreateToken.Error(),
			})
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Status:  exception.StatusInternalServer,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.TokenResponse{
		Code:    http.StatusOK,
		Status:  exception.StatusSuccess,
		Message: "Login Successful",
		Token:   data,
	})
}
