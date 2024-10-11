package scores

import (
	"time"
)

type Score struct {
	ID         int8      `json:"_id,omitempty"`
	QuestionID int8      `json:"question_id"`
	Email      string    `json:"email"`
	ScoreTpe   int       `json:"score_tpe"`
	Reason     string    `json:"reason"`
	IsDeleted  bool      `json:"is_deleted"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
