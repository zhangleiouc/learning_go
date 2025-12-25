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

func NewCouponRouter(env *bootstrap.Env, timeout time.Duration, db *sql.DB, group *gin.RouterGroup) {
	_ = env // reserved for future environment-driven behavior
	cr := repository.NewCouponRepository(db)
	cc := &controller.CouponController{
		CouponUsecase: usecase.NewCouponUsecase(cr, timeout),
	}

	group.POST("/user/coupons", cc.GetByUserID)
}
