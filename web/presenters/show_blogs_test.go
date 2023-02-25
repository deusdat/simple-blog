package presenters

import (
	"io"
	"net/http"
	"net/http/httptest"
	"simple-blog/domain"
	"testing"
	"time"
)

func TestNoArticlesFound(t *testing.T) {
	r, err := http.NewRequest("", "", nil)
	if err != nil {
		t.Fatalf("couldn't create request")
	}
	w := httptest.NewRecorder()
	p := GetArticlesPresenter{
		Writer:  w,
		Request: r,
	}

	p.Present(struct {
		Answer domain.GetArticleResults
		Err    error
	}{Answer: struct {
		Articles []domain.Article
		Next     domain.ArticleID
	}{Articles: nil}})

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)
	asStr := string(body)
	if `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title></title>
</head>
<body>
    
        <h2>No blogs found</h2>
    
</body>
</html>` != asStr {
		t.Fatalf("didn't generate correctly.")
	}
}

func TestSomeArticlesFound(t *testing.T) {
	r, err := http.NewRequest("", "", nil)
	if err != nil {
		t.Fatalf("couldn't create request")
	}
	w := httptest.NewRecorder()
	p := GetArticlesPresenter{
		Writer:  w,
		Request: r,
	}
	oneTime := time.Now()
	p.Present(struct {
		Answer domain.GetArticleResults
		Err    error
	}{Answer: struct {
		Articles []domain.Article
		Next     domain.ArticleID
	}{Articles: []domain.Article{
		{
			ID:          "asfsdf",
			Title:       "One",
			Content:     "Content doesn't have to be long Google. You're making recipes sites make crap.",
			CreatedDate: &oneTime,
			Author:      "JPD",
		},
	}}})

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)
	asStr := string(body)
	if `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title></title>
</head>
<body>
    
        
            <section>
                <h1>One</h1>
                <div>Content doesn't have to be long Google. You're making recipes sites make crap.</div>
            </section>
        
    
</body>
</html>` != asStr {
		t.Fatalf("didn't generate correctly.")
	}
}
