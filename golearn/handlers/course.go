package handlers

import (
	"math"
	"net/http"
	"strconv"
	"strings"

	"golearn/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateCourse(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.CreateCourseRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
			return
		}

		userID := c.GetUint("user_id")

		course := models.Course{
			Title:       req.Title,
			Description: req.Description,
			Category:    req.Category,
			TeacherID:   userID,
		}

		if err := db.Create(&course).Error; err != nil {
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to create course"})
			return
		}

		db.Preload("Teacher").First(&course, course.ID)

		c.JSON(http.StatusCreated, models.SuccessResponse{
			Message: "Course created successfully",
			Data:    course,
		})
	}
}

func GetCourses(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var pagination models.PaginationQuery
		if err := c.ShouldBindQuery(&pagination); err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
			return
		}

		if pagination.Page < 1 {
			pagination.Page = 1
		}
		if pagination.Limit < 1 || pagination.Limit > 100 {
			pagination.Limit = 10
		}

		var totalRows int64
		query := db.Model(&models.Course{})

		if pagination.Category != "" {
			query = query.Where("category = ?", pagination.Category)
		}

		query.Count(&totalRows)

		// Sort validation
		allowedSorts := map[string]bool{
			"created_at asc": true, "created_at desc": true,
			"title asc": true, "title desc": true,
		}
		sort := pagination.Sort
		if !allowedSorts[strings.ToLower(sort)] {
			sort = "created_at desc"
		}

		offset := (pagination.Page - 1) * pagination.Limit

		var courses []models.Course
		if err := query.Preload("Teacher").
			Order(sort).
			Offset(offset).
			Limit(pagination.Limit).
			Find(&courses).Error; err != nil {
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to fetch courses"})
			return
		}

		totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.Limit)))

		c.JSON(http.StatusOK, models.PaginatedResponse{
			Data:       courses,
			Page:       pagination.Page,
			Limit:      pagination.Limit,
			TotalRows:  totalRows,
			TotalPages: totalPages,
		})
	}
}

func GetCourse(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid course ID"})
			return
		}

		var course models.Course
		if err := db.Preload("Teacher").Preload("Lessons").Preload("Quizzes").First(&course, id).Error; err != nil {
			c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Course not found"})
			return
		}

		c.JSON(http.StatusOK, models.SuccessResponse{Data: course})
	}
}

func UpdateCourse(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid course ID"})
			return
		}

		var course models.Course
		if err := db.First(&course, id).Error; err != nil {
			c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Course not found"})
			return
		}

		userID := c.GetUint("user_id")
		if course.TeacherID != userID {
			c.JSON(http.StatusForbidden, models.ErrorResponse{Error: "You can only update your own courses"})
			return
		}

		var req models.UpdateCourseRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
			return
		}

		updates := map[string]interface{}{}
		if req.Title != "" {
			updates["title"] = req.Title
		}
		if req.Description != "" {
			updates["description"] = req.Description
		}
		if req.Category != "" {
			updates["category"] = req.Category
		}

		if err := db.Model(&course).Updates(updates).Error; err != nil {
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to update course"})
			return
		}

		db.Preload("Teacher").First(&course, id)

		c.JSON(http.StatusOK, models.SuccessResponse{
			Message: "Course updated successfully",
			Data:    course,
		})
	}
}

func DeleteCourse(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid course ID"})
			return
		}

		var course models.Course
		if err := db.First(&course, id).Error; err != nil {
			c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Course not found"})
			return
		}

		userID := c.GetUint("user_id")
		if course.TeacherID != userID {
			c.JSON(http.StatusForbidden, models.ErrorResponse{Error: "You can only delete your own courses"})
			return
		}

		if err := db.Delete(&course).Error; err != nil {
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to delete course"})
			return
		}

		c.JSON(http.StatusOK, models.SuccessResponse{Message: "Course deleted successfully"})
	}
}

func EnrollCourse(db *gorm.DB) gin.HandlerFunc {
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

		var existing models.Enrollment
		if err := db.Where("user_id = ? AND course_id = ?", userID, courseID).First(&existing).Error; err == nil {
			c.JSON(http.StatusConflict, models.ErrorResponse{Error: "Already enrolled in this course"})
			return
		}

		enrollment := models.Enrollment{
			UserID:   userID,
			CourseID: uint(courseID),
		}

		if err := db.Create(&enrollment).Error; err != nil {
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to enroll"})
			return
		}

		c.JSON(http.StatusCreated, models.SuccessResponse{
			Message: "Enrolled successfully",
			Data:    enrollment,
		})
	}
}
