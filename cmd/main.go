package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"simple-blog/web"
)

func main() {
	dbLocation := "test.db"
	//getFactory := func(ctx context.Context) factory.Factory {
	//	f, _ := factory.NewFactory(ctx, dbLocation)
	//	return f
	//}
	wrap := func(nf web.NeedsFactory) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			f := newFactory(dbLocation, w, r)
			nf(f)(w, r)
		}
	}
	r := chi.NewRouter()
	r.Route("/blogs", func(r chi.Router) {
		r.Get("/", wrap(web.GetArticles))
		r.Route("/{articleId}", func(r chi.Router) {
			r.Use(web.ArticleCtx)
			r.Post("/", wrap(web.PostArticle))
			r.Get("/edit", wrap(web.GetArticle))
		})
	})
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		return
	}
}
