package usecase

import (
	"testing"
  "github.com/stretchr/testify"
  "github.com/hawari17/hello-go/article/usecase"
)

func TestStore(t *testing.T) {
  repo := nil
  u := usecase.NewArticleUseCase(repo)

  a := &Article{
    Title: "Title",
    Content: "Content",
    url: "url"
  }
}
