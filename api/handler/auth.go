package handler

import (
	"app/api/models"
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login godoc
// @ID login
// @Router /auth/login [POST]
// @Summary Login
// @Description Login
// @Tags Auth
// @Accept json
// @Procedure json
// @Param login body models.LoginUser true "LoginRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) Login(c *gin.Context) {
	var loginUser models.LoginUser
	err := c.ShouldBindJSON(&loginUser)
	if err != nil {
		h.handlerResponse(c, "error login user should bind json", http.StatusBadRequest, err.Error())
		return
	}
	resp, err := h.strg.User().GetByID(context.Background(), &models.UserPrimaryKey{Email: loginUser.Email})
	if err != nil {
		if err.Error() == "no rows in result set" {
			h.handlerResponse(c, "User doesn't exist", http.StatusBadRequest, error.Error(errors.New("User does not exist!")))
			return
		}
		h.handlerResponse(c, "error in User GetByID Login", http.StatusBadRequest, err.Error())
		return
	}
	if resp.Password != loginUser.Password {
		h.handlerResponse(c, "", http.StatusBadRequest, error.Error(errors.New("Please, Enter Valid Code!")))
		return
	}

	h.handlerResponse(c, "Login succesfully", http.StatusOK, resp)
}

// Register godoc
// @ID register
// @Router /auth/register [POST]
// @Summary Register
// @Description Register
// @Tags Auth
// @Accept json
// @Procedure json
// @Param register body models.UserCreate true "CreateUserRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) Register(c *gin.Context) {

	var createUser models.UserCreate
	var id string
	err := c.ShouldBindJSON(&createUser)
	if err != nil {
		h.handlerResponse(c, "error user should bind json", http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.strg.User().GetByID(context.Background(), &models.UserPrimaryKey{Email: createUser.Email})
	if err != nil {
		if err.Error() == "no rows in result set" {
			id, err = h.strg.User().Create(context.Background(), &createUser)
			if err != nil {
				h.handlerResponse(c, "storage.user.create", http.StatusInternalServerError, err.Error())
				return
			}
		} else {
			h.handlerResponse(c, "User already exist", http.StatusInternalServerError, err.Error())
			return
		}
	} else if err == nil {
		h.handlerResponse(c, "User already exist", http.StatusBadRequest, nil)
		return
	}
	resp, err = h.strg.User().GetByID(context.Background(), &models.UserPrimaryKey{Id: id})

	h.handlerResponse(c, "create user response", http.StatusCreated, resp)
}
