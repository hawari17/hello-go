package usecase

import (
	"errors"
	"testing"
	"time"

	"github.com/hawari17/hello-go/article"
	"github.com/hawari17/hello-go/article/repository/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestStore(t *testing.T) {
	repo := new(mocks.ArticleRepository)
	repo.On("SelectByURL", mock.AnythingOfType("string")).Return(nil, nil)
	repo.On("Insert", mock.AnythingOfType("*article.Article")).Return(nil)

	u := NewArticleUseCase(repo)

	a := &article.Article{
		Title:     "Title",
		Content:   "Content",
		URL:       "http://www.test-url.com",
		CreatedAt: time.Now()}

	err := u.Store(a)
	assert.NoError(t, err)

	repo.AssertCalled(t, "SelectByURL", mock.AnythingOfType("string"))
	repo.AssertCalled(t, "Insert", mock.AnythingOfType("*article.Article"))
}

func TestStoreAndArticleAlreadyExists(t *testing.T) {
	repo := new(mocks.ArticleRepository)
	repo.On("SelectByURL", mock.AnythingOfType("string")).Return(&article.Article{}, nil)
	repo.On("Insert", mock.AnythingOfType("*article.Article")).Return(nil)

	u := NewArticleUseCase(repo)

	a := &article.Article{
		Title:     "Title",
		Content:   "Content",
		URL:       "http://www.test-url.com",
		CreatedAt: time.Now()}

	err := u.Store(a)
	assert.Error(t, err)

	repo.AssertCalled(t, "SelectByURL", mock.AnythingOfType("string"))
	repo.AssertNotCalled(t, "Insert", mock.AnythingOfType("*article.Article"))
}

func TestStoreAndRepositoryReturnsError(t *testing.T) {
	repo := new(mocks.ArticleRepository)
	repo.On("SelectByURL", mock.AnythingOfType("string")).Return(nil, errors.New("Undefined Errors"))
	repo.On("Insert", mock.AnythingOfType("*article.Article")).Return(nil)

	u := NewArticleUseCase(repo)

	a := &article.Article{
		Title:     "Title",
		Content:   "Content",
		URL:       "http://www.test-url.com",
		CreatedAt: time.Now()}

	err := u.Store(a)
	assert.Error(t, err)

	repo.AssertCalled(t, "SelectByURL", mock.AnythingOfType("string"))
	repo.AssertNotCalled(t, "Insert", mock.AnythingOfType("*article.Article"))
}

func TestFindByID(t *testing.T) {
	repo := new(mocks.ArticleRepository)
	repo.On("SelectByID", mock.AnythingOfType("int")).Return(&article.Article{}, nil)

	u := NewArticleUseCase(repo)

	_, err := u.FindByID(123)
	assert.NoError(t, err, "Should not produces error")

	repo.AssertCalled(t, "SelectByID", mock.AnythingOfType("int"))
}

func TestFindByIDAndArticleNotFound(t *testing.T) {
	repo := new(mocks.ArticleRepository)
	repo.On("SelectByID", mock.AnythingOfType("int")).Return(nil, nil)

	u := NewArticleUseCase(repo)

	_, err := u.FindByID(123)
	assert.Error(t, err, "Should produces error")

	repo.AssertCalled(t, "SelectByID", mock.AnythingOfType("int"))
}

func TestFindByIDAndRepositoryThrowsError(t *testing.T) {
	repo := new(mocks.ArticleRepository)
	repo.On("SelectByID", mock.AnythingOfType("int")).Return(nil, errors.New("Undefined Errors"))

	u := NewArticleUseCase(repo)

	_, err := u.FindByID(123)
	assert.Error(t, err, "Should produces error")

	repo.AssertCalled(t, "SelectByID", mock.AnythingOfType("int"))
}

func TestUpdate(t *testing.T) {
	repo := new(mocks.ArticleRepository)
	repo.On("SelectByID", mock.AnythingOfType("int")).Return(&article.Article{}, nil)
	repo.On("Update", mock.AnythingOfType("*article.Article")).Return(nil)

	u := NewArticleUseCase(repo)

	a := &article.Article{
		Title:     "Title",
		Content:   "Content",
		URL:       "http://www.test-url.com",
		CreatedAt: time.Now()}

	err := u.Update(123, a)
	assert.NoError(t, err, "Should not produces error")

	repo.AssertCalled(t, "SelectByID", mock.AnythingOfType("int"))
	repo.AssertCalled(t, "Update", mock.AnythingOfType("*article.Article"))
}

