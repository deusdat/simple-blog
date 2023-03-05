package domain

// ArticleWriter defines a means to persist a post.
type ArticleWriter interface {
	Write(post Article) (ArticleID, error)
}

type ArticleReaderSearch struct {
	ArticleID     ArticleID
	Page          int
	LastArticleID ArticleID
}
type ArticleReader interface {
	read(paging ArticlePaging)
}

type ArticleDeleter interface {
	Delete(ID ArticleID) error
}
