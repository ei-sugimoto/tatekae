package handler

import (
	"time"

	"github.com/ei-sugimoto/tatekae/api/model"
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
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
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

	c.JSON(201, RegisterResponse{
		ID:       res.ID,
		Username: res.Username,
		Email:    res.Email,
	})

}
