package handler

import (
	"app/api/models"
	"app/pkg/helper"
	"context"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// Create Courses godoc
// @ID create_Courses
// @Router /lms/api/v1/course [POST]
// @Summary Create Courses
// @Description Create Courses
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
		h.logger.Error("error Course Create Should Bind Json")
		c.JSON(http.StatusBadRequest, "error Course Create Should Bind Json")
		return
	}

	id, err := h.strg.Course().Create(c.Request.Context(), &createCourse)
	if err != nil {
		h.logger.Error("storage.Course.Create!")
		c.JSON(http.StatusInternalServerError, "storage.Course.Create!")
		return
	}

	resp, err := h.strg.Course().GetByID(c.Request.Context(), &models.CoursePrimaryKey{Id: id})
	if err != nil {
		h.logger.Error("storage.Course.GetById!")
		c.JSON(http.StatusInternalServerError, "storage.Course.GetById")
		return
	}

	h.logger.Info("Create Course Response!")
	c.JSON(http.StatusCreated, resp)
}

// @Security ApiKeyAuth
// GetByID Courses godoc
// @ID get_by_id_Courses
// @Router /lms/api/v1/course/{id} [GET]
// @Summary Get By ID Courses
// @Description Get By ID Courses
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
	id = c.Param("id")
	// // }

	if !helper.IsValidUUID(id) {
		h.logger.Error("is not valid uuid!")
		c.JSON(http.StatusBadRequest, "invalid id")
		return
	}

	resp, err := h.strg.Course().GetByID(c.Request.Context(), &models.CoursePrimaryKey{Id: id})
	if err != nil {
		h.logger.Error("storage.Course.GetById!")
		c.JSON(http.StatusInternalServerError, "storage.Course.GetById")
		return
	}

	h.logger.Info("GetById Course Response!")
	c.JSON(http.StatusOK, resp)
}

// GetList Courses godoc
// @ID get_list_Courses
// @Router /lms/api/v1/course [GET]
// @Summary Get List Courses
// @Description Get List Courses
// @Tags Course
// @Accept json
// @Procedure json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetListCourse(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.logger.Error("GetListCourse Offset")
		c.JSON(http.StatusBadRequest, "GetListCourse Offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.logger.Error("GetListCourse limit!")
		c.JSON(http.StatusBadRequest, "GetListCourse Limit")
		return
	}
	search, err := h.getSearchQuery(c.Query("search"))
	if err != nil {
		h.logger.Error("GetListCourse search!")
		c.JSON(http.StatusBadRequest, "GetListCourse search!")
		return
	}

	resp, err := h.strg.Course().GetList(c.Request.Context(), &models.CourseGetListRequest{
		Offset: offset,
		Limit:  limit,
		Search: search,
	})
	if err != nil {
		h.logger.Error("storage.Course.GetList!" + err.Error())
		c.JSON(http.StatusInternalServerError, "storage.Course.GetList!")
		return
	}
	h.logger.Info("GetListCourse Response!!")
	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// Update Courses godoc
// @ID update_Courses
// @Router /lms/api/v1/course/{id} [PUT]
// @Summary Update Courses
// @Description Update Courses
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
		h.logger.Error("is not valid uuid")
		c.JSON(http.StatusBadRequest, "invalid id")
		return
	}

	err := c.ShouldBindJSON(&updateCourse)
	if err != nil {
		h.logger.Error("error Course should bind json!")
		c.JSON(http.StatusBadRequest, "error Course should bind json")
		return
	}

	updateCourse.Id = id
	rowsAffected, err := h.strg.Course().Update(c.Request.Context(), &updateCourse)
	if err != nil {
		h.logger.Error("storage.Course.Update")
		c.JSON(http.StatusInternalServerError, "storage.Course.Update")
		return
	}

	if rowsAffected <= 0 {
		h.logger.Error("storage.Course.Update")
		c.JSON(http.StatusBadRequest, "no rows affected!")
		return
	}

	resp, err := h.strg.Course().GetByID(c.Request.Context(), &models.CoursePrimaryKey{Id: updateCourse.Id})
	if err != nil {
		h.logger.Error("storage.Course.GetById!")
		c.JSON(http.StatusInternalServerError, "storage.Course.GetById")
		return
	}

	h.logger.Info("Create Course Response!")
	c.JSON(http.StatusAccepted, resp)
}

// @Security ApiKeyAuth
// Delete Courses godoc
// @ID delete_Courses
// @Router /lms/api/v1/course/{id} [DELETE]
// @Summary Delete Courses
// @Description Delete Courses
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
		h.logger.Error("is not valid uuid!")
		c.JSON(http.StatusOK, "invalid uuid")

		return
	}

	err := h.strg.Course().Delete(c.Request.Context(), &models.CoursePrimaryKey{Id: id})
	if err != nil {
		h.logger.Error("storage.Course.Delete!")
		c.JSON(http.StatusInternalServerError, "storage.Course.Delete!")
		return
	}

	h.logger.Info("Create Course Response!")
	c.JSON(http.StatusNoContent, nil)
}

// Upload Photo godoc
// @ID photo_upload
// @Router /lms/api/v1/photoUpload [POST]
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
		h.logger.Error("error VideoLessons!")
		c.JSON(http.StatusInternalServerError, "error VideoLessons!")
		return
	}
	defer videoFile.Close()

	videoData, err := ioutil.ReadAll(videoFile)
	if err != nil {
		h.logger.Error("error ioutil.ReadAll.VideoLessons!")
		c.JSON(http.StatusInternalServerError, "error ioutil.ReadALl.VideoLessons!")
		return
	}
	result, err := h.strg.Course().UploadPhotos(c.Request.Context(), &models.VideoLessons{FileName: file.Filename, VideoData: videoData})
	if err != nil {
		h.logger.Error("error Course.UploadVideos!")
		c.JSON(http.StatusInternalServerError, "error Course.UploadVideos!")
		return
	}

	h.logger.Info("Video Uploaded Successfully!")
	c.JSON(http.StatusOK, result)
}

// Photo Get godoc
// @ID photo_get
// @Router /lms/api/v1/photo/{id} [Get]
// @Summary Photo Get
// @Description Photo Get
// @Tags Course
// @Accept json
// @Procedure json
// @Param id path string false "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) PhotoDownload(c *gin.Context) {

	id := c.Param("id")

	result, err := h.strg.Course().GetPhotos(c.Request.Context(), &models.VideoLessons{Id: id})
	if err != nil {
		h.logger.Error("error PhotoGet.PhotoDownloads!")
		c.JSON(http.StatusInternalServerError, "error PhotoGet.PhotoDownloads!")
		return
	}

	c.Header("Content-Type", "image/jpeg")
	c.Data(http.StatusOK, "photo", result.VideoData)
}

// GetList Course Of Users godoc
// @ID get_list_Course_of_Users
// @Router /lms/api/v1/course/users [GET]
// @Summary Get List Course Of Users
// @Description Get List Course Of Users
// @Tags Course
// @Accept json
// @Procedure json
// @Param id path string false "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) CourseOfUsers(c *gin.Context) {
	Id := c.Param("id")
	resp, err := h.strg.Course().GetListCourseOfUsers(context.Background(), &models.CoursePrimaryKey{Id: Id})
	if err != nil {
		h.logger.Error("error Course.GetListCourseOfUsers")
		c.JSON(http.StatusInternalServerError, "error Course.GetListCourseOfUsers!")
		return
	}

	h.logger.Info("Successfull!")
	c.JSON(http.StatusOK, resp)
}
