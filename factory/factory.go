package factory

import (
	"context"
	"database/sql"
	"github.com/deusdat/cleango"
	_ "modernc.org/sqlite"
	"net/http"
	"simple-blog/domain"
	"simple-blog/sqlite"
	"simple-blog/web/presenters"
)

type Factory struct {
	tx *sql.Tx
}

func NewFactory(ctx context.Context, dbLocation string) (Factory, error) {
	db, err := sqlite.InitDB(dbLocation)
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return Factory{}, &cleango.DomainError{
			Kind:            cleango.System,
			Message:         "failed to get db",
			UnderlyingCause: nil,
			Issues:          nil,
		}
	}
	return Factory{tx: tx}, nil
}

func (f *Factory) NewBlogCrud() *sqlite.BlogCrud {
	return &sqlite.BlogCrud{
		Tx: f.tx,
	}
}

func (f *Factory) BlogPostPresenter(w http.ResponseWriter, r *http.Request) cleango.Presenter[domain.Article] {
	return &presenters.BlogPostPresenter{
		Tx:                       f.tx,
		OriginalSource:           0,
		Writer:                   w,
		Request:                  r,
		ValidationErrorPresenter: presenters.GetEditBlogPostPresenter{},
	}
}

func (f *Factory) NewCreateOrModifyPost() cleango.UseCase[domain.Article, domain.Article] {
	return &domain.CreateOrModifyPost{
		Writer: f.NewBlogCrud(),
	}
}

func (f *Factory) NewGetArticles() cleango.UseCase[*domain.ArticleID, domain.GetArticleResults] {
	// Presently configured to work in an in-memory model.
	return &domain.GetArticleUseCase{}
}

func (f *Factory) NewGetArticlesPresenter(w http.ResponseWriter, r *http.Request) *presenters.GetArticlesPresenter {
	return &presenters.GetArticlesPresenter{
		Writer:  w,
		Request: r,
	}
}
