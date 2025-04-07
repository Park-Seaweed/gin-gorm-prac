package main

import (
	"gin-test/internal/config"
	"gin-test/internal/config/database"
	"gin-test/internal/controller"
	"gin-test/internal/logger"
	"gin-test/internal/middleware"
	"gin-test/internal/model"
	"gin-test/internal/repository"
	"gin-test/internal/router"
	"gin-test/internal/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	// 로그 초기화
	logger.InitLogger()
	defer logger.Log.Sync()

	cfg, err := config.Load("./config")
	if err != nil {
		logger.Log.Error("", zap.Error(err))
	}

	// DB 연결 및 마이그레이션
	db, err := database.NewDB(cfg)
	if err != nil {
		logger.Log.Error("", zap.Error(err))
	}
	if err := db.AutoMigrate(&model.User{}, &model.Post{}); err != nil {
		logger.Log.Fatal("AutoMigrate 실패", zap.Error(err))
	} else {
		logger.Log.Info("DB 마이그레이션 완료")
	}

	// Repository → Service → Controller
	jwtService := service.NewJwtServiceFromEnv(cfg)
	authController := controller.NewAuthController(jwtService)

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo, jwtService)
	userController := controller.NewUserController(userService)

	postRepo := repository.NewPostRepository(db)
	postService := service.NewPostService(postRepo, userRepo)

	postController := controller.NewPostController(postService) // 아직은 DB 미사용

	// Gin 인스턴스 및 미들웨어 등록
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.GinZapLogger())

	// 라우터 등록
	controllers := &controller.Controllers{
		UserController: userController,
		PostController: postController,
		AuthController: authController,
	}
	router.SetupRouter(r, controllers, jwtService)

	// 서버 시작
	r.Run(":8080")
}
