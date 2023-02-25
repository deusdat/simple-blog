package domain

import (
	"github.com/deusdat/cleango"
)

type DeletePost struct {
	deleter ArticleDeleter
}

func (d DeletePost) Execute(input ArticleID, p cleango.Presenter[ArticleID]) {
	if !IsValidID(input) {
		p.Present(struct {
			Answer ArticleID
			Err    error
		}{Answer: "", Err: &cleango.DomainError{
			Kind:            cleango.InvalidInput,
			Message:         "invalid",
			UnderlyingCause: nil,
			Issues: []cleango.ValidationIssue{
				{Path: "input", Message: "invalid ArticleID"},
			},
		}})
		return
	}
	err := d.deleter.Delete(input)
	p.Present(struct {
		Answer ArticleID
		Err    error
	}{Answer: input, Err: err})
}
