package interfaces

import "api-gateway/internal/domain"

type IArticlesManager interface {
	ListArticles() ([]domain.Article, error)
	GetArticle(id int) (domain.Article, error)
	Insert(article domain.Article) error
	Update(id int, Article domain.Article) error
	Delete(id int) (domain.Article, error)
}
