package controllers

import (
	"net/http"
	"order-manage/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// SaveOrder godoc
// @Summary Save Order
// @Description Insert Order Data
// @Tags Order
// @Security BearerAuth
// @Param Authorization header string true "Bearer Token"
// @Accept json
// @Produce json
// @Param order body models.Order true "Form Order"
// @Success 200 {object} models.Success "Successful response"
// @Router /order [post]
func (h *Handler) SaveOrder(c *gin.Context) {
	var input *models.Order

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.OrderUsecase.SaveOrder(c, input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	c.JSON(http.StatusOK, models.Success{
		Status:  http.StatusOK,
		Message: "success",
	})
}

// UpdateOrder godoc
// @Summary Update Order
// @Description Update Order Data
// @Tags Order
// @Security BearerAuth
// @Param Authorization header string true "Bearer Token"
// @Accept json
// @Produce json
// @Param order body models.Order true "Form Order"
// @Param id path string true "Order ID"
// @Success 200 {object} models.Success "Successful response"
// @Router /order/{id} [put]
func (h *Handler) UpdateOrder(c *gin.Context) {
	var input *models.Order

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	orderID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input.ID = orderID

	err = h.OrderUsecase.UpdateOrder(c, input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	c.JSON(http.StatusOK, models.Success{
		Status:  http.StatusOK,
		Message: "success",
	})
}

// GetOrderByID godoc
// @Summary Get Order By ID
// @Description Get Order By ID
// @Tags Order
// @Security BearerAuth
// @Param Authorization header string true "Bearer Token"
// @Accept json
// @Produce json
// @Param id path string true "Order ID"
// @Success 200 {object} models.Success "Successful response"
// @Router /order/{id} [get]
func (h *Handler) GetOrderByID(c *gin.Context) {
	orderID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order, err := h.OrderUsecase.GetOrderByID(c, orderID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	c.JSON(http.StatusOK, models.Success{
		Status:  http.StatusOK,
		Message: "success",
		Data:    order,
	})
}

// GetOrders godoc
// @Summary Get List Order
// @Description Get List Order
// @Tags Order
// @Security BearerAuth
// @Param Authorization header string true "Bearer Token"
// @Param page query int false "Page number"
// @Param limit query int false "Limit"
// @Param page query int false "Page number"
// @Param sorting query int false "Sorting"
// @Param search query int false "Search"
// @Accept json
// @Produce json
// @Success 200 {array} models.Success
// @Router /order/list [get]
func (h *Handler) GetOrders(c *gin.Context) {
	limit, err := strconv.ParseInt(c.Query("limit"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "limit value is incorrect."})
		return
	}

	page, err := strconv.ParseInt(c.Query("page"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "page value is incorrect."})
		return
	}

	param := models.ParamOrder{
		Limit:   uint64(limit),
		Page:    uint64(page),
		Sorting: c.Query("sorting"),
		Search:  c.Query("search"),
	}

	order, err := h.OrderUsecase.GetOrderList(c, param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	c.JSON(http.StatusOK, models.Success{
		Status:  http.StatusOK,
		Message: "success",
		Data:    order,
	})
}
