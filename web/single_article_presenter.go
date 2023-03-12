package web

import (
	"errors"
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
	Editing bool
}

func (p *GetSingleArticlePresenter) Present(answer cleango.Output[domain.Article]) {
	type ForDisplay struct {
		Title   string
		Content template.HTML
		ID      string
		Author  string
		Created time.Time
	}

	p.Writer.WriteHeader(200)
	p.Writer.Header().Set("Content-Type", "text/html")
	var err error
	if p.Editing {
		data := make(map[string]interface{})
		if answer.Err != nil {
			var domainErr *cleango.DomainError
			if !errors.As(answer.Err, &domainErr) ||
				domainErr.Kind != cleango.InvalidInput {
				http.Redirect(
					p.Writer,
					p.Request,
					fmt.Sprintf("/articles/error"),
					301)
				return
			}
			
		}
		data["Article"] = answer.Answer
		err = templates.Templates["edit.gohtml"].Execute(p.Writer, data)
	} else {
		err = templates.Templates["article.gohtml"].Execute(p.Writer, answer.Answer)
	}

	if err != nil {
		println("failed to write blogs", err.Error())
	}
}
