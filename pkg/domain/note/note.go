package note

import "time"

// Note is note entity
type Note struct {
	ID        string
	Content   string
	Comment   string
	EditedAt  time.Time
	CreatedAt time.Time
	Tags      []string
}
