package web

import (
	"context"
	"github.com/deusdat/cleango"
	"simple-blog/domain"
)

// FactoryFactory is a little homage to Java. This ties the internal structure of the factory
// to the request's context.
type FactoryFactory = func(ctx context.Context) Factory

// Factory defines what's needed to get to the data.
type Factory interface {
	// GetArticlesUseCase returns the use case responsible for fetching the articles.
	GetArticlesUseCase() cleango.UseCase[domain.ArticlePaging, domain.GetArticleResults]
	// GetArticlesPresenter the appropriate presenter to show the articles.
	GetArticlesPresenter() cleango.Presenter[domain.GetArticleResults]
}
