package models

import (
	"time"

	"github.com/google/uuid"
)

type Article struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	OwnerID   uuid.UUID `json:"owner_id" gorm:"type:uuid;not null"`
	Title     string    `json:"title" gorm:"size:255;not null"`
	Content   string    `json:"content" gorm:"type:text;not null"`
	Tags      []string  `json:"tags" gorm:"-"`
	UpVote    int       `json:"up_vote" gorm:"default:0"`
	DownVote  int       `json:"down_vote" gorm:"default:0"`
	Comments  []Comment `json:"comments" gorm:"-"`
}

type Comment struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key;"`
	OwnerID   uuid.UUID `json:"owner_id" gorm:"type:uuid;not null"`
	ArticleID uuid.UUID `json:"article_id" gorm:"type:uuid;not null"`
	Content   string    `json:"content" gorm:"type:text;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}
