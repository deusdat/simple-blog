package main

import (
	"github.com/deusdat/cleango"
	"net/http"
	"simple-blog/domain"
	"simple-blog/web"
)

type factory struct {
	w http.ResponseWriter
	r *http.Request
}

func (f factory) GetArticlesUseCase() cleango.UseCase[domain.ArticlePaging, domain.GetArticleResults] {
	return domain.GetArticleUseCase{}
}

func (f factory) GetArticlesPresenter() cleango.Presenter[domain.GetArticleResults] {
	return &web.GetArticlesPresenter{
		Writer:  f.w,
		Request: f.r,
	}
}

func newFactory(dbLocation string, w http.ResponseWriter, r *http.Request) factory {
	return factory{
		w: w,
		r: r,
	}
}
