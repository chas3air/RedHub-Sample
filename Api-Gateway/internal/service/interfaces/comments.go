package interfaces

import (
	"api-gateway/internal/domain"

	"github.com/google/uuid"
)

type ICommentsManager interface {
	ListComments() ([]domain.Comment, error)
	ListCommentsByUser(uid uuid.UUID) ([]domain.Comment, error)
	ListCommentsByArticle(aid uuid.UUID) ([]domain.Comment, error)
	InsertCommentToArticle(uid uuid.UUID, aid uuid.UUID, comment domain.Comment) (cid uuid.UUID, err error)
	UpdateCommentToArticle(uid uuid.UUID, old_cid uuid.UUID, new_comment domain.Comment) error
	DeleteCommentToArticle(uid uuid.UUID, cid uuid.UUID) (comment domain.Comment, err error)
}
