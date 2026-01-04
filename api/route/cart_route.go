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

func NewCartRouter(env *bootstrap.Env, timeout time.Duration, db *sql.DB, group *gin.RouterGroup) {
	_ = env // reserved for future environment-driven behavior
	cr := repository.NewCartRepository(db)
	cc := &controller.CartController{
		CartUsecase: usecase.NewCartUsecase(cr, timeout),
	}

	group.POST("/cart", cc.AddToCart)
	group.GET("/user/:user_id/cart", cc.GetCart)
	group.DELETE("/cart/:id", cc.DeleteCartItem)
}
