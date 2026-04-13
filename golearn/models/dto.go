package models

// Auth
type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Role     string `json:"role" binding:"required,oneof=teacher student"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type AuthResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

// Course
type CreateCourseRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Category    string `json:"category"`
}

type UpdateCourseRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Category    string `json:"category"`
}

// Lesson
type CreateLessonRequest struct {
	Title    string `json:"title" binding:"required"`
	Content  string `json:"content"`
	VideoURL string `json:"video_url"`
	Order    int    `json:"order"`
}

type UpdateLessonRequest struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	VideoURL string `json:"video_url"`
	Order    int    `json:"order"`
}

// Quiz
type CreateQuizRequest struct {
	Title     string                  `json:"title" binding:"required"`
	Questions []CreateQuestionRequest `json:"questions" binding:"required,dive"`
}

type CreateQuestionRequest struct {
	Text          string                `json:"text" binding:"required"`
	Choices       []CreateChoiceRequest `json:"choices" binding:"required,dive"`
	CorrectAnswer string                `json:"correct_answer" binding:"required"`
}

type CreateChoiceRequest struct {
	Label string `json:"label" binding:"required"`
	Text  string `json:"text" binding:"required"`
}

type SubmitQuizRequest struct {
	Answers []SubmitAnswerRequest `json:"answers" binding:"required"`
}

type SubmitAnswerRequest struct {
	QuestionID uint   `json:"question_id" binding:"required"`
	Answer     string `json:"answer" binding:"required"`
}

// Progress
type CompleteLessonRequest struct {
	CourseID uint `json:"course_id" binding:"required"`
	LessonID uint `json:"lesson_id" binding:"required"`
}

// Pagination
type PaginationQuery struct {
	Page     int    `form:"page,default=1"`
	Limit    int    `form:"limit,default=10"`
	Category string `form:"category"`
	Sort     string `form:"sort,default=created_at desc"`
}

// Generic Response
type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type PaginatedResponse struct {
	Data       interface{} `json:"data"`
	Page       int         `json:"page"`
	Limit      int         `json:"limit"`
	TotalRows  int64       `json:"total_rows"`
	TotalPages int         `json:"total_pages"`
}

// Progress Response
type ProgressResponse struct {
	CourseID         uint   `json:"course_id"`
	CourseTitle       string `json:"course_title"`
	TotalLessons     int64  `json:"total_lessons"`
	CompletedLessons int64  `json:"completed_lessons"`
	Percentage       float64 `json:"percentage"`
}
