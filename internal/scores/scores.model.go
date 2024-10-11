package scores

import (
	"github.com/DavidCast27/questions-dev-api/internal/questions"
	"time"
)

type Score struct {
	ID        int8               `json:"_id,omitempty"`
	Question  questions.Question `json:"question"`
	Email     string             `json:"email"`
	ScoreTpe  int                `json:"score_tpe"`
	Reason    string             `json:"reason"`
	IsDeleted bool               `json:"is_deleted"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}
