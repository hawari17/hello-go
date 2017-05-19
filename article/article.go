package article

import (
	"errors"
	"time"
)

// ErrAlreadyExists : Error stating that an article can't be stored because it already exists
// ErrNotFound : Error stating that no requested article can be found by provided condition (filters)
var (
	ErrAlreadyExists = errors.New("Article already exists")
	ErrNotFound      = errors.New("Article not found")
)

type Article struct {
	ID        int
	Title     string
	Content   string
	URL       string
	CreatedAt time.Time
}
