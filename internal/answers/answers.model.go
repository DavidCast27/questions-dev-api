package answers

import (
	"github.com/DavidCast27/questions-dev-api/internal/questions"
	"time"
)

type Answer struct {
	ID        int8               `json:"_id,omitempty"`
	Question  questions.Question `json:"question"`
	Content   string             `json:"content"`
	IsCorrect bool               `json:"is_correct"`
	IsDeleted bool               `json:"is_deleted"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}
