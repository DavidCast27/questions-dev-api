package tags

import (
	"github.com/DavidCast27/questions-dev-api/internal/questions"
	"time"
)

type Tag struct {
	ID        int8                 `json:"_id,omitempty"`
	Name      string               `json:"name"`
	IsDeleted bool                 `json:"is_deleted"`
	CreatedAt time.Time            `json:"created_at"`
	UpdatedAt time.Time            `json:"updated_at"`
	Questions []questions.Question `json:"questions"`
}
