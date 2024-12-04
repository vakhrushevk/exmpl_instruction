package memory

import (
	"context"
	"errors"
	"gkgkgkgk/internal/repository"
	"gkgkgkgk/internal/repository/model"
	"log"
	"sync"
)

type articleRepository struct {
	articles []model.Article
	sync.Mutex
}

// ArticleById implements repository.ArticleRepository.
func (a *articleRepository) ArticleById(ctx context.Context, id int) (*model.Article, error) {
	a.Lock()
	defer a.Unlock()
	for _, article := range a.articles {
		if article.Id == id {
			return &article, nil
		}
	}
	return nil, errors.New("article not found")
}

// Articles implements repository.ArticleRepository.
func (a *articleRepository) Articles(ctx context.Context) ([]model.Article, error) {
	a.Lock()
	defer a.Unlock()

	return a.articles, nil
}

// DeleteArticle implements repository.ArticleRepository.
func (a *articleRepository) DeleteArticle(ctx context.Context, id int) error {
	panic("unimplemented")
}

// NewArticle implements repository.ArticleRepository.
func (a *articleRepository) NewArticle(ctx context.Context, article model.Article) (int, error) {
	a.Lock()
	defer a.Unlock()
	var max int
	if article.Id == 0 {
		log.Println("id i's null find max id")
		for _, article := range a.articles {
			if max < article.Id {
				max = article.Id
			}
		}
		article.Id = max
	}
	a.articles = append(a.articles, article)

	return max, nil
}

func NewAritcleRepository() repository.ArticleRepository {
	articles := make([]model.Article, 0)

	articles = append(articles, model.Article{Id: 2,
		Title: "ti2i",
	})

	return &articleRepository{articles: articles}
}
