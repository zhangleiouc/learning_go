package route

import (
	"database/sql"
	"time"

	"github.com/gin-gonic/gin"
	"learning_go/api/controller"
	"learning_go/bootstrap"
	"learning_go/repository"
	"learning_go/usecase"
)

func NewOrderRouter(env *bootstrap.Env, timeout time.Duration, db *sql.DB, group *gin.RouterGroup) {
	or := repository.NewOrderRepository(db)
	oc := &controller.OrderController{
		OrderUsecase: usecase.NewOrderUsecase(or, timeout),
	}
	group.GET("/order/:id", oc.GetByID)
	group.POST("/order", oc.Create)
	group.GET("/order/customer/:customer_id", oc.GetByCustomerID)
	group.POST("/order/:id/pay-callback", oc.PayCallback)
}
