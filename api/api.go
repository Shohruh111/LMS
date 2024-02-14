package api

import (
	_ "app/api/docs"
	"app/api/handler"
	"app/config"
	"app/pkg/logger"
	"app/storage"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewApi(r *gin.Engine, cfg *config.Config, storage storage.StorageI, logger logger.LoggerI) {

	// @securityDefinitions.apikey ApiKeyAuth
	// @in header
	// @name Authorization

	handler := handler.NewHandler(cfg, storage, logger)

	r.Use(customCORSMiddleware())

	v1 := r.Group("/auth")
	// Login Api
	v1.POST("/login", handler.Login)

	// Register Api
	v1.POST("/register", handler.Register)
	v1.POST("/check_email", handler.CheckEmail)
	v1.POST("/check_code", handler.CheckCode)
	v1.POST("/send_exist_email", handler.SendCodeExistEmail)
	v1.POST("/update_password", handler.UpdatePassword)

	r.POST("/user", handler.CreateUser)
	r.GET("/user/:id", handler.GetByIdUser)
	r.GET("/user", handler.GetListUser)
	r.PUT("/user/:id", handler.UpdateUser)
	r.DELETE("/user/:id", handler.DeleteUser)

	r.POST("/role", handler.CreateRole)
	r.GET("/role/:id", handler.GetByIdRole)
	r.GET("/role", handler.GetListRole)
	r.PUT("/role/:id", handler.UpdateRole)
	r.DELETE("/role/:id", handler.DeleteRole)

	r.POST("/course", handler.CreateCourse)
	r.GET("/course/:id", handler.GetByIdCourse)
	r.GET("/course", handler.GetListCourse)
	r.PUT("/course/:id", handler.UpdateCourse)
	r.DELETE("/course/:id", handler.DeleteCourse)
	r.POST("/photo_upload", handler.PhotoUpload)
	r.GET("/photo/:id", handler.PhotoGet)

	url := ginSwagger.URL("swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}

func customCORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, PUT, PATCH, DELETE, OPTIONS, HEAD")
		c.Header("Access-Control-Allow-Headers", "Platform-Id, Content-Type, Accesp-Encoding, Authorization, Cache-Control")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
