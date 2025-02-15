package implementation

import (
	"api-gateway/internal/domain"

	"github.com/google/uuid"
)

type CommentsManager struct{}

func (cm *CommentsManager) ListComments() ([]domain.Comment, error)
func (cm *CommentsManager) ListCommentsByUser(uid uuid.UUID) ([]domain.Comment, error)
func (cm *CommentsManager) ListCommentsByArticle(aid uuid.UUID) ([]domain.Comment, error)
func (cm *CommentsManager) InsertCommentToArticle(uid uuid.UUID, aid uuid.UUID, comment domain.Comment) (cid uuid.UUID, err error)
func (cm *CommentsManager) UpdateCommentToArticle(uid uuid.UUID, old_cid uuid.UUID, new_comment domain.Comment) error
func (cm *CommentsManager) DeleteCommentToArticle(uid uuid.UUID, cid uuid.UUID) (comment domain.Comment, err error)
