package domain

import (
	"time"

	"github.com/google/uuid"
)

type Article struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	OwnerID   uuid.UUID `json:"owner_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Tags      []string  `json:"tags"`
	UpVote    int       `json:"up_vote"`
	DownVote  int       `json:"down_vote"`
	Comments  []Comment `json:"comments"`
}

type Comment struct {
	ID        uuid.UUID `json:"id"`
	OwnerID   uuid.UUID `json:"owner_id"`
	ArticleID uuid.UUID `json:"article_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}
