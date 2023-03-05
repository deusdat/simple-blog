package web

import (
	"context"
	"github.com/go-chi/chi/v5"
	"net/http"
)

const articleKey = "articleID"

func ArticleCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		articleID := chi.URLParam(r, articleKey)
		ctx := context.WithValue(r.Context(), "articleId", articleID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func PostArticle(f Factory) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		//if err := request.ParseForm(); err != nil {
		//	FactoryError(writer, request)
		//	return
		//}
		//articleID, ok := request.Context().Value(articleKey).(string)
		//if !ok {
		//	FactoryError(writer, request)
		//	return
		//}
		//title := request.Form.Get("title")
		//content := request.Form.Get("content")
		//author := request.Form.Get("author")
		//
		//p := f.BlogPostPresenter(writer, request)
		//post := f.NewCreateOrModifyPost()
		//post.Execute(domain.Article{
		//	ID:          domain.ArticleID(articleID),
		//	Title:       title,
		//	Content:     content,
		//	CreatedDate: nil,
		//	Author:      author,
		//}, p)
	}
}

func GetArticle(f Factory) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		//f := newF(request.Context())
		//if f != nil {
		//	FactoryError(writer, request)
		//	return
		//}
	}
}
