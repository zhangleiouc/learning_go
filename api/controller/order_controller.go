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

func (oc *OrderController) Create(c *gin.Context) {
	var request domain.CreateOrderRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{
			Message: "Invalid request body: " + err.Error(),
		})
		return
	}

	orderID, err := oc.OrderUsecase.Create(c.Request.Context(), &request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{
			Message: "Failed to create order: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, domain.SuccessResponse{
		Data: domain.CreateOrderResponse{
			OrderID: orderID,
		},
	})
}

func (oc *OrderController) GetByCustomerID(c *gin.Context) {
	customerIDParam := c.Param("customer_id")
	customerID, err := strconv.ParseInt(customerIDParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{
			Message: "Invalid customer ID",
		})
		return
	}

	orders, err := oc.OrderUsecase.GetByCustomerID(c.Request.Context(), customerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{
			Message: "Failed to get orders: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Data: orders,
	})
}

// PayCallback 订单付款回调接口
// 将订单由“待付款”状态改为“已付款”
// 如果订单状态不是“待付款”，则返回“订单状态已更新，请刷新后尝试”
func (oc *OrderController) PayCallback(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{
			Message: "Invalid order ID",
		})
		return
	}

	err = oc.OrderUsecase.MarkAsPaid(c.Request.Context(), id)
	if err != nil {
		// 订单状态已更新（例如已付款、已取消等），按需求返回固定提示
		if err == domain.ErrOrderStatusAlreadyUpdated {
			c.JSON(http.StatusOK, domain.ErrorResponse{
				Message: "订单状态已更新，请刷新后尝试",
			})
			return
		}

		// 订单不存在
		if err.Error() == "order not found" {
			c.JSON(http.StatusNotFound, domain.ErrorResponse{
				Message: "Order not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{
			Message: "Failed to update order status: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "订单支付成功",
	})
}
