package handler

import (
	"app/api/models"
	"app/pkg/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

// // Device godoc
// // @ID device
// // @Router /lms/api/v1/device [GET]
// // @Summary Devices
// // @Description Devices
// // @Tags Device
// // @Accept json
// // @Procedure json
// // @Success 200 {object} Response{data=string} "Success Request"
// // @Response 400 {object} Response{data=string} "Bad Request"
// // @Failure 500 {object} Response{data=string} "Server error"
// func (h *handler) DeviceInfo(c *gin.Context) {
// 	uaString := c.GetHeader("User-Agent")

// 	ua := user_agent.New(uaString)

// 	browser, version := ua.Browser()
// 	os := ua.OS()

// 	ip := c.ClientIP()

// 	c.JSON(200, models.Device{Browser: browser, BrowserVersion: version, DeviceName: os, IP: ip})

// }

// GetList Devices godoc
// @ID get_list_devices
// @Router /lms/api/v1/device [GET]
// @Summary Get List Devices
// @Description Get List Devices
// @Tags Device
// @Accept json
// @Procedure json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param userId query string false "user_id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetListDevice(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.logger.Error("GetListDevice INVALID Offset!")
		c.JSON(http.StatusBadRequest, "Invalid offset!")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.logger.Error("GetListDevice INVALID Limit!")
		c.JSON(http.StatusBadRequest, "Invalid Limit")
		return
	}
	user_id := c.Query("user_id")

	resp, err := h.strg.Device().GetList(c.Request.Context(), &models.DeviceGetListRequest{
		Offset: offset,
		Limit:  limit,
		UserId: user_id,
	})
	if err != nil {
		h.logger.Error("storage.Device.GetList!")
		c.JSON(http.StatusInternalServerError, "storage.Device.GetList")
		return
	}

	h.logger.Info("GetListDevice Response!")
	c.JSON(http.StatusOK, resp)
}

// Delete Device godoc
// @ID delete_devices
// @Router /lms/api/v1/device/{id} [DELETE]
// @Summary Delete Devices
// @Description Delete Devices
// @Tags Device
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) DeleteDevice(c *gin.Context) {

	var id string = c.Param("id")

	if !helper.IsValidUUID(id) {
		h.logger.Error("is nor valid uuid!")
		c.JSON(http.StatusBadRequest, "invalid id")
		return
	}

	err := h.strg.Role().Delete(c.Request.Context(), &models.RolePrimaryKey{Id: id})
	if err != nil {
		h.logger.Error("storage.Device.DELETE!")
		c.JSON(http.StatusInternalServerError, "storage.Device.DELETE")
		return
	}

	h.handlerResponse(c, "delete device resposne", http.StatusNoContent, nil)
	h.logger.Info("Deleted Device Successfully!")
	c.JSON(http.StatusNoContent, nil)
}
