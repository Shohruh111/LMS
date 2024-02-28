package handler

import (
	"app/api/models"
	"app/pkg/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)


// Create Roles godoc
// @ID create_Roles
// @Router /lms/api/v1/role [POST]
// @Summary Create Roles
// @Description Create Roles
// @Tags Role
// @Accept json
// @Procedure json
// @Param Role body models.RoleCreate true "CreateRoleRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) CreateRole(c *gin.Context) {

	var createRole models.RoleCreate
	err := c.ShouldBindJSON(&createRole)
	if err != nil {
		h.logger.Error("error Role Create Should Bind Json!")
		c.JSON(http.StatusBadRequest, "error Role Create Should Bind Json!")
		return
	}

	id, err := h.strg.Role().Create(c.Request.Context(), &createRole)
	if err != nil {
		h.logger.Error("storage.Role.Create!")
		c.JSON(http.StatusInternalServerError, "storage.Role.Create!")
		return
	}

	resp, err := h.strg.Role().GetByID(c.Request.Context(), &models.RolePrimaryKey{Id: id})
	if err != nil {
		h.logger.Error("storage.Role.GetByID!")
		c.JSON(http.StatusInternalServerError, "storage.Role.GetByID")
		return
	}

	h.logger.Info("Create Role Response!")
	c.JSON(http.StatusCreated, resp)
}


// GetByID Roles godoc
// @ID get_by_id_Roles
// @Router /lms/api/v1/role/{id} [GET]
// @Summary Get By ID Roles
// @Description Get By ID Roles
// @Tags Role
// @Accept json
// @Procedure json
// @Param id path string false "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetByIdRole(c *gin.Context) {
	var id string

	// _, exist := c.Get("Auth")
	// if !exist {
	// 	h.logger.Error("error in AUTH!")
	// 	c.JSON(http.StatusInternalServerError, "error in AUTH!")
	// 	return
	// }

	// RoleData := val.(helper.TokenInfo)
	// if len(RoleData.RoleID) > 0 {
	// 	id = RoleData.RoleID
	// } else {
	// 	id = c.Param("id")
	// }

	if !helper.IsValidUUID(id) {
		h.logger.Error("is not valid uuid!")
		c.JSON(http.StatusBadRequest, "invalid id")
		return
	}

	resp, err := h.strg.Role().GetByID(c.Request.Context(), &models.RolePrimaryKey{Id: id})
	if err != nil {
		h.logger.Error("storage.Role.GetByID!")
		c.JSON(http.StatusInternalServerError, "storage.Role.GetByID")
		return
	}

	h.logger.Error("GetByID Role Response!")
	c.JSON(http.StatusOK, resp)
}

// GetList Roles godoc
// @ID get_list_Roles
// @Router /lms/api/v1/role [GET]
// @Summary Get List Roles
// @Description Get List Roles
// @Tags Role
// @Accept json
// @Procedure json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetListRole(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.logger.Error("GetListRole INVALID Offset!")
		c.JSON(http.StatusBadRequest, "Invalid offset!")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.logger.Error("GetListRole INVALID Limit!")
		c.JSON(http.StatusBadRequest, "Invalid Limit")
		return
	}

	resp, err := h.strg.Role().GetList(c.Request.Context(), &models.RoleGetListRequest{
		Offset: offset,
		Limit:  limit,
	})
	if err != nil {
		h.logger.Error("storage.Role.GetList!")
		c.JSON(http.StatusInternalServerError, "storage.Role.GetList")
		return
	}

	h.logger.Info("GetListRole Response!")
	c.JSON(http.StatusOK, resp)
}


// Update Roles godoc
// @ID update_Roles
// @Router /lms/api/v1/role/{id} [PUT]
// @Summary Update Roles
// @Description Update Roles
// @Tags Role
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Param Role body models.RoleUpdate true "UpdateRoleRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) UpdateRole(c *gin.Context) {

	var (
		id         string = c.Param("id")
		updateRole models.RoleUpdate
	)

	if !helper.IsValidUUID(id) {
		h.logger.Error("is invalid uuid!")
		c.JSON(http.StatusBadRequest, "invalid id")
		return
	}

	err := c.ShouldBindJSON(&updateRole)
	if err != nil {
		h.logger.Error("error Role Should Bind Json!")
		c.JSON(http.StatusBadRequest, "error Role Should Bind Json")
		return
	}

	updateRole.Id = id
	rowsAffected, err := h.strg.Role().Update(c.Request.Context(), &updateRole)
	if err != nil {
		h.logger.Error("storage.Role.Update!")
		c.JSON(http.StatusInternalServerError, "storage.Role.Update")
		return
	}

	if rowsAffected <= 0 {
		h.logger.Error("storage.Role.Update!")
		c.JSON(http.StatusBadRequest, "no rows affected")
		return
	}

	resp, err := h.strg.Role().GetByID(c.Request.Context(), &models.RolePrimaryKey{Id: updateRole.Id})
	if err != nil {
		h.logger.Error("storage.Role.GetByID!")
		c.JSON(http.StatusInternalServerError, "storage.Role.GetByID")
		return
	}

	h.logger.Info("Update Role Successfully!")
	c.JSON(http.StatusAccepted, resp)
}


// Delete Roles godoc
// @ID delete_Roles
// @Router /lms/api/v1/role/{id} [DELETE]
// @Summary Delete Roles
// @Description Delete Roles
// @Tags Role
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) DeleteRole(c *gin.Context) {

	var id string = c.Param("id")

	if !helper.IsValidUUID(id) {
		h.logger.Error("is nor valid uuid!")
		c.JSON(http.StatusBadRequest, "invalid id")
		return
	}

	err := h.strg.Role().Delete(c.Request.Context(), &models.RolePrimaryKey{Id: id})
	if err != nil {
		h.logger.Error("storage.Role.DELETE!")
		c.JSON(http.StatusInternalServerError, "storage.Role.DELETE")
		return
	}

	h.handlerResponse(c, "create Role resposne", http.StatusNoContent, nil)
	h.logger.Info("Deleted Role Successfully!")
	c.JSON(http.StatusNoContent, nil)
}
