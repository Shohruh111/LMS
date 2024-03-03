package handler

import (
	"app/api/models"

	"github.com/gin-gonic/gin"
	"github.com/mssola/user_agent"
)

// Device godoc
// @ID device
// @Router /lms/api/v1/device [GET]
// @Summary Devices
// @Description Devices
// @Tags Device
// @Accept json
// @Procedure json
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) DeviceInfo(c *gin.Context) {
	uaString := c.GetHeader("User-Agent")

	ua := user_agent.New(uaString)

	browser, version := ua.Browser()
	os := ua.OS()

	ip := c.ClientIP()

	c.JSON(200, models.Device{Browser: browser, BrowserVersion: version, DeviceName: os, IP: ip})

}
