package domain

import (
	"github.com/deusdat/cleango"
	"time"
)

type GetArticleUseCase struct {
}

type GetArticleResults struct {
	Articles []Article
	Next     ArticleID
}

func (g GetArticleUseCase) Execute(articleID *ArticleID, p cleango.Presenter[GetArticleResults]) {
	firstDt := time.Now().AddDate(0, -1, 0)
	secondDt := time.Now()

	out := cleango.Output[GetArticleResults]{
		Answer: GetArticleResults{
			Articles: []Article{
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
					Title: "My most recent work",
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
			},
			Next: "",
		},
		Err: nil,
	}
	p.Present(out)
}