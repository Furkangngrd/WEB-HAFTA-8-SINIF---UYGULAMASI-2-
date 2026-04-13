package handlers

import (
	"net/http"
	"strconv"

	"golearn/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CompleteLesson(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.CompleteLessonRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
			return
		}

		userID := c.GetUint("user_id")

		// Check enrollment
		var enrollment models.Enrollment
		if err := db.Where("user_id = ? AND course_id = ?", userID, req.CourseID).First(&enrollment).Error; err != nil {
			c.JSON(http.StatusForbidden, models.ErrorResponse{Error: "You are not enrolled in this course"})
			return
		}

		// Check lesson exists in course
		var lesson models.Lesson
		if err := db.Where("id = ? AND course_id = ?", req.LessonID, req.CourseID).First(&lesson).Error; err != nil {
			c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Lesson not found in this course"})
			return
		}

		// Check if already completed
		var existing models.Progress
		if err := db.Where("user_id = ? AND course_id = ? AND lesson_id = ?", userID, req.CourseID, req.LessonID).First(&existing).Error; err == nil {
			c.JSON(http.StatusConflict, models.ErrorResponse{Error: "Lesson already completed"})
			return
		}

		progress := models.Progress{
			UserID:    userID,
			CourseID:  req.CourseID,
			LessonID:  req.LessonID,
			Completed: true,
		}

		if err := db.Create(&progress).Error; err != nil {
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to mark lesson as completed"})
			return
		}

		c.JSON(http.StatusCreated, models.SuccessResponse{
			Message: "Lesson marked as completed",
			Data:    progress,
		})
	}
}

func GetProgress(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		courseID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid course ID"})
			return
		}

		userID := c.GetUint("user_id")

		var course models.Course
		if err := db.First(&course, courseID).Error; err != nil {
			c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Course not found"})
			return
		}

		var totalLessons int64
		db.Model(&models.Lesson{}).Where("course_id = ?", courseID).Count(&totalLessons)

		var completedLessons int64
		db.Model(&models.Progress{}).Where("user_id = ? AND course_id = ? AND completed = ?", userID, courseID, true).Count(&completedLessons)

		percentage := float64(0)
		if totalLessons > 0 {
			percentage = (float64(completedLessons) / float64(totalLessons)) * 100
		}

		response := models.ProgressResponse{
			CourseID:         uint(courseID),
			CourseTitle:      course.Title,
			TotalLessons:     totalLessons,
			CompletedLessons: completedLessons,
			Percentage:       percentage,
		}

		c.JSON(http.StatusOK, models.SuccessResponse{Data: response})
	}
}
