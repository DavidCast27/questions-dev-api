package answers

import (
	"time"
)

type Answer struct {
	ID         int8      `json:"_id,omitempty"`
	QuestionID int8      `json:"question_id"`
	Content    string    `json:"content"`
	IsCorrect  bool      `json:"is_correct"`
	IsDeleted  bool      `json:"is_deleted"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
