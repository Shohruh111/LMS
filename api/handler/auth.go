package handler

import (
	"app/api/models"
	"app/pkg/helper"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Login godoc
// @ID /auth/login
// @Router /lms/api/auth/login [POST]
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
		h.logger.Error("error LOGIN user should bind json")
		c.JSON(http.StatusBadRequest, "error LOGIN user should bind json")
		return
	}

	resp, err := h.strg.User().GetByID(context.Background(), &models.UserPrimaryKey{Email: loginUser.Email})
	if err != nil {
		if err.Error() == "no rows in result set" {
			h.logger.Error("User doesn't exist")
			c.JSON(http.StatusBadRequest, "User doesn't exist")
			return
		}
		h.logger.Error("error in User GetByID Login")
		c.JSON(http.StatusBadRequest, "error in User GetByID Login")
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(resp.Password), []byte(loginUser.Password))
	if err != nil {
		h.logger.Error("error CompareHashAndPassword")
		c.JSON(http.StatusBadRequest, "Invalid Password!")
		return
	}
	resp.Password = ""

	// result, err := helper.GenerateJWT(map[string]interface{}{"user_id": resp.Id}, time.Hour*360, h.cfg.SecretKey)
	// if err != nil {
	// 	h.logger.Error("error in tokens GENERATEJWT.LOGIN")
	// 	c.JSON(http.StatusInternalServerError, "Server Error!")
	// }

	var credentails = map[string]interface{}{
		"user_id": resp.Id,
	}
	accessToken, err := helper.GenerateJWT(credentails, time.Hour*360, h.cfg.SecretKey)
	if err != nil {
		h.logger.Error("error in tokens GENERATEJWT.LOGIN")
		c.JSON(http.StatusInternalServerError, "Server Error!")
		return
	}

	user := models.LoginResponse{
		AccessToken: accessToken,
		User:        *resp,
	}

	c.JSON(http.StatusOK, user)

}

// @Security ApiKeyAuth
// Register godoc
// @ID /auth/register
// @Router /lms/api/auth/register [POST]
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
		h.logger.Error("error REGISTER user should bind json")
		c.JSON(http.StatusBadRequest, "error REGISTER user should bind json")
		return
	}
	if len(createUser.RoleId) == 0 {
		userRole, err := h.strg.Role().GetByID(context.Background(), &models.RolePrimaryKey{Type: "Student"})
		if err != nil {
			h.logger.Error("error Role.GetByID")
			c.JSON(http.StatusInternalServerError, "error Role.GetByID")
			return
		}
		createUser = models.UserCreate{
			RoleId:      userRole.Id,
			FirstName:   createUser.FirstName,
			LastName:    createUser.LastName,
			Email:       createUser.Email,
			PhoneNumber: createUser.PhoneNumber,
			Password:    createUser.Password,
		}
		id, err = h.strg.User().Create(context.Background(), &createUser)
		if err != nil {
			h.logger.Error("storage.user.create")
			c.JSON(http.StatusInternalServerError, "storage.user.create")
			return
		}
		resp, err := h.strg.User().GetByID(context.Background(), &models.UserPrimaryKey{Id: id})
		if err != nil {
			h.logger.Error("storage.user.getbyid")
			c.JSON(http.StatusInternalServerError, "storage.user.getbyid")
			return
		}
		h.logger.Info("User Created Successfully!")
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err = h.strg.User().Create(context.Background(), &createUser)
	if err != nil {
		h.logger.Error("storage.user.create")
		c.JSON(http.StatusInternalServerError, "storage.user.create")
		return
	}

	resp, err := h.strg.User().GetByID(context.Background(), &models.UserPrimaryKey{Id: id})
	if err != nil {
		h.logger.Error("storage.user.getbyid")
		c.JSON(http.StatusInternalServerError, "storage.user.getbyid")
		return
	}

	h.logger.Info("create user response")
	c.JSON(http.StatusCreated, resp)
}

