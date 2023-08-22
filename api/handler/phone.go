package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"app/api/models"
	"app/pkg/helper"
)

// Create phone godoc
// @ID create_user_phone
// @Router /user/phone [POST]
// @Summary Create Phone
// @Description Create Phone
// @Tags Phone
// @Accept json
// @Procedure json
// @Param Phone body models.CreatePhone true "CreatePhoneRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) CreatePhone(c *gin.Context) {

	var createPhone models.CreatePhone
	err := c.ShouldBindJSON(&createPhone)
	if err != nil {
		h.handlerResponse(c, "error Phone should bind json", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.strg.Phone().Create(c.Request.Context(), &createPhone)
	if err != nil {
		h.handlerResponse(c, "storage.Phone.create", http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.strg.Phone().GetByID(c.Request.Context(), &models.PhonePrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.Phone.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create Phone resposne", http.StatusCreated, resp)
}

// GetByID Phone godoc
// @ID get_by_id_Phone
// @Router /user/phone/{id} [GET]
// @Summary Get By ID Phone
// @Description Get By ID Phone
// @Tags Phone
// @Accept json
// @Procedure json
// @Param id path string false "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetByIdPhone(c *gin.Context) {
	var id string

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	resp, err := h.strg.Phone().GetByID(c.Request.Context(), &models.PhonePrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.Phone.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get by id Phone resposne", http.StatusOK, resp)
}

// GetList Phone godoc
// @ID get_list_phone
// @Router /user/phone [GET]
// @Summary Get List Phone
// @Description Get List Phone
// @Tags Phone
// @Accept json
// @Procedure json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetListPhone(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list Phone offset", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list Phone limit", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.strg.Phone().GetList(c.Request.Context(), &models.GetListPhoneRequest{
		Offset: offset,
		Limit:  limit,
	})
	if err != nil {
		h.handlerResponse(c, "storage.Phone.get_list", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list Phone resposne", http.StatusOK, resp)
}

// Update Phone godoc
// @ID update_phone
// @Router /user/phone/{id} [PUT]
// @Summary Update Phone
// @Description Update Phone
// @Tags Phone
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Param Phone body models.UpdatePhone true "UpdatePhoneRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) UpdatePhone(c *gin.Context) {

	var (
		id          string = c.Param("id")
		updatePhone models.UpdatePhone
	)

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := c.ShouldBindJSON(&updatePhone)
	if err != nil {
		h.handlerResponse(c, "error Phone should bind json", http.StatusBadRequest, err.Error())
		return
	}

	updatePhone.Id = id
	rowsAffected, err := h.strg.Phone().Update(c.Request.Context(), &updatePhone)
	if err != nil {
		h.handlerResponse(c, "storage.Phone.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.Phone.update", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.strg.Phone().GetByID(c.Request.Context(), &models.PhonePrimaryKey{Id: updatePhone.Id})
	if err != nil {
		h.handlerResponse(c, "storage.Phone.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create Phone resposne", http.StatusAccepted, resp)
}

// Delete Phone godoc
// @ID delete_Phone
// @Router /user/phone/{id} [DELETE]
// @Summary Delete Phone
// @Description Delete Phone
// @Tags Phone
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) DeletePhone(c *gin.Context) {

	var id string = c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := h.strg.Phone().Delete(c.Request.Context(), &models.PhonePrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.Phone.delete", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create Phone resposne", http.StatusNoContent, nil)
}
