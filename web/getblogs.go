package web

import (
	"net/http"
	"simple-blog/factory"
)

func GetArticles(dbLocation string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// This is a code smell due to the implementation of the factory,
		// not clean architecture. If I don't mind leaking the IO library, and I
		// don't since I'm in another IO layer, I could get the values from the context.
		f, _ := factory.NewFactory(r.Context(), dbLocation)
		uc := f.NewGetArticles()
		p := f.NewGetArticlesPresenter(w, r)
		uc.Execute(nil, p)
	}
}