// @Security ApiKeyAuth
// CheckEmail godoc
// @ID /auth/check_email
// @Router /lms/api/auth/checkEmail [POST]
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
		h.logger.Error("error CheckEmail Auth")
		c.JSON(http.StatusInternalServerError, "error CheckEmail Auth")
		return
	}
	_, err = h.strg.User().GetByID(context.Background(), &models.UserPrimaryKey{Email: checkEmail.Email})
	if err != nil {
		if err.Error() == "no rows in result set" {
			verifyCode, err := helper.SendEmail(checkEmail.Email)
			if err != nil {
				h.logger.Error("error sendEMail")
				c.JSON(http.StatusInternalServerError, "error sendEMail")
				return
			}
			requestId, err = h.strg.User().CheckOTP(context.Background(), &checkEmail, verifyCode)
			if err != nil {
				h.logger.Error("error in strg.User.CheckOTP")
				c.JSON(http.StatusInternalServerError, "error in strg.User.CheckOTP")
				return
			}
		} else {
			h.logger.Error("User already exist")
			c.JSON(http.StatusBadRequest, "User already exist")
			return
		}
	} else if err == nil {
		h.logger.Error("User already exist")
		c.JSON(http.StatusBadRequest, "User already exist")
		return
	}
	Id := models.ConfirmCode{RequestId: requestId}
	checkEmail.RequestId = Id.RequestId

	h.logger.Info("Email sent Successfully!")
	c.JSON(http.StatusCreated, checkEmail)
}

// @Security ApiKeyAuth
// CheckCode godoc
// @ID /auth/check_code
// @Router /lms/api/auth/checkCode [POST]
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
		h.logger.Error("error Should Bind Json CheckCode")
		c.JSON(http.StatusBadRequest, "error Should Bind Json CheckCode")
		return
	}

	ConfirmCode, err := h.strg.User().GetOTP(context.Background(), &code)
	if err != nil {
		h.logger.Error("error in strg.User.GetOTP")
		c.JSON(http.StatusBadRequest, "error in strg.User.GetOTP")
		return
	}

	h.logger.Info("Successfully Checked!")
	c.JSON(http.StatusOK, ConfirmCode)
}

// @Security ApiKeyAuth
// SendEmail godoc
// @ID /auth/send_exist_email
// @Router /lms/api/auth/sendExistEmail [POST]
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
		h.logger.Error("error Should Bind Json SendCodeExistEmail")
		c.JSON(http.StatusBadRequest, "error Should Bind Json SendCodeExistEmail")
		return
	}
	_, err = h.strg.User().GetByID(context.Background(), &models.UserPrimaryKey{Email: email.Email})
	if err != nil {
		if err.Error() == "no rows in result set" {
			h.logger.Error("Email is not registered!")
			c.JSON(http.StatusBadRequest, "Email is not registered!")
			return
		}
		h.logger.Error("error user.GetById.SendCodeExistEmail")
		c.JSON(http.StatusInternalServerError, "error user.GetById.SendCodeExistEmail")
		return
	}
	verifyCode, err := helper.SendEmail(email.Email)
	if err != nil {
		h.logger.Error("error SendCodeExistEmail helper.SendEmail")
		c.JSON(http.StatusBadRequest, "error SendCodeExistEmail helper.SendEmail")
		return
	}

	requestId, err := h.strg.User().CheckOTP(context.Background(), &email, verifyCode)
	if err != nil {
		h.logger.Error("error in strg.User.CheckOTP")
		c.JSON(http.StatusBadRequest, "error in strg.User.CheckOTP")
		return
	}
	informations := models.CheckEmail{RequestId: requestId, Email: email.Email}
	h.logger.Info("Email Sent Successfully!")
	c.JSON(http.StatusOK, informations)
}

// UpdatePassword godoc
// @ID /auth/update_password
// @Router /lms/api/auth/updatePassword [POST]
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
	// val, exist := c.Get("user_id")
	// fmt.Println(val)
	// if !exist {
	// 	c.JSON(http.StatusBadRequest, "invalid token")
	// 	return
	// }
	var update models.UpdatePassword

	err := c.ShouldBindJSON(&update)
	if err != nil {
		h.logger.Error("error Should Bind Json UpdatePassword")
		c.JSON(http.StatusBadRequest, "error Should Bind Json UpdatePassword")
		return
	}

	_, email, err := h.strg.User().UpdatePassword(context.Background(), &update)
	if err != nil {
		h.logger.Error("error in User.UpdatePassword")
		c.JSON(http.StatusInternalServerError, "error in User.UpdatePassword")
		return
	}

	user, err := h.strg.User().GetByID(context.Background(), &models.UserPrimaryKey{Email: email})
	if err != nil {
		h.logger.Error("error User.GetById.UpdatePassword")
		c.JSON(http.StatusInternalServerError, "error User.GetById.UpdatePassword")
		return
	}

	h.logger.Info("Password Updated Successfully!")
	c.JSON(http.StatusOK, user)
}
