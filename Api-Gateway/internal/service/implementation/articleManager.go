package implementation

import "api-gateway/internal/domain"

type ArticlesManager struct{}

func (am *ArticlesManager) ListArticles() ([]domain.Article, error)
func (am *ArticlesManager) GetArticle(id int) (domain.Article, error)
func (am *ArticlesManager) Insert(article domain.Article) error
func (am *ArticlesManager) Update(id int, Article domain.Article) error
func (am *ArticlesManager) Delete(id int) (domain.Article, error)
