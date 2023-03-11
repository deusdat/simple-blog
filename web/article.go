package web

import (
	"context"
	"fmt"
	"github.com/deusdat/cleango"
	"github.com/go-chi/chi/v5"
	"net/http"
	"simple-blog/domain"
)

const articleKey = "articleID"

const ctxArticleIDKey = "articleID"

func ArticleCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		articleID := chi.URLParam(r, articleKey)
		ctx := context.WithValue(r.Context(), ctxArticleIDKey, articleID)
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
		useCase := f.GetSingleArticleUseCase()
		p := f.GetSingleArticlePresenter()
		articleID := request.Context().Value(ctxArticleIDKey)
		if articleID == nil || articleID == "" {
			p.Present(struct {
				Answer domain.Article
				Err    error
			}{Err: &cleango.DomainError{
				Kind:    cleango.InvalidInput,
				Message: "articleID missing",
			}})
			return
		}

		useCase.Execute(domain.ArticleID(fmt.Sprintf("%v", articleID)), p)
	}
}
