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
	articleExists, err := u.ArticleRepository.SelectByURL(a.URL)
	if err != nil {
		return err
	}

	if articleExists != nil {
		return article.ErrAlreadyExists
	}

	return u.ArticleRepository.Insert(a)
}

func (u *articleUseCase) FindByID(id int) (*article.Article, error) {
	articleReturned, err := u.ArticleRepository.SelectByID(id)
	if err != nil {
		return nil, err
	}

	if articleReturned == nil {
		return nil, article.ErrAlreadyExists
	}

	return articleReturned, nil
}

func (u *articleUseCase) Update(id int, a *article.Article) error {
	articleExists, err := u.ArticleRepository.SelectByID(id)
	if err != nil {
		return err
	}

	if articleExists == nil {
		return article.ErrNotFound
	}

	return u.ArticleRepository.Update(a)
}

func (u *articleUseCase) Delete(id int) error {
	articleExists, err := u.ArticleRepository.SelectByID(id)
	if err != nil {
		return err
	}

	if articleExists == nil {
		return article.ErrNotFound
	}

	return u.ArticleRepository.Delete(id)
}

func NewArticleUseCase(repo repository.ArticleRepository) ArticleUsecase {
	return &articleUseCase{repo}
}
