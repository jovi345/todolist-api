package handler

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/todos-api/jovi345/formatter"
	"github.com/todos-api/jovi345/token"
	"github.com/todos-api/jovi345/user"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.UserRegistrationInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		msg := formatter.SendResponse("Failed", err.Error())
		c.JSON(http.StatusBadRequest, msg)
		return
	}

	response, err := h.userService.RegisterUser(input)
	if err != nil {
		msg := formatter.SendResponse("Failed", err.Error())
		c.JSON(http.StatusBadRequest, msg)
		return
	}

	msg := formatter.SendResponse("Success", response)
	c.JSON(http.StatusCreated, msg)
}

func (h *userHandler) LoginUser(c *gin.Context) {
	var input user.UserLoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		msg := formatter.SendResponse("Failed", err.Error())
		c.JSON(http.StatusBadRequest, msg)
		return
	}

	_, err = h.userService.Login(input)
	if err != nil {
		msg := formatter.SendResponse("Failed", err.Error())
		c.JSON(http.StatusBadRequest, msg)
		return
	}

	accessToken, err := token.GenerateAccessToken(input.Email)
	if err != nil {
		msg := formatter.SendResponse("Failed", err.Error())
		c.JSON(http.StatusBadRequest, msg)
		return
	}

	refreshToken, err := token.GenerateRefreshToken(input.Email)
	if err != nil {
		msg := formatter.SendResponse("Failed", err.Error())
		c.JSON(http.StatusBadRequest, msg)
		return
	}

	c.SetCookie("refresh_token", refreshToken, int(7*24*3600), "/", "", false, true)

	msg := formatter.SendResponse("Success", accessToken)
	c.JSON(http.StatusOK, msg)
}

func (h *userHandler) RefreshToken(c *gin.Context) {
	cookie, err := c.Request.Cookie("refresh_token")
	if err != nil {
		c.SetCookie("refresh_token", "", -1, "/", "", false, true)
		msg := formatter.SendResponse("Failed", err.Error())
		c.JSON(http.StatusBadRequest, msg)
		return
	}

	secretKey := os.Getenv("JWT_REFRESH_KEY")
	refreshToken := cookie.Value

	_, err = token.ValidateToken(refreshToken, secretKey)
	if err != nil {
		msg := formatter.SendResponse("Failed", err.Error())
		c.JSON(http.StatusForbidden, msg)
		return
	}

	accessToken, err := h.userService.RefreshToken(refreshToken)
	if err != nil {
		msg := formatter.SendResponse("Failed", err.Error())
		c.JSON(http.StatusBadRequest, msg)
		return
	}

	msg := formatter.SendResponse("Success", accessToken)
	c.JSON(http.StatusOK, msg)
}

func (h *userHandler) Logout(c *gin.Context) {
	c.SetCookie("refresh_token", "", -1, "/", "", false, true)

	c.JSON(http.StatusBadRequest, "Logged out")
}
