package main

import (
	"github.com/go-chi/chi/v5"
	"simple-blog/web"
)

func main() {
	dbLocation := "test.db"
	r := chi.NewRouter()
	r.Route("/blogs", func(r chi.Router) {
		r.Post("/", web.GetArticles(dbLocation))
		r.Route("/{articleId}", func(r chi.Router) {
			r.Use(web.ArticleCtx)
			r.Post("/", web.PostArticle(dbLocation))
			r.Get("/edit", web.GetArticle(dbLocation))
		})
	})
}
