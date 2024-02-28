package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"app/api/models"
	"app/pkg/helper"
)

const (
	students = "Student"
	mentors  = "Mentor"
)

// Create Users godoc
// @ID create_users
// @Router /lms/api/v1/user [POST]
// @Summary Create Users
// @Description Create Users
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

	role, err := h.strg.Role().GetByID(c.Request.Context(), &models.RolePrimaryKey{Type: createUser.UserType})
	if err != nil {
		h.logger.Error("error Role.GetByID.CreateUser")
		c.JSON(http.StatusInternalServerError, "Serve Error!")
	}
	createUser.RoleId = role.Id

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

// GetByID Users godoc
// @ID get_by_id_users
// @Router /lms/api/v1/user/{id} [GET]
// @Summary Get By ID Users
// @Description Get By ID Users
// @Tags User
// @Accept json
// @Procedure json
// @Param id path string false "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetByIdUser(c *gin.Context) {
	var id string

	// val := c.GetHeader("Authorization")

	// res1, err := helper.ExtractToken(val)
	// if err != nil {
	// 	h.logger.Error("helper ExtractToken" + err.Error())
	// 	return
	// }

	// res, err := helper.ParseClaims(res1, h.cfg.SecretKey)
	// if err != nil {
	// 	h.logger.Error("token error: \n" + err.Error())
	// 	return
	// }
	// if res["role_id"] != "214fd852-b158-4e9a-9004-1cdc94c72835" {
	// 	h.logger.Error("Unauthorized!")
	// 	c.JSON(401, "UnAuthorized!")
	// 	return
	// }
	id = c.Param("id")

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
	resp.Password = ""
	h.logger.Info("GetByID User Response!")
	c.JSON(http.StatusOK, resp)
}

// GetList Users godoc
// @ID get_list_users
// @Router /lms/api/v1/user [GET]
// @Summary Get List Users
// @Description Get List Users
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

	path := c.FullPath()
	fmt.Println("\t\n: " + path)

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

// Update Users godoc
// @ID update_users
// @Router /lms/api/v1/user/{id} [PUT]
// @Summary Update Users
// @Description Update Users
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

	h.logger.Info("Update User Response!")
	c.JSON(http.StatusAccepted, resp)
}

// Delete Users godoc
// @ID delete_users
// @Router /lms/api/v1/user/{id} [DELETE]
// @Summary Delete Users
// @Description Delete Users
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

// GetList Students godoc
// @ID get_list_students
// @Router /lms/api/v1/students [GET]
// @Summary Get List Students
// @Description Get List Students
// @Tags User
// @Accept json
// @Procedure json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
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

	search, err := h.getSearchQuery(c.Query("search"))
	if err != nil {
		h.logger.Error("GetListStudents INVALID SEARCH!")
		c.JSON(http.StatusBadRequest, "INVALID SEARCH")
		return
	}

	resp, err := h.strg.User().GetList(c.Request.Context(), &models.UserGetListRequest{
		Offset: offset,
		Limit:  limit,
		Filter: students,
		Search: search,
	})

	if err != nil {
		h.logger.Error("storage.User.GetListStudents!" + err.Error())
		c.JSON(http.StatusInternalServerError, "storage.User.GetLustStudents")
		return
	}

	h.logger.Info("GetList Students Response!")
	c.JSON(http.StatusOK, resp)
}

// GetList Mentors godoc
// @ID get_list_mentors
// @Router /lms/api/v1/mentors [GET]
// @Summary Get List Mentors
// @Description Get List Mentors
// @Tags User
// @Accept json
// @Procedure json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
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

	search, err := h.getSearchQuery(c.Query("search"))
	if err != nil {
		h.logger.Error("GetListMentors INVALID SEARCH!")
		c.JSON(http.StatusBadRequest, "INVALID SEARCH")
		return
	}

	resp, err := h.strg.User().GetList(c.Request.Context(), &models.UserGetListRequest{
		Offset: offset,
		Limit:  limit,
		Filter: mentors,
		Search: search,
	})

	if err != nil {
		h.logger.Error("storage.User.GetListMentors!")
		c.JSON(http.StatusInternalServerError, "storage.User.GetListMentors!")
		return
	}

	h.logger.Info("GetListMentors Response!")
	c.JSON(http.StatusOK, resp)
}
