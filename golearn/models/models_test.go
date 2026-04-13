package models

import (
	"testing"
)

func TestUserRole(t *testing.T) {
	student := User{Role: "student"}
	teacher := User{Role: "teacher"}

	if student.Role != "student" {
		t.Errorf("Expected role 'student', got '%s'", student.Role)
	}

	if teacher.Role != "teacher" {
		t.Errorf("Expected role 'teacher', got '%s'", teacher.Role)
	}
}
