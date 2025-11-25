package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"learning_go/domain"
)

type OrderController struct {
	OrderUsecase domain.OrderUsecase
}

func (oc *OrderController) GetByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{
			Message: "Invalid order ID",
		})
		return
	}

	order, err := oc.OrderUsecase.GetByID(c.Request.Context(), id)
	if err != nil {
		if err.Error() == "order not found" {
			c.JSON(http.StatusNotFound, domain.ErrorResponse{
				Message: "Order not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Data: order,
	})
}
