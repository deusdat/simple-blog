package web

import (
	"database/sql"
	"fmt"
	"github.com/deusdat/cleango"
	"html/template"
	"net/http"
	"simple-blog/domain"
	"simple-blog/web/templates"
	"time"
)

type BlogEditSource int

const (
	NewPost BlogEditSource = iota
	EditPost
)

type BlogPostPresenter struct {
	Tx                       *sql.Tx
	OriginalSource           BlogEditSource
	Writer                   http.ResponseWriter
	Request                  *http.Request
	ValidationErrorPresenter GetEditBlogPostPresenter
}

func (b *BlogPostPresenter) Present(answer cleango.Output[domain.Article]) {
	if answer.Err != nil {
		b.ValidationErrorPresenter.Present(answer)
		return
	}

	_ = b.Tx.Commit()
	http.Redirect(
		b.Writer,
		b.Request,
		fmt.Sprintf("/blogs/"),
		301)
}

type GetEditBlogPostPresenter struct {
}

func (b *GetEditBlogPostPresenter) Present(answer cleango.Output[domain.Article]) {

}

type GetArticlesPresenter struct {
	Writer  http.ResponseWriter
	Request *http.Request
}

func (p *GetArticlesPresenter) Present(answer cleango.Output[domain.GetArticleResults]) {
	type ForDisplay struct {
		Title   string
		Content template.HTML
		ID      string
		Author  string
		Created time.Time
	}

	var forDisplays []ForDisplay
	for _, article := range answer.Answer.Articles {
		createdOn := time.Now()
		if article.CreatedDate != nil {
			createdOn = *article.CreatedDate
		}
		forDisplays = append(forDisplays, ForDisplay{
			Title:   article.Title,
			Content: template.HTML(article.Content),
			ID:      string(article.ID),
			Author:  article.Author,
			Created: createdOn,
		})
	}
	midPoint := len(forDisplays)
	for i := 0; i < midPoint/2; i++ {
		j := midPoint - i - 1
		forDisplays[j], forDisplays[i] = forDisplays[i], forDisplays[j]
	}
	input := make(map[string]interface{})
	input["Articles"] = forDisplays

	p.Writer.WriteHeader(200)
	p.Writer.Header().Set("Content-Type", "text/html")
	t := templates.Templates["blogs.gohtml"]
	if err := t.Execute(p.Writer, input); err != nil {
		println("failed to write blogs")
	}
}
