package models

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"not null"`
	Email     string    `json:"email" gorm:"uniqueIndex;not null"`
	Password  string    `json:"-" gorm:"not null"`
	Role      string    `json:"role" gorm:"not null;default:student"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Course struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title" gorm:"not null"`
	Description string    `json:"description"`
	Category    string    `json:"category" gorm:"index"`
	TeacherID   uint      `json:"teacher_id" gorm:"not null;index"`
	Teacher     User      `json:"teacher" gorm:"foreignKey:TeacherID"`
	Lessons     []Lesson  `json:"lessons,omitempty" gorm:"foreignKey:CourseID"`
	Quizzes     []Quiz    `json:"quizzes,omitempty" gorm:"foreignKey:CourseID"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Lesson struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	CourseID   uint      `json:"course_id" gorm:"not null;index"`
	Title     string    `json:"title" gorm:"not null"`
	Content   string    `json:"content"`
	VideoURL  string    `json:"video_url"`
	Order     int       `json:"order" gorm:"not null;default:0"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Quiz struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	CourseID   uint      `json:"course_id" gorm:"not null;index"`
	Title     string     `json:"title" gorm:"not null"`
	Questions []Question `json:"questions,omitempty" gorm:"foreignKey:QuizID"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

type Question struct {
	ID            uint     `json:"id" gorm:"primaryKey"`
	QuizID        uint     `json:"quiz_id" gorm:"not null;index"`
	Text          string   `json:"text" gorm:"not null"`
	Choices       []Choice `json:"choices,omitempty" gorm:"foreignKey:QuestionID"`
	CorrectAnswer string   `json:"correct_answer" gorm:"not null"`
}

type Choice struct {
	ID         uint   `json:"id" gorm:"primaryKey"`
	QuestionID uint   `json:"question_id" gorm:"not null;index"`
	Text       string `json:"text" gorm:"not null"`
	Label      string `json:"label" gorm:"not null"`
}

type QuizResult struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	QuizID    uint      `json:"quiz_id" gorm:"not null;index"`
	UserID    uint      `json:"user_id" gorm:"not null;index"`
	Score     float64   `json:"score"`
	Total     int       `json:"total"`
	Correct   int       `json:"correct"`
	Answers   []QuizAnswer `json:"answers,omitempty" gorm:"foreignKey:QuizResultID"`
	CreatedAt time.Time `json:"created_at"`
}

type QuizAnswer struct {
	ID           uint   `json:"id" gorm:"primaryKey"`
	QuizResultID uint   `json:"quiz_result_id" gorm:"not null;index"`
	QuestionID   uint   `json:"question_id" gorm:"not null"`
	Answer       string `json:"answer"`
	IsCorrect    bool   `json:"is_correct"`
}

type Enrollment struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id" gorm:"not null;index"`
	CourseID  uint      `json:"course_id" gorm:"not null;index"`
	User      User      `json:"user" gorm:"foreignKey:UserID"`
	Course    Course    `json:"course" gorm:"foreignKey:CourseID"`
	CreatedAt time.Time `json:"created_at"`
}

type Progress struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id" gorm:"not null;index"`
	CourseID  uint      `json:"course_id" gorm:"not null;index"`
	LessonID  uint      `json:"lesson_id" gorm:"not null;index"`
	Completed bool      `json:"completed" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at"`
}
