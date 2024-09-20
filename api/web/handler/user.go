package handler

import (
	"time"

	"github.com/ei-sugimoto/tatekae/api/model"
	"github.com/ei-sugimoto/tatekae/api/pkg"
	"github.com/ei-sugimoto/tatekae/api/usecase"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserUsecase usecase.IUserUsecase
}

func NewUserHandler(userusecase usecase.IUserUsecase) *UserHandler {
	return &UserHandler{
		UserUsecase: userusecase,
	}
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterResponse struct {
	Token string `json:"token"`
}

func (h *UserHandler) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	newUser := &model.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
		CreateAt: time.Now(),
	}

	res, err := h.UserUsecase.Create(newUser)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	token, err := pkg.NewToken(res)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, RegisterResponse{Token: token})

}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *UserHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, err := h.UserUsecase.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Err})
		return
	}

	token, TokenErr := pkg.NewToken(user)
	if TokenErr != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, RegisterResponse{Token: token})
}
