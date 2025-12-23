package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"learning_go/domain"
)

type CouponController struct {
	CouponUsecase domain.CouponUsecase
}

// GetByUserID returns coupon details for a given user id.
func (cc *CouponController) GetByUserID(c *gin.Context) {
	var req domain.UserCouponRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{
			Message: "Invalid request body: " + err.Error(),
		})
		return
	}

	coupons, err := cc.CouponUsecase.GetByUserID(c.Request.Context(), req.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{
			Message: "Failed to get coupons: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Data: coupons,
	})
}
