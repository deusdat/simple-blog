package web

import (
	"fmt"
	"github.com/deusdat/cleango"
	"html/template"
	"net/http"
	"simple-blog/domain"
	"simple-blog/web/templates"
	"time"
)

type GetSingleArticlePresenter struct {
	Writer  http.ResponseWriter
	Request *http.Request
}

func (p *GetSingleArticlePresenter) Present(answer cleango.Output[domain.Article]) {
	type ForDisplay struct {
		Title   string
		Content template.HTML
		ID      string
		Author  string
		Created time.Time
	}
	if answer.Err != nil {
		http.Redirect(
			p.Writer,
			p.Request,
			fmt.Sprintf("/articles/error"),
			301)
		return
	}

	p.Writer.WriteHeader(200)
	p.Writer.Header().Set("Content-Type", "text/html")
	t := templates.Templates["article.gohtml"]
	if err := t.Execute(p.Writer, answer.Answer); err != nil {
		println("failed to write blogs")
	}
}
