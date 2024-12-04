package repository

import (
	"context"
	"gkgkgkgk/internal/repository/model"
)

type ArticleRepository interface {
	NewArticle(ctx context.Context, article model.Article) (int, error)
	ArticleById(ctx context.Context, id int) (*model.Article, error)
	Articles(ctx context.Context) ([]model.Article, error)
	DeleteArticle(ctx context.Context, id int) error
}
