package router

import (
	"go_clean_architecture/Infrastructure/middleware"
	"go_clean_architecture/bootstrap"
	"go_clean_architecture/delivery/controller"
	"go_clean_architecture/domain"
	"go_clean_architecture/repository"
	"go_clean_architecture/usecases"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db mongo.Database, gin *gin.Engine) {
	// Initialize services
	jwtService := middleware.NewJWTService(env.AccessTokenSecret)
	middleware.SetJWTService(jwtService)
	passwordService := middleware.NewPasswordService()

	publicRouter := gin.Group("/api/v1")
	NewSignUpRoutes(publicRouter, env, timeout, db, passwordService)
	NewLoginRoutes(publicRouter, env, timeout, db, passwordService, jwtService)
	NewRefreshTokenRoutes(publicRouter, env, timeout, db, jwtService)

	protectedRouter := gin.Group("/user/me")
	protectedRouter.Use(middleware.AuthMiddleware())
	NewTaskRoutes(protectedRouter, env, timeout, db)
}

func NewSignUpRoutes(router *gin.RouterGroup, env *bootstrap.Env, timeout time.Duration, db mongo.Database, passwordService *middleware.PasswordService) {
	sr := repository.NewSignUpRepository(db, domain.CollectionUser, passwordService)
	sc := &controller.SignUpController{
		SignUpUsecase: usecases.NewSignUpUsecase(sr, timeout),
	}
	router.POST("/signup", sc.SignUp)
	router.GET("/users", sc.GetUser)
}

func NewLoginRoutes(router *gin.RouterGroup, env *bootstrap.Env, timeout time.Duration, db mongo.Database, passwordService *middleware.PasswordService, jwtService *middleware.JWTService) {
	lr := repository.NewLoginRepository(db, domain.CollectionUser, passwordService, jwtService)
	lc := &controller.LoginController{
		LoginUsecase: usecases.NewLoginUsecase(lr, timeout),
	}
	router.POST("/login", lc.Login)
}

func NewRefreshTokenRoutes(router *gin.RouterGroup, env *bootstrap.Env, timeout time.Duration, db mongo.Database, jwtService *middleware.JWTService) {
	rr := repository.NewRefreshTokenRepository(db, domain.CollectionUser, jwtService)
	rc := &controller.RefreshTokenController{
		RefreshTokenUsecase: usecases.NewRefreshTokenUsecase(rr, timeout),
	}
	router.POST("/refresh-token", rc.RefreshToken)
}

func NewTaskRoutes(router *gin.RouterGroup, env *bootstrap.Env, timeout time.Duration, db mongo.Database) {
	tr := repository.NewTaskRepository(db, domain.CollectionTask)
	tc := &controller.TaskController{
		TaskUsecase: usecases.NewTaskUsecase(tr, timeout),
	}
	router.POST("/task", tc.CreateTask)
	router.GET("/task", tc.GetAllTasks)
	router.GET("/task/:id", tc.GetTask)
	router.PUT("/task/:id", tc.UpdateTask)
	router.DELETE("/task/:id", middleware.RoleMiddleware("admin"), tc.DeleteTask)
}
