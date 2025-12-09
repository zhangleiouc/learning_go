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

func NewProductRouter(env *bootstrap.Env, timeout time.Duration, db *sql.DB, group *gin.RouterGroup) {
	_ = env // reserved for future environment-driven behavior
	pr := repository.NewProductRepository(db)
	pc := &controller.ProductController{
		ProductUsecase: usecase.NewProductUsecase(pr, timeout),
	}

	group.GET("/products", pc.Search)
}
