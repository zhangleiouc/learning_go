package route

import (
	"database/sql"
	"time"

	"github.com/gin-gonic/gin"
	"learning_go/bootstrap"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db *sql.DB, gin *gin.Engine) {
	publicRouter := gin.Group("")
	// All Public APIs
	NewOrderRouter(env, timeout, db, publicRouter)
}
