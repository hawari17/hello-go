package repository

import (
	"github.com/hawari17/hello-go/article"
)

type ArticleRepository interface {
	Insert(a *article.Article) error
	Update(a *article.Article) error
	Delete(id int) error
	SelectByID(id int) (*article.Article, error)
	SelectByURL(url string) (*article.Article, error)
}
