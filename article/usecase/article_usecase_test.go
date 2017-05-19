package usecase

import (
  "errors"
	"testing"
  "time"

  "github.com/hawari17/hello-go/article"
  "github.com/hawari17/hello-go/article/repository/mocks"
  "github.com/hawari17/hello-go/article/usecase"

  "github.com/stretchr/testify/assert"
  "github.com/stretchr/testify/mock"
)

func TestStore(t *testing.T) {
  repo := new(mocks.ArticleRepository)
  repo.On("FindByURL", mock.AnythingOfType("string")).Return(nil, nil)
  repo.On("Store", mock.AnythingOfType("*article.Article")).Return(nil)

  u := usecase.NewArticleUseCase(repo)

  a := &article.Article{
    Title: "Title",
    Content: "Content",
    URL: "http://www.test-url.com",
    CreatedAt: time.Now()
  }

  err := u.Store(a)
  assert.NoError(t, err)

  repo.AssertCalled(t, "FindByURL", mock.AnythingOfType("string"))
  repo.AssertCalled(t, "Store", mock.AnythingOfType("*article.Article"))
}

func TestStoreAndArticleAlreadyExists(t *testing.T) {
  repo := new(mocks.ArticleRepository)
  repo.On("FindByURL", mock.AnythingOfType("string")).Return(&article.Article{}, nil)
  repo.On("Store", mock.AnythingOfType("*article.Article")).Return(nil)

  u := usecase.NewArticleUseCase(repo)

  a := &article.Article{
    Title: "Title",
    Content: "Content",
    URL: "http://www.test-url.com",
    CreatedAt: time.Now()
  }

  err := u.Store(a)
  assert.Error(t, err)

  repo.AssertCalled(t, "FindByURL", mock.AnythingOfType("string"))
  repo.AssertNotCalled(t, "Store", mock.AnythingOfType("*article.Article"))
}

func TestStoreAndRepositoryReturnsError(t *testing.T) {
  repo := new(mocks.ArticleRepository)
  repo.On("FindByURL", mock.AnythingOfType("string")).Return(nil, errors.New("Undefined Errors"))
  repo.On("Store", mock.AnythingOfType("*article.Article")).Return(nil)

  u := usecase.NewArticleUseCase(repo)

  a := &article.Article{
    Title: "Title",
    Content: "Content",
    URL: "http://www.test-url.com",
    CreatedAt: time.Now()
  }

  err := u.Store(a)
  assert.Error(t, err)

  repo.AssertCalled(t, "FindByURL", mock.AnythingOfType("string"))
  repo.AssertNotCalled(t, "Store", mock.AnythingOfType("*article.Article"))
}
