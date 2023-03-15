package web

import (
	"fmt"
	"github.com/deusdat/cleango"
	"net/http"
	"simple-blog/domain"
	"simple-blog/web/templates"
)

type PostArticlePresenter struct {
	Writer         http.ResponseWriter
	Request        *http.Request
	ErrorPresenter cleango.Presenter[domain.Article]
}

func (b *PostArticlePresenter) Present(answer cleango.Output[domain.Article]) {
	if answer.Err != nil {
		b.ErrorPresenter.Present(answer)
		return
	}

	http.Redirect(
		b.Writer,
		b.Request,
		fmt.Sprintf("/articles/%s", answer.Answer.ID),
		301)
}

type GetArticlesPresenter struct {
	Writer  http.ResponseWriter
	Request *http.Request
}

func (p *GetArticlesPresenter) Present(answer cleango.Output[domain.GetArticleResults]) {
	type ForDisplay struct {
		Title string
		ID    string
	}

	var forDisplays []ForDisplay
	for _, article := range answer.Answer.Articles {
		forDisplays = append(forDisplays, ForDisplay{
			Title: article.Title,
			ID:    string(article.ID),
		})
	}
	midPoint := len(forDisplays)
	for i := 0; i < midPoint/2; i++ {
		j := midPoint - i - 1
		forDisplays[j], forDisplays[i] = forDisplays[i], forDisplays[j]
	}
	input := make(map[string]interface{})
	input["Articles"] = forDisplays
	input["NextID"] = answer.Answer.Next

	p.Writer.WriteHeader(200)
	p.Writer.Header().Set("Content-Type", "text/html")
	t := templates.Templates["articles.gohtml"]
	if err := t.Execute(p.Writer, input); err != nil {
		println("failed to write blogs")
	}
}
