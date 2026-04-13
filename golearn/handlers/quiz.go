package handlers

import (
	"net/http"
	"strconv"

	"golearn/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateQuiz(db *gorm.DB) gin.HandlerFunc {
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
			c.JSON(http.StatusForbidden, models.ErrorResponse{Error: "You can only create quizzes for your own courses"})
			return
		}

		var req models.CreateQuizRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
			return
		}

		quiz := models.Quiz{
			CourseID: uint(courseID),
			Title:   req.Title,
		}

		if err := db.Create(&quiz).Error; err != nil {
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to create quiz"})
			return
		}

		for _, q := range req.Questions {
			question := models.Question{
				QuizID:        quiz.ID,
				Text:          q.Text,
				CorrectAnswer: q.CorrectAnswer,
			}

			if err := db.Create(&question).Error; err != nil {
				c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to create question"})
				return
			}

			for _, ch := range q.Choices {
				choice := models.Choice{
					QuestionID: question.ID,
					Label:      ch.Label,
					Text:       ch.Text,
				}
				if err := db.Create(&choice).Error; err != nil {
					c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to create choice"})
					return
				}
			}
		}

		db.Preload("Questions.Choices").First(&quiz, quiz.ID)

		c.JSON(http.StatusCreated, models.SuccessResponse{
			Message: "Quiz created successfully",
			Data:    quiz,
		})
	}
}

func GetQuizzes(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		courseID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid course ID"})
			return
		}

		var quizzes []models.Quiz
		if err := db.Where("course_id = ?", courseID).Find(&quizzes).Error; err != nil {
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to fetch quizzes"})
			return
		}

		c.JSON(http.StatusOK, models.SuccessResponse{Data: quizzes})
	}
}

func GetQuiz(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		quizID, err := strconv.Atoi(c.Param("quizId"))
		if err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid quiz ID"})
			return
		}

		var quiz models.Quiz
		if err := db.Preload("Questions.Choices").First(&quiz, quizID).Error; err != nil {
			c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Quiz not found"})
			return
		}

		c.JSON(http.StatusOK, models.SuccessResponse{Data: quiz})
	}
}

func SubmitQuiz(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		quizID, err := strconv.Atoi(c.Param("quizId"))
		if err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid quiz ID"})
			return
		}

		var quiz models.Quiz
		if err := db.Preload("Questions").First(&quiz, quizID).Error; err != nil {
			c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Quiz not found"})
			return
		}

		var req models.SubmitQuizRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
			return
		}

		userID := c.GetUint("user_id")

		correctMap := make(map[uint]string)
		for _, q := range quiz.Questions {
			correctMap[q.ID] = q.CorrectAnswer
		}

		total := len(quiz.Questions)
		correct := 0
		var answers []models.QuizAnswer

		for _, ans := range req.Answers {
			isCorrect := false
			if correctAns, ok := correctMap[ans.QuestionID]; ok {
				if correctAns == ans.Answer {
					isCorrect = true
					correct++
				}
			}
			answers = append(answers, models.QuizAnswer{
				QuestionID: ans.QuestionID,
				Answer:     ans.Answer,
				IsCorrect:  isCorrect,
			})
		}

		score := float64(0)
		if total > 0 {
			score = (float64(correct) / float64(total)) * 100
		}

		result := models.QuizResult{
			QuizID:  uint(quizID),
			UserID:  userID,
			Score:   score,
			Total:   total,
			Correct: correct,
			Answers: answers,
		}

		if err := db.Create(&result).Error; err != nil {
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to submit quiz"})
			return
		}

		c.JSON(http.StatusOK, models.SuccessResponse{
			Message: "Quiz submitted successfully",
			Data:    result,
		})
	}
}

func GetQuizResults(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		quizID, err := strconv.Atoi(c.Param("quizId"))
		if err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid quiz ID"})
			return
		}

		userID := c.GetUint("user_id")
		role, _ := c.Get("role")

		var results []models.QuizResult

		query := db.Where("quiz_id = ?", quizID).Preload("Answers")

		if role.(string) == "student" {
			query = query.Where("user_id = ?", userID)
		}

		if err := query.Find(&results).Error; err != nil {
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to fetch results"})
			return
		}

		c.JSON(http.StatusOK, models.SuccessResponse{Data: results})
	}
}
