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

func NewInvoiceRouter(env *bootstrap.Env, timeout time.Duration, db *sql.DB, group *gin.RouterGroup) {
	_ = env // reserved for future environment-driven behavior
	ir := repository.NewInvoiceRepository(db)
	ic := &controller.InvoiceController{
		InvoiceUsecase: usecase.NewInvoiceUsecase(ir, timeout),
	}

	group.GET("/order/:order_id/invoices", ic.GetByOrderID)
}
