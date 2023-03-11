package domain

// ArticleWriter defines a means to persist a post.
type ArticleWriter interface {
	Write(post Article) (ArticleID, error)
}

type ArticleReaderSearch struct {
	ArticleID     ArticleID
	LastArticleID ArticleID
}
type ArticleReader interface {
	Read(paging ArticleReaderSearch) ([]Article, error)
}

type ArticleDeleter interface {
	Delete(ID ArticleID) error
}
