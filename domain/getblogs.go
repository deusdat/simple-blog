package domain

import (
	"github.com/deusdat/cleango"
)

type GetArticleUseCase struct {
	ArticleReader ArticleReader
}

type GetArticleResults struct {
	Articles []Article
	Next     ArticleID
}

type ArticlePaging struct {
	LastID ArticleID
}

func (g *GetArticleUseCase) Execute(paging ArticlePaging, p cleango.Presenter[GetArticleResults]) {
	articles, err := g.ArticleReader.Read(ArticleReaderSearch{
		ArticleID:     "",
		LastArticleID: paging.LastID,
	})

	if err != nil {
		p.Present(struct {
			Answer GetArticleResults
			Err    error
		}{Err: err})
		return
	}

	var lastId ArticleID
	if len(articles) > 0 {
		lastId = articles[len(articles)-1].ID
	}
	p.Present(struct {
		Answer GetArticleResults
		Err    error
	}{Answer: struct {
		Articles []Article
		Next     ArticleID
	}{
		Articles: articles, Next: lastId},
		Err: err,
	})
}
