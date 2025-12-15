package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"learning_go/domain"
)

type InvoiceController struct {
	InvoiceUsecase domain.InvoiceUsecase
}

// GetByOrderID returns invoices for a given order id.
func (ic *InvoiceController) GetByOrderID(c *gin.Context) {
	orderIDParam := c.Param("order_id")
	orderID, err := strconv.ParseInt(orderIDParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{
			Message: "Invalid order ID",
		})
		return
	}

	invoices, err := ic.InvoiceUsecase.GetByOrderID(c.Request.Context(), orderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{
			Message: "Failed to get invoices: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Data: invoices,
	})
}
