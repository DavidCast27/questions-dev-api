package levels

import "time"

type Level struct {
	ID        int8      `json:"_id,omitempty"`
	Name      string    `json:"name"`
	IsDeleted bool      `json:"is_deleted"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
