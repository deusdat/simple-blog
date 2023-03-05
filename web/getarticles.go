package web

import (
	"net/http"
	"simple-blog/domain"
	"strconv"
)

func GetArticles(f Factory) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uc := f.GetArticlesUseCase()
		p := f.GetArticlesPresenter()
		nextID := r.URL.Query().Get("next")
		page := r.URL.Query().Get("page")

		nextPage, _ := strconv.Atoi(page)
		uc.Execute(domain.ArticlePaging{
			Number: nextPage,
			LastID: domain.ArticleID(nextID),
		}, p)
	}
}
