package domain

import (
	"strings"
	"time"
)

type ArticleID string

type Article struct {
	ID          ArticleID
	Title       string
	Content     string
	CreatedDate *time.Time
	Author      string
}

func (p ArticleID) String() string {
	return string(p)
}

func IsValidID(id ArticleID) bool {
	return strings.TrimSpace(id.String()) != ""
}