func TestUpdateAndArticleNotFound(t *testing.T) {
	repo := new(mocks.ArticleRepository)
	repo.On("SelectByID", mock.AnythingOfType("int")).Return(nil, nil)
	repo.On("Update", mock.AnythingOfType("*article.Article")).Return(nil)

	u := NewArticleUseCase(repo)

	a := &article.Article{
		Title:     "Title",
		Content:   "Content",
		URL:       "http://www.test-url.com",
		CreatedAt: time.Now()}

	err := u.Update(123, a)
	assert.Error(t, err, "Should produces error")

	repo.AssertCalled(t, "SelectByID", mock.AnythingOfType("int"))
	repo.AssertNotCalled(t, "Update", mock.AnythingOfType("*article.Article"))
}

func TestUpdateAndRepositoryThrowsError(t *testing.T) {
	repo := new(mocks.ArticleRepository)
	repo.On("SelectByID", mock.AnythingOfType("int")).Return(nil, errors.New("Undefined Errors"))
	repo.On("Update", mock.AnythingOfType("*article.Article")).Return(nil)

	u := NewArticleUseCase(repo)

	a := &article.Article{
		Title:     "Title",
		Content:   "Content",
		URL:       "http://www.test-url.com",
		CreatedAt: time.Now()}

	err := u.Update(123, a)
	assert.Error(t, err, "Should produces error")

	repo.AssertCalled(t, "SelectByID", mock.AnythingOfType("int"))
	repo.AssertNotCalled(t, "Update", mock.AnythingOfType("*article.Article"))
}

func TestUpdateAndRepositoryThrowsErrorWhenUpdate(t *testing.T) {
	repo := new(mocks.ArticleRepository)
	repo.On("SelectByID", mock.AnythingOfType("int")).Return(&article.Article{}, nil)
	repo.On("Update", mock.AnythingOfType("*article.Article")).Return(errors.New("Error when updating"))

	u := NewArticleUseCase(repo)

	a := &article.Article{
		Title:     "Title",
		Content:   "Content",
		URL:       "http://www.test-url.com",
		CreatedAt: time.Now()}

	err := u.Update(123, a)
	assert.Error(t, err, "Should produces error")

	repo.AssertCalled(t, "SelectByID", mock.AnythingOfType("int"))
	repo.AssertCalled(t, "Update", mock.AnythingOfType("*article.Article"))
}

func TestDelete(t *testing.T) {
	repo := new(mocks.ArticleRepository)
	repo.On("SelectByID", mock.AnythingOfType("int")).Return(&article.Article{}, nil)
	repo.On("Delete", mock.AnythingOfType("int")).Return(nil)

	u := NewArticleUseCase(repo)

	err := u.Delete(123)
	assert.NoError(t, err, "Should not produces error")

	repo.AssertCalled(t, "SelectByID", mock.AnythingOfType("int"))
	repo.AssertCalled(t, "Delete", mock.AnythingOfType("int"))
}

func TestDeleteAndArticleNotFound(t *testing.T) {
	repo := new(mocks.ArticleRepository)
	repo.On("SelectByID", mock.AnythingOfType("int")).Return(nil, nil)
	repo.On("Delete", mock.AnythingOfType("int")).Return(nil)

	u := NewArticleUseCase(repo)

	err := u.Delete(123)
	assert.Error(t, err, "Should not produces error")

	repo.AssertCalled(t, "SelectByID", mock.AnythingOfType("int"))
	repo.AssertNotCalled(t, "Delete", mock.AnythingOfType("int"))
}

func TestDeleteAndRepositoryThrowsError(t *testing.T) {
	repo := new(mocks.ArticleRepository)
	repo.On("SelectByID", mock.AnythingOfType("int")).Return(nil, errors.New("Undefined errors"))
	repo.On("Delete", mock.AnythingOfType("int")).Return(nil)

	u := NewArticleUseCase(repo)

	err := u.Delete(123)
	assert.Error(t, err, "Should not produces error")

	repo.AssertCalled(t, "SelectByID", mock.AnythingOfType("int"))
	repo.AssertNotCalled(t, "Delete", mock.AnythingOfType("int"))
}

func TestDeleteAndRepositoryThrowsErrorWhenDelete(t *testing.T) {
	repo := new(mocks.ArticleRepository)
	repo.On("SelectByID", mock.AnythingOfType("int")).Return(&article.Article{}, nil)
	repo.On("Delete", mock.AnythingOfType("int")).Return(errors.New("Error when deleting"))

	u := NewArticleUseCase(repo)

	err := u.Delete(123)
	assert.Error(t, err, "Should not produces error")

	repo.AssertCalled(t, "SelectByID", mock.AnythingOfType("int"))
	repo.AssertCalled(t, "Delete", mock.AnythingOfType("int"))
}
