package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"app/api/models"
	"app/pkg/helper"
)

const (
	students = "Oquvchi"
	mentors  = "Mentors"
)

// @Security ApiKeyAuth
// Create user godoc
// @ID create_user
// @Router /lms/api/user [POST]
// @Summary Create User
// @Description Create User
// @Tags User
// @Accept json
// @Procedure json
// @Param user body models.UserCreate true "CreateUserRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) CreateUser(c *gin.Context) {

	var createUser models.UserCreate
	err := c.ShouldBindJSON(&createUser)
	if err != nil {
		h.logger.Error("error User Should Bind Json!")
		c.JSON(http.StatusBadRequest, "error User Should Bind Json")
		return
	}

	id, err := h.strg.User().Create(c.Request.Context(), &createUser)
	if err != nil {
		h.logger.Error("storage.User.Create!")
		c.JSON(http.StatusInternalServerError, "storage.User.Create")
		return
	}

	resp, err := h.strg.User().GetByID(c.Request.Context(), &models.UserPrimaryKey{Id: id})
	if err != nil {
		h.logger.Error("storage.User.GetByID!")
		c.JSON(http.StatusInternalServerError, "storage.User.GetByID")
		return
	}

	h.logger.Info("Create User Successfully!!")
	c.JSON(http.StatusCreated, resp)
}

// @Security ApiKeyAuth
// GetByID user godoc
// @ID get_by_id_user
// @Router /lms/api/user/{id} [GET]
// @Summary Get By ID User
// @Description Get By ID User
// @Tags User
// @Accept json
// @Procedure json
// @Param id path string false "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetByIdUser(c *gin.Context) {
	var id string

	// val, exist := c.Get("Auth")
	// if !exist {
	// 	h.handlerResponse(c, "Here", http.StatusInternalServerError, nil)
	// 	return
	// }

	// userData := val.(helper.TokenInfo)
	// if len(userData.UserID) > 0 {
	// 	id = userData.UserID
	// } else {
	id = c.Param("id")
	// }

	if !helper.IsValidUUID(id) {
		h.logger.Error("is valid uuid!")
		c.JSON(http.StatusBadRequest, "invalid id")
		return
	}

	resp, err := h.strg.User().GetByID(c.Request.Context(), &models.UserPrimaryKey{Id: id})
	if err != nil {
		h.logger.Error("storage.User.GetByID!")
		c.JSON(http.StatusInternalServerError, "storage.User.GetByID")
		return
	}

	h.handlerResponse(c, "get by id user resposne", http.StatusOK, resp)
	h.logger.Info("GetByID User Response!")
	c.JSON(http.StatusOK, resp)
}

// GetList user godoc
// @ID get_list_user
// @Router /lms/api/user [GET]
// @Summary Get List User
// @Description Get List User
// @Tags User
// @Accept json
// @Procedure json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param filter query string false "filter"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetListUser(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.logger.Error("GetListUser INVALID OFFSET!")
		c.JSON(http.StatusBadRequest, "INVALID OFFSET")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.logger.Error("GetListUser INVALID LIMIT!")
		c.JSON(http.StatusBadRequest, "INVALID LIMIT")
		return
	}
	filter := c.Query("filter")

	resp, err := h.strg.User().GetList(c.Request.Context(), &models.UserGetListRequest{
		Offset: offset,
		Limit:  limit,
		Filter: filter,
	})
	if err != nil {
		h.logger.Error("storage.User.GetList!")
		c.JSON(http.StatusInternalServerError, "storage.User.GetList!")
		return
	}

	h.logger.Info("GetListUser Response!")
	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// Update user godoc
