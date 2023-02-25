package domain

import (
	"fmt"
	"github.com/deusdat/cleango"
	"strings"
)

type CreateOrModifyPost struct {
	Writer ArticleWriter
}

func ValidateBlogPost(post Article) *cleango.DomainError {
	var issues []cleango.ValidationIssue
	if strings.TrimSpace(post.Title) == "" {
		issues = append(issues, cleango.ValidationIssue{
			Path:    "/Title",
			Message: "must exist",
		})
	}

	if strings.TrimSpace(post.Content) == "" {
		issues = append(issues, cleango.ValidationIssue{
			Path:    "/Content",
			Message: "must exist",
		})
	}

	if strings.TrimSpace(post.Author) == "" {
		issues = append(issues, cleango.ValidationIssue{
			Path:    "/Author",
			Message: "must exist",
		})
	}

	if len(issues) > 0 {
		return &cleango.DomainError{
			Kind:            cleango.InvalidInput,
			Message:         "invalid blog post",
			UnderlyingCause: nil,
			Issues:          issues,
		}
	}
	return nil
}

func (c *CreateOrModifyPost) Execute(input Article, p cleango.Presenter[Article]) {
	op := "CreateOrModifyPost - %w"
	if err := ValidateBlogPost(input); err != nil {
		p.Present(cleango.Output[Article]{
			Answer: input,
			Err:    fmt.Errorf(op, err),
		})
		return
	}
	id, writeErr := c.Writer.Write(input)
	if writeErr != nil {
		p.Present(cleango.Output[Article]{
			Answer: input,
			Err:    fmt.Errorf(op, writeErr),
		})
		return
	}
	input.ID = id
	p.Present(cleango.Output[Article]{
		Answer: input,
		Err:    nil,
	})
	return // just for good measure.
}
