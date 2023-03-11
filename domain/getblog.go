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
	if err != nil {
		p.Present(struct {
			Answer Article
			Err    error
		}{Err: fmt.Errorf("GetSingleArticleUseCase.Execute %w", err)})
		return
	}
	if len(articles) == 0 {
		p.Present(struct {
			Answer Article
			Err    error
		}{Err: &cleango.DomainError{
			Kind:    cleango.NotFound,
			Message: fmt.Sprintf("GetSingleArticleUseCase.Execute could not find %s", articleID),
		}})
		return
	}
	if len(articles) != 1 {
		p.Present(struct {
			Answer Article
			Err    error
		}{Err: &cleango.DomainError{
			Kind:    cleango.NotFound,
			Message: fmt.Sprintf("GetSingleArticleUseCase.Execute too many articles with %s", articleID),
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
