package controllers

import (
	"net/http"
	"order-manage/models"

	"github.com/gin-gonic/gin"
)

// GetProfile godoc
// @Summary Get Profile User
// @Description Retrieve profile user
// @Tags User
// @Security BearerAuth
// @Param Authorization header string true "Bearer Token"
// @Accept json
// @Produce json
// @Success 200 {array} models.User
// @Router /user/profile [get]
func (h *Handler) GetProfile(c *gin.Context) {
	userID, exist := c.Get("user_id")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user id not found"})
		return
	}

	id := uint64(userID.(float64))

	user, err := h.UserUsecase.GetProfile(c, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	c.JSON(http.StatusOK, models.Success{
		Status: http.StatusOK,
		Data:   user,
	})
}
