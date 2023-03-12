package memory

import (
	"github.com/google/uuid"
	"simple-blog/domain"
	"sync"
)

type ArticleCache struct {
	sync.Mutex
	// a slice is easy to reason about for a demo. Also preserves order.
	articles []domain.Article
}

func (a *ArticleCache) Write(post domain.Article) (domain.ArticleID, error) {
	a.Mutex.Lock()
	defer a.Mutex.Unlock()
	for i, _ := range a.articles {
		art := &a.articles[i]
		if art.ID == post.ID {
			art.Title, art.CreatedDate, art.Author, art.Content =
				post.Title, post.CreatedDate, post.Author, post.Content
			return post.ID, nil
		}
	}
	post.ID = domain.ArticleID(uuid.New().String())
	a.articles = append(a.articles, post)
	return post.ID, nil
}

func (a *ArticleCache) Read(paging domain.ArticleReaderSearch) ([]domain.Article, error) {
	a.Mutex.Lock()
	defer a.Mutex.Unlock()
	// Get the unique article.
	if domain.IsValidID(paging.ArticleID) {
		for _, art := range a.articles {
			if art.ID == paging.ArticleID {
				return []domain.Article{art}, nil
			}
		}
	}
	var asCopy []domain.Article
	pageSize := 3
	for _, art := range a.articles {
		if art.ID == paging.ArticleID {
			asCopy = []domain.Article{art}
			break
		}
		if domain.IsValidID(paging.LastArticleID) {
			if art.ID == paging.LastArticleID {
				asCopy = nil
			}
			// we need to start paging at article.
			asCopy = append(asCopy, art)
		} else {
			// we want the first pageSize
			asCopy = append(asCopy, art)
		}

		if len(asCopy) > pageSize {
			break
		}
	}
	return asCopy, nil
}
