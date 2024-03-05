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

	v1 := r.Group("/lms/api/v1")

	// Login Api
	v1.POST("/auth/login", handler.Login)

	// v1.Use(handler.AuthMiddleware())

	// Register Api
	v1.POST("/auth/register", handler.Register)

	v1.POST("/auth/checkEmail", handler.CheckEmail)
	v1.POST("/auth/checkCode", handler.CheckCode)
	v1.POST("/auth/restorePassword", handler.RestorePassword)
	v1.POST("/auth/updatePassword", handler.UpdatePassword)

	// v1.GET("/device", handler.DeviceInfo)

	v1.GET("/ping", handler.PingPong)

	v1.POST("/user", handler.CreateUser)
	v1.GET("/user/:id", handler.GetByIdUser)
	v1.GET("/user", handler.GetListUser)
	v1.PUT("/user/:id", handler.UpdateUser)
	v1.DELETE("/user/:id", handler.DeleteUser)
	v1.GET("/students", handler.GetListStudents)
	v1.GET("/mentors", handler.GetListMentors)

	v1.GET("/excel/students", handler.StudentsExcelDownload)
	v1.GET("/excel/mentors", handler.MentorsExcelDownload)
	v1.GET("/excel/courses", handler.CoursesExcelDownload)

	v1.POST("/role", handler.CreateRole)
	v1.GET("/role/:id", handler.GetByIdRole)
	v1.GET("/role", handler.GetListRole)
	v1.PUT("/role/:id", handler.UpdateRole)
	v1.DELETE("/role/:id", handler.DeleteRole)

	v1.POST("/group", handler.CreateGroup)
	v1.GET("/group/:id", handler.GetByIdGroup)
	v1.GET("/group", handler.GetListGroup)
	v1.PUT("/group/:id", handler.UpdateGroup)
	v1.DELETE("/group/:id", handler.DeleteGroup)

	v1.POST("/lesson", handler.CreateLesson)
	v1.GET("/lesson/:id", handler.GetByIdLesson)
	v1.GET("/lesson", handler.GetListLesson)
	v1.PUT("/lesson/:id", handler.UpdateLesson)
	v1.DELETE("/lesson/:id", handler.DeleteLesson)

	v1.POST("/course", handler.CreateCourse)
	v1.GET("/course/:id", handler.GetByIdCourse)
	v1.GET("/course", handler.GetListCourse)
	v1.PUT("/course/:id", handler.UpdateCourse)
	v1.DELETE("/course/:id", handler.DeleteCourse)
	v1.POST("/photoUpload", handler.PhotoUpload)
	v1.GET("/photo/:id", handler.PhotoDownload)
	v1.GET("/course/users", handler.CourseOfUsers)

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
