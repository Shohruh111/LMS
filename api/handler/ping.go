package handler

import "github.com/gin-gonic/gin"


// Ping godoc
// @ID ping
// @Router /lms/api/v1/ping [GET]
// @Summary Ping
// @Description Ping
// @Tags Ping
// @Accept json
// @Procedure json
func(h *handler) PingPong(c *gin.Context){
	c.JSON(200, "ping")
}