package controller

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"learning_go/domain"
)

type ProductController struct {
	ProductUsecase domain.ProductUsecase
}

func (pc *ProductController) Search(c *gin.Context) {
	keyword := strings.TrimSpace(c.Query("name"))
	if keyword == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{
			Message: "name query parameter is required",
		})
		return
	}

	products, err := pc.ProductUsecase.SearchByName(c.Request.Context(), keyword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{
			Message: "Failed to search products: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Data: products,
	})
}
