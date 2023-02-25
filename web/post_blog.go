package web

import (
	"context"
	"github.com/go-chi/chi/v5"
	"net/http"
	"simple-blog/domain"
	"simple-blog/factory"
)

const articleKey = "articleID"

func ArticleCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		articleID := chi.URLParam(r, articleKey)
		ctx := context.WithValue(r.Context(), "articleId", articleID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func PostArticle(dbLocation string) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		f, err := factory.NewFactory(request.Context(), dbLocation)
		err = request.ParseForm()
		if err != nil {
			return
		}
		articleID, ok := request.Context().Value(articleKey).(string)
		if err != nil || !ok {
			FactoryError(writer, request)
			return
		}
		title := request.Form.Get("title")
		content := request.Form.Get("content")
		author := request.Form.Get("author")

		p := f.BlogPostPresenter(writer, request)
		post := f.NewCreateOrModifyPost()
		post.Execute(domain.Article{
			ID:          domain.ArticleID(articleID),
			Title:       title,
			Content:     content,
			CreatedDate: nil,
			Author:      author,
		}, p)
	}
}

func GetArticle(dbLocation string) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		_, err := factory.NewFactory(request.Context(), dbLocation)
		if err != nil {
			FactoryError(writer, request)
			return
		}
	}
}
