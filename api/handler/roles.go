package handler

import (
	"app/api/models"
	"app/pkg/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// Create Role godoc
// @ID create_Role
// @Router /role [POST]
// @Summary Create Role
// @Description Create Role
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
		h.handlerResponse(c, "error Role should bind json", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.strg.Role().Create(c.Request.Context(), &createRole)
	if err != nil {
		h.handlerResponse(c, "storage.Role.create", http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.strg.Role().GetByID(c.Request.Context(), &models.RolePrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.Role.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create Role resposne", http.StatusCreated, resp)
}

// @Security ApiKeyAuth
// GetByID Role godoc
// @ID get_by_id_Role
// @Router /role/{id} [GET]
// @Summary Get By ID Role
// @Description Get By ID Role
// @Tags Role
// @Accept json
// @Procedure json
// @Param id path string false "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetByIdRole(c *gin.Context) {
	var id string

	_, exist := c.Get("Auth")
	if !exist {
		h.handlerResponse(c, "Here", http.StatusInternalServerError, nil)
		return
	}

	// RoleData := val.(helper.TokenInfo)
	// if len(RoleData.RoleID) > 0 {
	// 	id = RoleData.RoleID
	// } else {
	// 	id = c.Param("id")
	// }

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	resp, err := h.strg.Role().GetByID(c.Request.Context(), &models.RolePrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.Role.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get by id Role resposne", http.StatusOK, resp)
}

// GetList Role godoc
// @ID get_list_Role
// @Router /role [GET]
// @Summary Get List Role
// @Description Get List Role
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
		h.handlerResponse(c, "get list Role offset", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list Role limit", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.strg.Role().GetList(c.Request.Context(), &models.RoleGetListRequest{
		Offset: offset,
		Limit:  limit,
	})
	if err != nil {
		h.handlerResponse(c, "storage.Role.get_list", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list Role resposne", http.StatusOK, resp)
}

// @Security ApiKeyAuth
// Update Role godoc
// @ID update_Role
// @Router /role/{id} [PUT]
// @Summary Update Role
// @Description Update Role
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
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := c.ShouldBindJSON(&updateRole)
	if err != nil {
		h.handlerResponse(c, "error Role should bind json", http.StatusBadRequest, err.Error())
		return
	}

	updateRole.Id = id
	rowsAffected, err := h.strg.Role().Update(c.Request.Context(), &updateRole)
	if err != nil {
		h.handlerResponse(c, "storage.Role.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.Role.update", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.strg.Role().GetByID(c.Request.Context(), &models.RolePrimaryKey{Id: updateRole.Id})
	if err != nil {
		h.handlerResponse(c, "storage.Role.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create Role resposne", http.StatusAccepted, resp)
}

// @Security ApiKeyAuth
// Delete Role godoc
// @ID delete_Role
// @Router /role/{id} [DELETE]
// @Summary Delete Role
// @Description Delete Role
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
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := h.strg.Role().Delete(c.Request.Context(), &models.RolePrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.Role.delete", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create Role resposne", http.StatusNoContent, nil)
}
