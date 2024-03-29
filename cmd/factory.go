package main

import (
	"github.com/deusdat/cleango"
	"net/http"
	"simple-blog/domain"
	"simple-blog/memory"
	"simple-blog/web"
	"time"
)

type factory struct {
	w http.ResponseWriter
	r *http.Request
}

func (f factory) PostArticlePresenter() cleango.Presenter[domain.Article] {
	return &web.PostArticlePresenter{
		Writer:         f.w,
		Request:        f.r,
		ErrorPresenter: f.GetEditArticlePresenter(),
	}
}

func (f factory) PostEditArticleUseCase() cleango.UseCase[domain.Article, domain.Article] {
	return &domain.CreateOrModifyPost{
		Writer: articleCache,
	}
}

func (f factory) GetEditArticlePresenter() cleango.Presenter[domain.Article] {
	return &web.GetSingleArticlePresenter{
		Writer:  f.w,
		Request: f.r,
		Editing: true,
	}
}

func (f factory) GetSingleArticleUseCase() cleango.UseCase[domain.ArticleID, domain.Article] {
	return &domain.GetSingleArticleUseCase{ArticleReader: articleCache}
}

var articleCache *memory.ArticleCache

func init() {
	articleCache = &memory.ArticleCache{}
	now := time.Now()
	firstDt := now.AddDate(0, -1, 10)
	secondDt := now.AddDate(0, -1, 0)
	thirdDt := now.AddDate(0, 0, -1)
	fourth := now

	for _, article := range []domain.Article{
		{
			ID:    "random1",
			Title: "My First article",
			Content: `<div><b>Sometimes you just need to have a test article.</b></div>
<div>Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor 
incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud 
exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure 
dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. 
Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit 
anim id est laborum.</div>
`,
			CreatedDate: &firstDt,
			Author:      "Milton Bradly",
		},
		{
			ID:    "randomw2",
			Title: "My Second Article",
			Content: `Sed ut perspiciatis unde omnis iste natus error sit 
voluptatem accusantium doloremque laudantium, totam rem aperiam, eaque ipsa quae ab illo 
inventore veritatis et quasi architecto beatae vitae dicta sunt explicabo. Nemo enim ipsam
voluptatem quia voluptas sit aspernatur aut odit aut fugit, sed quia consequuntur magni
dolores eos qui ratione voluptatem sequi nesciunt. Neque porro quisquam est, qui dolorem
ipsum quia dolor sit amet, consectetur, adipisci velit, sed quia non numquam eius modi
tempora incidunt ut labore et dolore magnam aliquam quaerat voluptatem. Ut enim ad minima
veniam, quis nostrum exercitationem ullam corporis suscipit laboriosam, nisi ut aliquid ex 
ea commodi consequatur? Quis autem vel eum iure reprehenderit qui in ea voluptate velit esse
quam nihil molestiae consequatur, vel illum qui dolorem eum fugiat quo voluptas nulla pariatur?`,
			CreatedDate: &secondDt,
			Author:      "Tiler Durden",
		},
		{
			ID:    "randomw3",
			Title: "My Third",
			Content: `Sed ut perspiciatis unde omnis iste natus error sit 
voluptatem accusantium doloremque laudantium, totam rem aperiam, eaque ipsa quae ab illo 
inventore veritatis et quasi architecto beatae vitae dicta sunt explicabo. Nemo enim ipsam
voluptatem quia voluptas sit aspernatur aut odit aut fugit, sed quia consequuntur magni
dolores eos qui ratione voluptatem sequi nesciunt. Neque porro quisquam est, qui dolorem
ipsum quia dolor sit amet, consectetur, adipisci velit, sed quia non numquam eius modi
tempora incidunt ut labore et dolore magnam aliquam quaerat voluptatem. Ut enim ad minima
veniam, quis nostrum exercitationem ullam corporis suscipit laboriosam, nisi ut aliquid ex 
ea commodi consequatur? Quis autem vel eum iure reprehenderit qui in ea voluptate velit esse
quam nihil molestiae consequatur, vel illum qui dolorem eum fugiat quo voluptas nulla pariatur?`,
			CreatedDate: &thirdDt,
			Author:      "Abe Froman",
		},
		{
			ID:    "randomw4",
			Title: "My Fourth",
			Content: `Sed ut perspiciatis unde omnis iste natus error sit 
voluptatem accusantium doloremque laudantium, totam rem aperiam, eaque ipsa quae ab illo 
inventore veritatis et quasi architecto beatae vitae dicta sunt explicabo. Nemo enim ipsam
voluptatem quia voluptas sit aspernatur aut odit aut fugit, sed quia consequuntur magni
dolores eos qui ratione voluptatem sequi nesciunt. Neque porro quisquam est, qui dolorem
ipsum quia dolor sit amet, consectetur, adipisci velit, sed quia non numquam eius modi
tempora incidunt ut labore et dolore magnam aliquam quaerat voluptatem. Ut enim ad minima
veniam, quis nostrum exercitationem ullam corporis suscipit laboriosam, nisi ut aliquid ex 
ea commodi consequatur? Quis autem vel eum iure reprehenderit qui in ea voluptate velit esse
quam nihil molestiae consequatur, vel illum qui dolorem eum fugiat quo voluptas nulla pariatur?`,
			CreatedDate: &fourth,
			Author:      "Mr Tibbs",
		},
	} {
		_, _ = articleCache.Write(article)
	}
}

func (f factory) GetArticlesUseCase() cleango.UseCase[domain.ArticlePaging, domain.GetArticleResults] {
	return &domain.GetArticleUseCase{
		ArticleReader: articleCache,
	}
}

func (f factory) GetArticlesPresenter() cleango.Presenter[domain.GetArticleResults] {
	return &web.GetArticlesPresenter{
		Writer:  f.w,
		Request: f.r,
	}
}

func newFactory(dbLocation string, w http.ResponseWriter, r *http.Request) factory {
	return factory{
		w: w,
		r: r,
	}
}

func (f factory) GetSingleArticlePresenter() cleango.Presenter[domain.Article] {
	return &web.GetSingleArticlePresenter{
		Writer:  f.w,
		Request: f.r,
		Editing: false,
	}
}
