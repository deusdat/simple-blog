package sqlite

import (
	"database/sql"
	"fmt"
	"github.com/deusdat/cleango"
	_ "modernc.org/sqlite"
	"simple-blog/domain"
	"time"
)

var db *sql.DB

// CloseDb should always be called when the app shuts down.
func CloseDb() {
	if db == nil {
		return
	}
	err := db.Close()
	if err != nil {
		println(err.Error())
		return
	}
}

func InitDB(dbLocation string) (*sql.DB, error) {
	var err error
	if db == nil {
		db, err = sql.Open("sqlite", dbLocation)
		if err != nil {
			return nil, err
		}

	}
	return db, nil
}

type BlogCrud struct {
	Tx *sql.Tx
}

func (s *BlogCrud) insert(post domain.Article) (domain.ArticleID, error) {
	st := `INSERT INTO blog (title, content, created_dt, author)
			VALUES ($1, $2, $3, $4)`

	dtStr := (*post.CreatedDate).Format(time.RFC3339)
	rs, err := s.Tx.Exec(st, post.Title, post.Content, dtStr, post.Author)
	if err != nil {
		return "", &cleango.DomainError{
			Kind:            cleango.System,
			Message:         "failed to write blog",
			UnderlyingCause: err,
			Issues:          nil,
		}
	}

	id, err := rs.LastInsertId()
	if err != nil {
		return "", &cleango.DomainError{
			Kind:            cleango.System,
			Message:         "failed to get blog id",
			UnderlyingCause: err,
			Issues:          nil,
		}
	}

	return domain.ArticleID(fmt.Sprintf("%d", id)), nil
}

func (s *BlogCrud) Write(post domain.Article) (domain.ArticleID, error) {
	if post.CreatedDate == nil {
		now := time.Now()
		post.CreatedDate = &now
	}

	if post.ID == "" {
		return s.insert(post)
	}
	return s.update(post)
}

func (s *BlogCrud) update(_ domain.Article) (domain.ArticleID, error) {
	return "", &cleango.DomainError{
		Kind:            cleango.System,
		Message:         "update not implemented",
		UnderlyingCause: nil,
		Issues:          nil,
	}
}
