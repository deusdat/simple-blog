package web

import (
	"net/http"
	"simple-blog/domain"
)

func GetArticles(f Factory) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uc := f.GetArticlesUseCase()
		p := f.GetArticlesPresenter()
		nextID := r.URL.Query().Get("next")

		uc.Execute(domain.ArticlePaging{
			LastID: domain.ArticleID(nextID),
		}, p)
	}
}
