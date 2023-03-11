package memory

import (
	"simple-blog/domain"
	"testing"
	"time"
)

func TestArticleCache(t *testing.T) {
	cache := ArticleCache{}
	articles, err := cache.Read(struct {
		ArticleID     domain.ArticleID
		LastArticleID domain.ArticleID
	}{ArticleID: "One", LastArticleID: ""})

	if err != nil {
		t.Fatalf("1. Failed to search for articles\n")
	}
	if len(articles) != 0 {
		t.Fatalf("2. There should be no articles at this point\n")
	}

	now := time.Now()
	articleID, err := cache.Write(struct {
		ID          domain.ArticleID
		Title       string
		Content     string
		CreatedDate *time.Time
		Author      string
	}{
		ID:          "One",
		Title:       "My first test article.",
		Content:     "Hello, world!",
		CreatedDate: &now,
		Author:      "Patrick"})
	if err != nil {
		t.Fatalf("Failed to write first article\n")
	}
	if articleID == "" {
		t.Fatalf("articleID is wrong %s.\n", articleID)
	}

	secondID, err := cache.Write(struct {
		ID          domain.ArticleID
		Title       string
		Content     string
		CreatedDate *time.Time
		Author      string
	}{
		Title:       "My second test article.",
		Content:     "Good bye, commute!",
		CreatedDate: &now,
		Author:      "Patrick"})
	if err != nil {
		t.Fatalf("Failed to write second article\n")
	}

	articles, err = cache.Read(struct {
		ArticleID     domain.ArticleID
		LastArticleID domain.ArticleID
	}{ArticleID: articleID, LastArticleID: ""})

	if len(articles) != 1 ||
		articles[0].Title != "My first test article." ||
		articles[0].Author != "Patrick" ||
		articles[0].Content != "Hello, world!" {
		t.Fatalf("Failed while searching for first article\n")
	}

	articles, err = cache.Read(struct {
		ArticleID     domain.ArticleID
		LastArticleID domain.ArticleID
	}{ArticleID: secondID, LastArticleID: ""})

	if len(articles) != 1 ||
		articles[0].ID != secondID ||
		articles[0].Title != "My second test article." ||
		articles[0].Author != "Patrick" ||
		articles[0].Content != "Good bye, commute!" {
		t.Fatalf("Failed while searching for second article\n")
	}
}
