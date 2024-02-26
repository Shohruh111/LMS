package handler

import (
	"app/api/models"
	"app/pkg/helper"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	studentFilter = "Student"
	mentorFilter  = "Mentor"
)

// Students Excel godoc
// @ID get_list_students_excel
// @Router /lms/api/excel/students [GET]
// @Summary Get List Students
// @Description Get List Students Excel Format
// @Tags Excel
// @Accept json
// @Procedure json
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) StudentsExcelDownload(c *gin.Context) {

	var (
		headers = []string{"ID", "Phone Number", "First Name", "Last Name", "Email"}
		datas   [][]interface{}
	)

	resp, err := h.strg.User().GetAllUsersForExcel(c.Request.Context(), &models.UserGetListRequest{Filter: studentFilter})
	if err != nil {
		h.logger.Error("error StudentsExcelDownload User.GetAllStudentsForExcel")
		c.JSON(http.StatusInternalServerError, "Serve Error!")
		return
	}
	for count, val := range resp.Users {
		datas = append(datas, []interface{}{count + 1, val.PhoneNumber, val.FirstName, val.LastName, val.Email})
	}
	excel := helper.Excel{
		Headers: headers,
		Datas:   datas,
	}

	fmt.Println(resp)

	file, err := helper.GenerateExcelFile(&excel)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error generating Excel file")
		return
	}

	fileName := "students.xlsx"

	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header("Content-Disposition", "attachment; filename="+fileName)

	err = file.Write(c.Writer)
	if err != nil {
		fmt.Println("Error writing to response:", err)
		c.String(http.StatusInternalServerError, "Error writing Excel file to response")
	}
}

// Mentors Excel godoc
// @ID get_list_mentors_excel
// @Router /lms/api/excel/mentors [GET]
// @Summary Get List Mentors
// @Description Get List Mentors Excel Format
// @Tags Excel
// @Accept json
// @Procedure json
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) MentorsExcelDownload(c *gin.Context) {

	var (
		headers = []string{"ID", "Phone Number", "First Name", "Last Name", "Email"}
		datas   [][]interface{}
	)

	resp, err := h.strg.User().GetAllUsersForExcel(c.Request.Context(), &models.UserGetListRequest{Filter: mentorFilter})
	if err != nil {
		h.logger.Error("error MentorsExcelDownload User.GetAllUsersForExcel")
		c.JSON(http.StatusInternalServerError, "Serve Error!")
		return
	}
	for count, val := range resp.Users {
		datas = append(datas, []interface{}{count + 1, val.PhoneNumber, val.FirstName, val.LastName, val.Email})
	}
	excel := helper.Excel{
		Headers: headers,
		Datas:   datas,
	}

	fmt.Println(resp)

	file, err := helper.GenerateExcelFile(&excel)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error generating Excel file")
		return
	}

	fileName := "mentors.xlsx"

	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header("Content-Disposition", "attachment; filename="+fileName)

	err = file.Write(c.Writer)
	if err != nil {
		fmt.Println("Error writing to response:", err)
		c.String(http.StatusInternalServerError, "Error writing Excel file to response")
	}
}
