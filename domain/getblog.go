package domain

import (
	"fmt"
	"github.com/deusdat/cleango"
)

type GetSingleArticleUseCase struct {
	ArticleReader ArticleReader
}

func (g *GetSingleArticleUseCase) Execute(articleID ArticleID, p cleango.Presenter[Article]) {
	articles, err := g.ArticleReader.Read(struct {
		ArticleID     ArticleID
		LastArticleID ArticleID
	}{ArticleID: articleID})
	op := "getsinglearticleusecase.execute "
	if err != nil {
		p.Present(struct {
			Answer Article
			Err    error
		}{Err: cleango.ToDomainError(op+"failed to get article repo", err)})
		return
	}
	if len(articles) == 0 {
		p.Present(struct {
			Answer Article
			Err    error
		}{Err: &cleango.DomainError{
			Kind:    cleango.NotFound,
			Message: fmt.Sprintf("%s could not find %s", op, articleID),
		}})
		return
	}
	if len(articles) != 1 {
		p.Present(struct {
			Answer Article
			Err    error
		}{Err: &cleango.DomainError{
			Kind:    cleango.NotFound,
			Message: fmt.Sprintf("%s too many articles with %s", op, articleID),
		}})
		return
	}
	p.Present(struct {
		Answer Article
		Err    error
	}{
		Answer: articles[0],
	})

}