// @ID update_user
// @Router /lms/api/user/{id} [PUT]
// @Summary Update User
// @Description Update User
// @Tags User
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Param user body models.UserUpdate true "UpdateUserRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) UpdateUser(c *gin.Context) {

	var (
		id         string = c.Param("id")
		updateUser models.UserUpdate
	)

	if !helper.IsValidUUID(id) {
		h.logger.Error("is not valid uuid")
		c.JSON(http.StatusBadRequest, "invalid id")
		return
	}

	err := c.ShouldBindJSON(&updateUser)
	if err != nil {
		h.logger.Error("error User Should Bind Json")
		c.JSON(http.StatusBadRequest, "error User Should Bind Json")
		return
	}

	updateUser.Id = id
	rowsAffected, err := h.strg.User().Update(c.Request.Context(), &updateUser)
	if err != nil {
		h.logger.Error("storage.User.Update!")
		c.JSON(http.StatusInternalServerError, "storage.User.Update")
		return
	}

	if rowsAffected <= 0 {
		h.logger.Error("storage.User.Update!")
		c.JSON(http.StatusBadRequest, "no rows affected!")
		return
	}

	resp, err := h.strg.User().GetByID(c.Request.Context(), &models.UserPrimaryKey{Id: updateUser.Id})
	if err != nil {
		h.logger.Error("storage.User.GetByID!")
		c.JSON(http.StatusInternalServerError, "storage.User.GetByID")
		return
	}

	h.logger.Info("Create User Response!")
	c.JSON(http.StatusAccepted, resp)
}

// @Security ApiKeyAuth
// Delete user godoc
// @ID delete_user
// @Router /lms/api/user/{id} [DELETE]
// @Summary Delete User
// @Description Delete User
// @Tags User
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) DeleteUser(c *gin.Context) {

	var id string = c.Param("id")

	if !helper.IsValidUUID(id) {
		h.logger.Error("is nor valid uuid!")
		c.JSON(http.StatusBadRequest, "invalid id!")
		return
	}

	err := h.strg.User().Delete(c.Request.Context(), &models.UserPrimaryKey{Id: id})
	if err != nil {
		h.logger.Error("storage.User.Delete!")
		c.JSON(http.StatusInternalServerError, "storage.User.Delete")
		return
	}

	h.logger.Info("User Deleted Successfully!")
	c.JSON(http.StatusNoContent, nil)
}

// GetList students godoc
// @ID get_list_students
// @Router /lms/api/students [GET]
// @Summary Get List Students
// @Description Get List Students
// @Tags User
// @Accept json
// @Procedure json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetListStudents(c *gin.Context) {
	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.logger.Error("GetListStudents INVALID OFFSET!")
		c.JSON(http.StatusBadRequest, "INVALID OFFSET")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.logger.Error("GetListStudents INVALID LIMIT!")
		c.JSON(http.StatusBadRequest, "INVALID LIMIT")
		return
	}

	resp, err := h.strg.User().GetList(c.Request.Context(), &models.UserGetListRequest{
		Offset: offset,
		Limit:  limit,
		Filter: students,
	})

	if err != nil {
		h.logger.Error("storage.User.GetListStudents!")
		c.JSON(http.StatusInternalServerError, "storage.User.GetLustStudents")
		return
	}

	h.logger.Info("GetList Students Response!")
	c.JSON(http.StatusOK, resp)
}

// GetList students godoc
// @ID get_list_mentors
// @Router /lms/api/mentors [GET]
// @Summary Get List Mentors
// @Description Get List Mentors
// @Tags User
// @Accept json
// @Procedure json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetListMentors(c *gin.Context) {
	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.logger.Error("GetListMentors INVALID OFFSET!")
		c.JSON(http.StatusBadRequest, "INVALID OFFSET")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.logger.Error("GetListMentors INVALID LIMIT!")
		c.JSON(http.StatusBadRequest, "INVALID LIMIT")
		return
	}

	resp, err := h.strg.User().GetList(c.Request.Context(), &models.UserGetListRequest{
		Offset: offset,
		Limit:  limit,
		Filter: mentors,
	})

	if err != nil {
		h.logger.Error("storage.User.GetListMentors!")
		c.JSON(http.StatusInternalServerError, "storage.User.GetListMentors!")
		return
	}

	h.logger.Info("GetListMentors Response!")
	c.JSON(http.StatusOK, resp)
}
