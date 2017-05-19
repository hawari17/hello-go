package usecase

import (
	"github.com/hawari17/hello-go/article"
	"github.com/hawari17/hello-go/article/repository"
)

type ArticleUsecase interface {
	Store(a *article.Article) error
	FindByID(id int) (*article.Article, error)
	Update(id int, a *article.Article) error
	Delete(id int) error
}

type articleUseCase struct {
	ArticleRepository repository.ArticleRepository
}

func (u *articleUseCase) Store(a *article.Article) error {
	articleExists, err := u.ArticleRepository.FindByURL(a.URL)
	if err != nil {
		return err
	}

	if articleExists != nil {
		return article.ErrAlreadyExists
	}

	return u.ArticleRepository.Store(a)
}

func (u *articleUseCase) FindByID(id int) (*article.Article, error) {
	return nil, nil
}

func (u *articleUseCase) Update(id int, a *article.Article) error {
	return nil
}

func (u *articleUseCase) Delete(id int) error {
	return nil
}

func NewArticleUseCase(repo repository.ArticleRepository) ArticleUsecase {
	return &articleUseCase{repo}
}
