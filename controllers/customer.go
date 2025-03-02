package controllers

import (
	"net/http"
	"order-manage/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// SaveCustomer godoc
// @Summary Save Customer
// @Description Insert Customer Data
// @Tags Customer
// @Security BearerAuth
// @Param Authorization header string true "Bearer Token"
// @Accept json
// @Produce json
// @Param order body models.Customer true "Form Customer"
// @Success 200 {object} models.Success "Successful response"
// @Router /customer [post]
func (h *Handler) SaveCustomer(c *gin.Context) {
	var input *models.Customer

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.CustomerUsecase.SaveCustomer(c, input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	c.JSON(http.StatusOK, models.Success{
		Status:  http.StatusOK,
		Message: "success",
	})
}

// UpdateCustomer godoc
// @Summary Update Customer
// @Description Update Customer Data
// @Tags Customer
// @Security BearerAuth
// @Param Authorization header string true "Bearer Token"
// @Accept json
// @Produce json
// @Param order body models.Customer true "Form Customer"
// @Param id path string true "Customer ID"
// @Success 200 {object} models.Success "Successful response"
// @Router /customer/{id} [put]
func (h *Handler) UpdateCustomer(c *gin.Context) {
	var input *models.Customer

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customerID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input.ID = customerID

	err = h.CustomerUsecase.UpdateCustomer(c, input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	c.JSON(http.StatusOK, models.Success{
		Status:  http.StatusOK,
		Message: "success",
	})
}

// GetCustomerByID godoc
// @Summary Get Customer By ID
// @Description Get Customer By ID
// @Tags Customer
// @Security BearerAuth
// @Param Authorization header string true "Bearer Token"
// @Accept json
// @Produce json
// @Param id path string true "Customer ID"
// @Success 200 {object} models.Success "Successful response"
// @Router /customer/{id} [get]
func (h *Handler) GetCustomerByID(c *gin.Context) {
	customerID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customer, err := h.CustomerUsecase.GetCustomerByID(c, customerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	c.JSON(http.StatusOK, models.Success{
		Status:  http.StatusOK,
		Message: "success",
		Data:    customer,
	})
}

// GetCustomers godoc
// @Summary Get List Customer
// @Description Get List Customer
// @Tags Customer
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
// @Router /customer/list [get]
func (h *Handler) GetCustomers(c *gin.Context) {
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

	param := models.ParamCustomer{
		Limit:   uint64(limit),
		Page:    uint64(page),
		Sorting: c.Query("sorting"),
		Search:  c.Query("search"),
	}

	customer, err := h.CustomerUsecase.GetCustomerList(c, param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	c.JSON(http.StatusOK, models.Success{
		Status:  http.StatusOK,
		Message: "success",
		Data:    customer,
	})
}
