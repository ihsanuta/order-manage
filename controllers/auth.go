package controllers

import (
	"net/http"
	"order-manage/models"
	"order-manage/usecase"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	AuthUsecase     usecase.AuthUsecase
	UserUsecase     usecase.UserUsecase
	CustomerUsecase usecase.CustomerUsecase
	OrderUsecase    usecase.OrderUsecase
}

// Register godoc
// @Summary Register User
// @Description Register User
// @Tags auth
// @Accept json
// @Produce json
// @Success 200
// @Param order body models.RegisterInput true "Form Register"
// @Router /register [post]
func (h *Handler) Register(c *gin.Context) {

	var input models.RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := &models.User{
		Username: input.Username,
		Password: input.Password,
	}

	err := h.AuthUsecase.Register(c, u)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "registration success"})

}

// Login godoc
// @Summary Login User
// @Description Login User
// @Tags auth
// @Accept json
// @Produce json
// @Param order body models.LoginInput true "Form Register"
// @Success 200
// @Router /login [post]
func (h *Handler) Login(c *gin.Context) {
	var input models.LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.User{
		Username: input.Username,
		Password: input.Password,
	}

	token, err := h.AuthUsecase.LoginCheck(c, u.Username, u.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})

}
