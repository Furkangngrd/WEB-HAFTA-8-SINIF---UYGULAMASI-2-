package handlers

import (
	"net/http"
	"strconv"

	"golearn/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateLesson(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		courseID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid course ID"})
			return
		}

		var course models.Course
		if err := db.First(&course, courseID).Error; err != nil {
			c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Course not found"})
			return
		}

		userID := c.GetUint("user_id")
		if course.TeacherID != userID {
			c.JSON(http.StatusForbidden, models.ErrorResponse{Error: "You can only add lessons to your own courses"})
			return
		}

		var req models.CreateLessonRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
			return
		}

		lesson := models.Lesson{
			CourseID:  uint(courseID),
			Title:    req.Title,
			Content:  req.Content,
			VideoURL: req.VideoURL,
			Order:    req.Order,
		}

		if err := db.Create(&lesson).Error; err != nil {
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to create lesson"})
			return
		}

		c.JSON(http.StatusCreated, models.SuccessResponse{
			Message: "Lesson created successfully",
			Data:    lesson,
		})
	}
}

func GetLessons(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		courseID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid course ID"})
			return
		}

		var course models.Course
		if err := db.First(&course, courseID).Error; err != nil {
			c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Course not found"})
			return
		}

		var lessons []models.Lesson
		if err := db.Where("course_id = ?", courseID).Order("\"order\" asc").Find(&lessons).Error; err != nil {
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to fetch lessons"})
			return
		}

		c.JSON(http.StatusOK, models.SuccessResponse{Data: lessons})
	}
}

func GetLesson(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		courseID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid course ID"})
			return
		}

		lessonID, err := strconv.Atoi(c.Param("lessonId"))
		if err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid lesson ID"})
			return
		}

		var lesson models.Lesson
		if err := db.Where("id = ? AND course_id = ?", lessonID, courseID).First(&lesson).Error; err != nil {
			c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Lesson not found"})
			return
		}

		c.JSON(http.StatusOK, models.SuccessResponse{Data: lesson})
	}
}

func UpdateLesson(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		courseID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid course ID"})
			return
		}

		lessonID, err := strconv.Atoi(c.Param("lessonId"))
		if err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid lesson ID"})
			return
		}

		var course models.Course
		if err := db.First(&course, courseID).Error; err != nil {
			c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Course not found"})
			return
		}

		userID := c.GetUint("user_id")
		if course.TeacherID != userID {
			c.JSON(http.StatusForbidden, models.ErrorResponse{Error: "You can only update lessons in your own courses"})
			return
		}

		var lesson models.Lesson
		if err := db.Where("id = ? AND course_id = ?", lessonID, courseID).First(&lesson).Error; err != nil {
			c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Lesson not found"})
			return
		}

		var req models.UpdateLessonRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
			return
		}

		updates := map[string]interface{}{}
		if req.Title != "" {
			updates["title"] = req.Title
		}
		if req.Content != "" {
			updates["content"] = req.Content
		}
		if req.VideoURL != "" {
			updates["video_url"] = req.VideoURL
		}
		if req.Order != 0 {
			updates["order"] = req.Order
		}

		if err := db.Model(&lesson).Updates(updates).Error; err != nil {
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to update lesson"})
			return
		}

		db.Where("id = ? AND course_id = ?", lessonID, courseID).First(&lesson)

		c.JSON(http.StatusOK, models.SuccessResponse{
			Message: "Lesson updated successfully",
			Data:    lesson,
		})
	}
}

func DeleteLesson(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		courseID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid course ID"})
			return
		}

		lessonID, err := strconv.Atoi(c.Param("lessonId"))
		if err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid lesson ID"})
			return
		}

		var course models.Course
		if err := db.First(&course, courseID).Error; err != nil {
			c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Course not found"})
			return
		}

		userID := c.GetUint("user_id")
		if course.TeacherID != userID {
			c.JSON(http.StatusForbidden, models.ErrorResponse{Error: "You can only delete lessons from your own courses"})
			return
		}

		var lesson models.Lesson
		if err := db.Where("id = ? AND course_id = ?", lessonID, courseID).First(&lesson).Error; err != nil {
			c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Lesson not found"})
			return
		}

		if err := db.Delete(&lesson).Error; err != nil {
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to delete lesson"})
			return
		}

		c.JSON(http.StatusOK, models.SuccessResponse{Message: "Lesson deleted successfully"})
	}
}
