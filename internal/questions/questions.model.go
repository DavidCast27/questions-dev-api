package questions

import (
	"github.com/DavidCast27/questions-dev-api/internal/answers"
	"github.com/DavidCast27/questions-dev-api/internal/categories"
	"github.com/DavidCast27/questions-dev-api/internal/languages"
	"github.com/DavidCast27/questions-dev-api/internal/levels"
	"github.com/DavidCast27/questions-dev-api/internal/scores"
	"github.com/DavidCast27/questions-dev-api/internal/tags"
	"time"
)

type Question struct {
	ID               int8                `json:"_id,omitempty"`
	Title            string              `json:"title"`
	Description      string              `json:"description"`
	Category         categories.Category `json:"category"`
	Language         languages.Language  `json:"language"`
	Level            levels.Level        `json:"level"`
	Tags             []tags.Tag          `json:"tags"`
	Scores           []scores.Score      `json:"scores"`
	Answers          []answers.Answer    `json:"answers"`
	IsMultipleChoice bool                `json:"is_multiple_choice"`
	IsActive         bool                `json:"is_active"`
	IsDeleted        bool                `json:"is_deleted"`
	CreatedAt        time.Time           `json:"created_at"`
	UpdatedAt        time.Time           `json:"updated_at"`
}
