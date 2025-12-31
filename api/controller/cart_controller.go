package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"learning_go/domain"
)

type CartController struct {
	CartUsecase domain.CartUsecase
}

// AddToCart handles POST request to add a product to cart.
func (cc *CartController) AddToCart(c *gin.Context) {
	var request domain.AddToCartRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{
			Message: "Invalid request body: " + err.Error(),
		})
		return
	}

	cartID, err := cc.CartUsecase.AddOrUpdate(c.Request.Context(), request.UserID, request.ProductID, request.Quantity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{
			Message: "Failed to add to cart: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, domain.SuccessResponse{
		Data: map[string]int64{
			"cart_id": cartID,
		},
	})
}

// GetCart handles GET request to retrieve user's cart items.
func (cc *CartController) GetCart(c *gin.Context) {
	userIDParam := c.Param("user_id")
	userID, err := strconv.ParseInt(userIDParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{
			Message: "Invalid user ID",
		})
		return
	}

	cartItems, err := cc.CartUsecase.GetByUserID(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{
			Message: "Failed to get cart: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Data: cartItems,
	})
}

// DeleteCartItem handles DELETE request to remove a cart item.
func (cc *CartController) DeleteCartItem(c *gin.Context) {
	cartIDParam := c.Param("id")
	cartID, err := strconv.ParseInt(cartIDParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{
			Message: "Invalid cart ID",
		})
		return
	}

	err = cc.CartUsecase.Delete(c.Request.Context(), cartID)
	if err != nil {
		if err.Error() == "cart item not found" {
			c.JSON(http.StatusNotFound, domain.ErrorResponse{
				Message: "Cart item not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{
			Message: "Failed to delete cart item: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Cart item deleted successfully",
	})
}
