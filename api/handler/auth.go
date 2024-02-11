package handler

import (
	"app/api/models"
	"app/pkg/helper"
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login godoc
// @ID /auth/login
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
// @ID /auth/register
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
	var (
		id string
	)
	err := c.ShouldBindJSON(&createUser)
	if err != nil {
		h.handlerResponse(c, "error user should bind json", http.StatusBadRequest, err.Error())
		return
	}
	// id, err = h.strg.User().Create(context.Background(), &createUser)
	// if err != nil {
	// 	h.handlerResponse(c, "storage.user.create", http.StatusInternalServerError, err.Error())
	// 	return
	// }

	resp, err := h.strg.User().GetByID(context.Background(), &models.UserPrimaryKey{Id: id})

	h.handlerResponse(c, "create user response", http.StatusCreated, resp)
}

// CheckEmail godoc
// @ID /auth/check_email
// @Router /auth/check_email [POST]
// @Summary CheckEmail
// @Description CheckEmail
// @Tags Auth
// @Accept json
// @Procedure json
// @Param checkEmail body models.CheckEmail true "CheckEmailRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) CheckEmail(c *gin.Context) {
	var checkEmail models.CheckEmail
	var requestId string
	err := c.ShouldBindJSON(&checkEmail)
	if err != nil {
		h.handlerResponse(c, "error CheckEmail Auth", http.StatusInternalServerError, err.Error())
		return
	}
	_, err = h.strg.User().GetByID(context.Background(), &models.UserPrimaryKey{Email: checkEmail.Email})
	if err != nil {
		if err.Error() == "no rows in result set" {
			verifyCode, err := helper.SendEmail(checkEmail.Email)
			if err != nil {
				h.handlerResponse(c, "error sendEMail", http.StatusInternalServerError, err.Error())
				return
			}
			requestId, err = h.strg.User().CheckOTP(context.Background(), &checkEmail, verifyCode)
			if err != nil {
				h.handlerResponse(c, "error in strg.User.CheckOTP", http.StatusInternalServerError, err.Error())
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

	h.handlerResponse(c, "Email Sent Successfully!", http.StatusCreated, requestId)
}

// CheckCode godoc
// @ID /auth/check_code
// @Router /auth/check_code [POST]
// @Summary CheckCode
// @Description CheckCode
// @Tags Auth
// @Accept json
// @Procedure json
// @Param checkEmail body models.CheckCode true "CheckCodeRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) CheckCode(c *gin.Context) {
	var code models.CheckCode

	err := c.ShouldBindJSON(&code)
	if err != nil {
		h.handlerResponse(c, "error Should Bind Json CheckCode", http.StatusBadRequest, err.Error())
		return
	}

	ConfirmCode, err := h.strg.User().GetOTP(context.Background(), &code)
	if err != nil {
		h.handlerResponse(c, "error User.GetOTP", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "Succcessfully checked", http.StatusOK, ConfirmCode)
}
