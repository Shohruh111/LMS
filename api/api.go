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

	r.POST("/user", handler.CreateUser)
	r.GET("/user/:id", handler.GetByIdUser)
	r.GET("/user", handler.GetListUser)
	r.PUT("/user/:id", handler.UpdateUser)
	r.DELETE("/user/:id", handler.DeleteUser)

	r.POST("/role", handler.CreateUser)
	r.GET("/role/:id", handler.GetByIdUser)
	r.GET("/role", handler.GetListUser)
	r.PUT("/role/:id", handler.UpdateUser)
	r.DELETE("/role/:id", handler.DeleteUser)

	r.POST("/course", handler.CreateUser)
	r.GET("/course/:id", handler.GetByIdUser)
	r.GET("/course", handler.GetListUser)
	r.PUT("/course/:id", handler.UpdateUser)
	r.DELETE("/course/:id", handler.DeleteUser)

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
