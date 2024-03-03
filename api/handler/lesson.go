package handler

import (
	"app/api/models"
	"app/pkg/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create Lessons godoc
// @ID create_lessons
// @Router /lms/api/v1/lesson [POST]
// @Summary Create Lessons
// @Description Create Lessons
// @Tags Lesson
// @Accept json
// @Procedure json
// @Param Lesson body models.LessonCreate true "CreateLessonRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) CreateLesson(c *gin.Context) {

	var createLesson models.LessonCreate
	err := c.ShouldBindJSON(&createLesson)
	if err != nil {
		h.logger.Error("error Lesson Create Should Bind Json!")
		c.JSON(http.StatusBadRequest, "error Lesson Create Should Bind Json!")
		return
	}

	id, err := h.strg.Lesson().Create(c.Request.Context(), &createLesson)
	if err != nil {
		h.logger.Error("storage.Lesson.Create!" + err.Error())
		c.JSON(http.StatusInternalServerError, "storage.Lesson.Create!")
		return
	}

	resp, err := h.strg.Lesson().GetByID(c.Request.Context(), &models.LessonPrimaryKey{ID: id})
	if err != nil {
		h.logger.Error("storage.Lesson.GetByID!")
		c.JSON(http.StatusInternalServerError, "storage.Lesson.GetByID")
		return
	}

	h.logger.Info("Create Lesson Response!")
	c.JSON(http.StatusCreated, resp)
}

// GetByID Lessons godoc
// @ID get_by_id_lessons
// @Router /lms/api/v1/lesson/{id} [GET]
// @Summary Get By ID Lessons
// @Description Get By ID Lessons
// @Tags Lesson
// @Accept json
// @Procedure json
// @Param id path string false "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetByIdLesson(c *gin.Context) {
	var id string

	if !helper.IsValidUUID(id) {
		h.logger.Error("is not valid uuid!")
		c.JSON(http.StatusBadRequest, "invalid id")
		return
	}

	resp, err := h.strg.Lesson().GetByID(c.Request.Context(), &models.LessonPrimaryKey{ID: id})
	if err != nil {
		h.logger.Error("storage.Lesson.GetByID!")
		c.JSON(http.StatusInternalServerError, "storage.Lesson.GetByID")
		return
	}

	h.logger.Error("GetByID Lesson Response!")
	c.JSON(http.StatusOK, resp)
}

// GetList Lessons godoc
// @ID get_list_lessons
// @Router /lms/api/v1/lesson [GET]
// @Summary Get List Lessons
// @Description Get List Lessons
// @Tags Lesson
// @Accept json
// @Procedure json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param course_id query string false "course_id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetListLesson(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.logger.Error("GetListLesson INVALID Offset!")
		c.JSON(http.StatusBadRequest, "Invalid offset!")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.logger.Error("GetListLesson INVALID Limit!")
		c.JSON(http.StatusBadRequest, "Invalid Limit")
		return
	}
	courseId := c.Query("course_id")

	resp, err := h.strg.Lesson().GetList(c.Request.Context(), &models.LessonGetListRequest{
		Offset:   offset,
		Limit:    limit,
		CourseId: courseId,
	})
	if err != nil {
		h.logger.Error("storage.Lesson.GetList!" + err.Error())
		c.JSON(http.StatusInternalServerError, "storage.Lesson.GetList")
		return
	}

	h.logger.Info("GetListLesson Response!")
	c.JSON(http.StatusOK, resp)
}

// Update Lessons godoc
// @ID update_lessons
// @Router /lms/api/v1/lesson/{id} [PUT]
// @Summary Update Lessons
// @Description Update Lessons
// @Tags Lesson
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Param Lesson body models.LessonUpdate true "UpdateLessonRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) UpdateLesson(c *gin.Context) {

	var (
		id           string = c.Param("id")
		updateLesson models.LessonUpdate
	)

	if !helper.IsValidUUID(id) {
		h.logger.Error("is invalid uuid!")
		c.JSON(http.StatusBadRequest, "invalid id")
		return
	}

	err := c.ShouldBindJSON(&updateLesson)
	if err != nil {
		h.logger.Error("error Lesson Should Bind Json!")
		c.JSON(http.StatusBadRequest, "error Lesson Should Bind Json")
		return
	}

	updateLesson.ID = id
	rowsAffected, err := h.strg.Lesson().Update(c.Request.Context(), &updateLesson)
	if err != nil {
		h.logger.Error("storage.Lesson.Update!")
		c.JSON(http.StatusInternalServerError, "storage.Lesson.Update")
		return
	}

	if rowsAffected <= 0 {
		h.logger.Error("storage.Lesson.Update!")
		c.JSON(http.StatusBadRequest, "no rows affected")
		return
	}

	resp, err := h.strg.Lesson().GetByID(c.Request.Context(), &models.LessonPrimaryKey{ID: updateLesson.ID})
	if err != nil {
		h.logger.Error("storage.Lesson.GetByID!")
		c.JSON(http.StatusInternalServerError, "storage.Lesson.GetByID")
		return
	}

	h.logger.Info("Update Lesson Successfully!")
	c.JSON(http.StatusAccepted, resp)
}

// Delete Lessons godoc
// @ID delete_lessons
// @Router /lms/api/v1/lesson/{id} [DELETE]
// @Summary Delete Lessons
// @Description Delete Lessons
// @Tags Lesson
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) DeleteLesson(c *gin.Context) {

	var id string = c.Param("id")

	if !helper.IsValidUUID(id) {
		h.logger.Error("is nor valid uuid!")
		c.JSON(http.StatusBadRequest, "invalid id")
		return
	}

	err := h.strg.Lesson().Delete(c.Request.Context(), &models.LessonPrimaryKey{ID: id})
	if err != nil {
		h.logger.Error("storage.Lesson.DELETE!")
		c.JSON(http.StatusInternalServerError, "storage.Lesson.DELETE")
		return
	}

	h.handlerResponse(c, "delete Lesson resposne", http.StatusNoContent, nil)
	h.logger.Info("Deleted Lesson Successfully!")
	c.JSON(http.StatusNoContent, nil)
}
