package handler

import (
	"app/api/models"
	"app/pkg/helper"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// Create Course godoc
// @ID create_Course
// @Router /course [POST]
// @Summary Create Course
// @Description Create Course
// @Tags Course
// @Accept json
// @Procedure json
// @Param Course body models.CourseCreate true "CreateCourseRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) CreateCourse(c *gin.Context) {

	var createCourse models.CourseCreate
	err := c.ShouldBindJSON(&createCourse)
	if err != nil {
		h.handlerResponse(c, "error Course should bind json", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.strg.Course().Create(c.Request.Context(), &createCourse)
	if err != nil {
		h.handlerResponse(c, "storage.Course.create", http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.strg.Course().GetByID(c.Request.Context(), &models.CoursePrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.Course.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create Course resposne", http.StatusCreated, resp)
}

// @Security ApiKeyAuth
// GetByID Course godoc
// @ID get_by_id_Course
// @Router /course/{id} [GET]
// @Summary Get By ID Course
// @Description Get By ID Course
// @Tags Course
// @Accept json
// @Procedure json
// @Param id path string false "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetByIdCourse(c *gin.Context) {
	var id string

	// _, exist := c.Get("Auth")
	// if !exist {
	// 	h.handlerResponse(c, "Here", http.StatusInternalServerError, nil)
	// 	return
	// }

	// CourseData := val.(helper.TokenInfo)
	// if len(CourseData.CourseID) > 0 {
	// 	id = CourseData.CourseID
	// } else {
	// id = c.Param("id")
	// // }

	// if !helper.IsValidUUID(id) {
	// 	h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
	// 	return
	// }

	resp, err := h.strg.Course().GetByID(c.Request.Context(), &models.CoursePrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.Course.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get by id Course resposne", http.StatusOK, resp)
}

// GetList Course godoc
// @ID get_list_Course
// @Router /course [GET]
// @Summary Get List Course
// @Description Get List Course
// @Tags Course
// @Accept json
// @Procedure json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetListCourse(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list Course offset", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list Course limit", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.strg.Course().GetList(c.Request.Context(), &models.CourseGetListRequest{
		Offset: offset,
		Limit:  limit,
	})
	if err != nil {
		h.handlerResponse(c, "storage.Course.get_list", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list Course resposne", http.StatusOK, resp)
}

// @Security ApiKeyAuth
// Update Course godoc
// @ID update_Course
// @Router /course/{id} [PUT]
// @Summary Update Course
// @Description Update Course
// @Tags Course
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Param Course body models.CourseUpdate true "UpdateCourseRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) UpdateCourse(c *gin.Context) {

	var (
		id           string = c.Param("id")
		updateCourse models.CourseUpdate
	)

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := c.ShouldBindJSON(&updateCourse)
	if err != nil {
		h.handlerResponse(c, "error Course should bind json", http.StatusBadRequest, err.Error())
		return
	}

	updateCourse.Id = id
	rowsAffected, err := h.strg.Course().Update(c.Request.Context(), &updateCourse)
	if err != nil {
		h.handlerResponse(c, "storage.Course.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.Course.update", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.strg.Course().GetByID(c.Request.Context(), &models.CoursePrimaryKey{Id: updateCourse.Id})
	if err != nil {
		h.handlerResponse(c, "storage.Course.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create Course resposne", http.StatusAccepted, resp)
}

// @Security ApiKeyAuth
// Delete Course godoc
// @ID delete_Course
// @Router /course/{id} [DELETE]
// @Summary Delete Course
// @Description Delete Course
// @Tags Course
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) DeleteCourse(c *gin.Context) {

	var id string = c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := h.strg.Course().Delete(c.Request.Context(), &models.CoursePrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.Course.delete", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create Course resposne", http.StatusNoContent, nil)
}

// Upload Photo godoc
// @ID photo_upload
// @Router /photo_upload [POST]
// @Summary Photo Upload
// @Description Photo Upload
// @Tags Course
// @Accept json
// @Procedure json
// @Param photo formData file true "Photo file to upload"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) PhotoUpload(c *gin.Context) {

	file, err := c.FormFile("photo")
	if err != nil {
		return
	}

	videoFile, err := file.Open()
	if err != nil {
		h.handlerResponse(c, "error VideoLessons", http.StatusInternalServerError, err.Error())
		return
	}
	defer videoFile.Close()

	videoData, err := ioutil.ReadAll(videoFile)
	if err != nil {
		h.handlerResponse(c, "error ioutil.ReadAll VideoLessons", http.StatusInternalServerError, err.Error())
		return
	}
	result, err := h.strg.Course().UploadPhotos(c.Request.Context(), &models.VideoLessons{FileName: file.Filename, VideoData: videoData})
	if err != nil {
		h.handlerResponse(c, "error Course.UploadVideos", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "Video uploaded successfully!", http.StatusOK, result)
}

// Photo Get godoc
// @ID photo_get
// @Router /photo/{id} [Get]
// @Summary Photo Get
// @Description Photo Get
// @Tags Course
// @Accept json
// @Procedure json
// @Param id path string false "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) PhotoGet(c *gin.Context) {

	id := c.Param("id")

	result, err := h.strg.Course().GetPhotos(c.Request.Context(), &models.VideoLessons{Id: id})
	if err != nil {
		h.handlerResponse(c, "error PhotoGet, GetPhotos", http.StatusInternalServerError, err.Error())
		return
	}

	c.Header("Content-Type", "image/jpeg")
	c.Data(http.StatusOK, "photo", result.VideoData)
}
