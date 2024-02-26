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

	v1 := r.Group("/lms/api/auth")

	// Login Api
	v1.POST("/login", handler.Login)

	// v1.Use(handler.AuthMiddleware())

	// Register Api
	v1.POST("/register", handler.Register)

	v1.POST("/checkEmail", handler.CheckEmail)
	v1.POST("/checkCode", handler.CheckCode)
	v1.POST("/sendExistEmail", handler.SendCodeExistEmail)
	v1.POST("/updatePassword", handler.UpdatePassword)

	v2 := r.Group("/lms/api")

	v2.POST("/user", handler.CreateUser)
	v2.GET("/user/:id", handler.GetByIdUser)
	v2.GET("/user", handler.GetListUser)
	v2.PUT("/user/:id", handler.UpdateUser)
	v2.DELETE("/user/:id", handler.DeleteUser)
	v2.GET("/students", handler.GetListStudents)
	v2.GET("/mentors", handler.GetListMentors)

	v2.GET("/excel/students", handler.StudentsExcelDownload)
	v2.GET("/excel/mentors", handler.MentorsExcelDownload)

	v2.POST("/role", handler.CreateRole)
	v2.GET("/role/:id", handler.GetByIdRole)
	v2.GET("/role", handler.GetListRole)
	v2.PUT("/role/:id", handler.UpdateRole)
	v2.DELETE("/role/:id", handler.DeleteRole)

	v2.POST("/course", handler.CreateCourse)
	v2.GET("/course/:id", handler.GetByIdCourse)
	v2.GET("/course", handler.GetListCourse)
	v2.PUT("/course/:id", handler.UpdateCourse)
	v2.DELETE("/course/:id", handler.DeleteCourse)
	v2.POST("/photoUpload", handler.PhotoUpload)
	v2.GET("/photo/:id", handler.PhotoDownload)
	v2.GET("/course/users", handler.CourseOfUsers)

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
