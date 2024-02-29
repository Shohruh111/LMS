package handler

import (
	"app/api/models"
	"app/pkg/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create Groups godoc
// @ID create_groups
// @Router /lms/api/v1/group [POST]
// @Summary Create Groups
// @Description Create Groups
// @Tags Group
// @Accept json
// @Procedure json
// @Param Group body models.GroupCreate true "CreateGroupRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) CreateGroup(c *gin.Context) {

	var createGroup models.GroupCreate
	err := c.ShouldBindJSON(&createGroup)
	if err != nil {
		h.logger.Error("error Group Create Should Bind Json!")
		c.JSON(http.StatusBadRequest, "error Group Create Should Bind Json!")
		return
	}

	id, err := h.strg.Group().Create(c.Request.Context(), &createGroup)
	if err != nil {
		h.logger.Error("storage.Group.Create!" + err.Error())
		c.JSON(http.StatusInternalServerError, "storage.Group.Create!")
		return
	}

	resp, err := h.strg.Group().GetByID(c.Request.Context(), &models.GroupPrimaryKey{ID: id})
	if err != nil {
		h.logger.Error("storage.Group.GetByID!")
		c.JSON(http.StatusInternalServerError, "storage.Group.GetByID")
		return
	}

	h.logger.Info("Create Group Response!")
	c.JSON(http.StatusCreated, resp)
}

// GetByID Groups godoc
// @ID get_by_id_groups
// @Router /lms/api/v1/group/{id} [GET]
// @Summary Get By ID Groups
// @Description Get By ID Groups
// @Tags Group
// @Accept json
// @Procedure json
// @Param id path string false "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetByIdGroup(c *gin.Context) {
	var id string

	// _, exist := c.Get("Auth")
	// if !exist {
	// 	h.logger.Error("error in AUTH!")
	// 	c.JSON(http.StatusInternalServerError, "error in AUTH!")
	// 	return
	// }

	// GroupData := val.(helper.TokenInfo)
	// if len(GroupData.GroupID) > 0 {
	// 	id = GroupData.GroupID
	// } else {
	// 	id = c.Param("id")
	// }

	if !helper.IsValidUUID(id) {
		h.logger.Error("is not valid uuid!")
		c.JSON(http.StatusBadRequest, "invalid id")
		return
	}

	resp, err := h.strg.Group().GetByID(c.Request.Context(), &models.GroupPrimaryKey{ID: id})
	if err != nil {
		h.logger.Error("storage.Group.GetByID!")
		c.JSON(http.StatusInternalServerError, "storage.Group.GetByID")
		return
	}

	h.logger.Error("GetByID Group Response!")
	c.JSON(http.StatusOK, resp)
}

// GetList Groups godoc
// @ID get_list_groups
// @Router /lms/api/v1/group [GET]
// @Summary Get List Groups
// @Description Get List Groups
// @Tags Group
// @Accept json
// @Procedure json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param course_id query string false "course_id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetListGroup(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.logger.Error("GetListGroup INVALID Offset!")
		c.JSON(http.StatusBadRequest, "Invalid offset!")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.logger.Error("GetListGroup INVALID Limit!")
		c.JSON(http.StatusBadRequest, "Invalid Limit")
		return
	}
	courseId := c.Query("course_id")

	resp, err := h.strg.Group().GetList(c.Request.Context(), &models.GroupGetListRequest{
		Offset:   offset,
		Limit:    limit,
		CourseId: courseId,
	})
	if err != nil {
		h.logger.Error("storage.Group.GetList!" + err.Error())
		c.JSON(http.StatusInternalServerError, "storage.Group.GetList")
		return
	}

	h.logger.Info("GetListGroup Response!")
	c.JSON(http.StatusOK, resp)
}

// Update Groups godoc
// @ID update_groups
// @Router /lms/api/v1/group/{id} [PUT]
// @Summary Update Groups
// @Description Update Groups
// @Tags Group
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Param Group body models.GroupUpdate true "UpdateGroupRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) UpdateGroup(c *gin.Context) {

	var (
		id          string = c.Param("id")
		updateGroup models.GroupUpdate
	)

	if !helper.IsValidUUID(id) {
		h.logger.Error("is invalid uuid!")
		c.JSON(http.StatusBadRequest, "invalid id")
		return
	}

	err := c.ShouldBindJSON(&updateGroup)
	if err != nil {
		h.logger.Error("error Group Should Bind Json!")
		c.JSON(http.StatusBadRequest, "error Group Should Bind Json")
		return
	}

	updateGroup.ID = id
	rowsAffected, err := h.strg.Group().Update(c.Request.Context(), &updateGroup)
	if err != nil {
		h.logger.Error("storage.Group.Update!")
		c.JSON(http.StatusInternalServerError, "storage.Group.Update")
		return
	}

	if rowsAffected <= 0 {
		h.logger.Error("storage.Group.Update!")
		c.JSON(http.StatusBadRequest, "no rows affected")
		return
	}

	resp, err := h.strg.Group().GetByID(c.Request.Context(), &models.GroupPrimaryKey{ID: updateGroup.ID})
	if err != nil {
		h.logger.Error("storage.Group.GetByID!")
		c.JSON(http.StatusInternalServerError, "storage.Group.GetByID")
		return
	}

	h.logger.Info("Update Group Successfully!")
	c.JSON(http.StatusAccepted, resp)
}

// Delete Groups godoc
// @ID delete_groups
// @Router /lms/api/v1/group/{id} [DELETE]
// @Summary Delete Groups
// @Description Delete Groups
// @Tags Group
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) DeleteGroup(c *gin.Context) {

	var id string = c.Param("id")

	if !helper.IsValidUUID(id) {
		h.logger.Error("is nor valid uuid!")
		c.JSON(http.StatusBadRequest, "invalid id")
		return
	}

	err := h.strg.Group().Delete(c.Request.Context(), &models.GroupPrimaryKey{ID: id})
	if err != nil {
		h.logger.Error("storage.Group.DELETE!")
		c.JSON(http.StatusInternalServerError, "storage.Group.DELETE")
		return
	}

	h.handlerResponse(c, "create Group resposne", http.StatusNoContent, nil)
	h.logger.Info("Deleted Group Successfully!")
	c.JSON(http.StatusNoContent, nil)
}
