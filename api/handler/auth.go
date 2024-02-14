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
	if len(createUser.RoleId) == 0 {
		userRole, err := h.strg.Role().GetByID(context.Background(), &models.RolePrimaryKey{Type: "Oquvchi"})
		if err != nil {
			h.handlerResponse(c, "error Role.GetByID", http.StatusInternalServerError, err.Error())
			return
		}
		createUser.RoleId = userRole.Id
		id, err = h.strg.User().Create(context.Background(), &createUser)
		if err != nil {
			h.handlerResponse(c, "storage.user.create", http.StatusInternalServerError, err.Error())
			return
		}
		resp, err := h.strg.User().GetByID(context.Background(), &models.UserPrimaryKey{Id: id})
		if err != nil {
			h.handlerResponse(c, "storage.user.getbyid", http.StatusInternalServerError, err.Error())
			return
		}
		h.handlerResponse(c, "User Created Successfully!", http.StatusOK, resp)
		return
	}

	id, err = h.strg.User().Create(context.Background(), &createUser)
	if err != nil {
		h.handlerResponse(c, "storage.user.create", http.StatusInternalServerError, err.Error())
		return
	}

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
	Id := models.ConfirmCode{RequestId: requestId}
	checkEmail.RequestId = Id.RequestId

	h.handlerResponse(c, "Email Sent Successfully!", http.StatusCreated, checkEmail)
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

// SendEmail godoc
// @ID /auth/send_exist_email
// @Router /auth/send_exist_email [POST]
// @Summary SendEmail
// @Description SendEmail
// @Tags Auth
// @Accept json
// @Procedure json
// @Param checkEmail body models.CheckEmail true "CheckEmailRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) SendCodeExistEmail(c *gin.Context) {
	var email models.CheckEmail

	err := c.ShouldBindJSON(&email)
	if err != nil {
		h.handlerResponse(c, "error Should Bind Json SendCodeExistEmail", http.StatusBadRequest, err.Error())
		return
	}

	_, err = h.strg.User().GetByID(context.Background(), &models.UserPrimaryKey{Email: email.Email})
	if err != nil {
		h.handlerResponse(c, "User already exist!", http.StatusOK, error.Error(errors.New("User already exist!")))
		return
	}
	verifyCode, err := helper.SendEmail(email.Email)
	if err != nil {
		h.handlerResponse(c, "error in sendcodeexistEmail helper.SendEmail", http.StatusInternalServerError, err.Error())
		return
	}

	requestId, err := h.strg.User().CheckOTP(context.Background(), &email, verifyCode)
	if err != nil {
		h.handlerResponse(c, "error in strg.User.CheckOTP", http.StatusInternalServerError, err.Error())
		return
	}
	reqId := models.CheckEmail{RequestId: requestId}
	h.handlerResponse(c, "Email sent successfully!", http.StatusOK, reqId)
}

// UpdatePassword godoc
// @ID /auth/update_password
// @Router /auth/update_password [POST]
// @Summary UpdatePassword
// @Description UpdatePassword
// @Tags Auth
// @Accept json
// @Procedure json
// @Param checkEmail body models.UpdatePassword true "UpdatePasswordRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) UpdatePassword(c *gin.Context) {
	var update models.UpdatePassword

	err := c.ShouldBindJSON(&update)
	if err != nil {
		h.handlerResponse(c, "error Should Bind Json UpdatePassword", http.StatusInternalServerError, err.Error())
		return
	}

	_, err = h.strg.User().UpdatePassword(context.Background(), &update)
	if err != nil {
		h.handlerResponse(c, "error in User.UpdatePassword", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "Password updated successfully!", http.StatusOK, "Password updated successfully!")
}
