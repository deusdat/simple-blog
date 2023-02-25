package domain

// ArticleWriter defines a means to persist a post.
type ArticleWriter interface {
	Write(post Article) (ArticleID, error)
}

type ArticleReader interface {
	Read(id ArticleID) (Article, error)
}

type ArticleDeleter interface {
	Delete(ID ArticleID) error
}
