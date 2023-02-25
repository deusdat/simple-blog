package domain

import (
	"github.com/deusdat/cleango"
	"testing"
)

type GeneralAnswer[T any] struct {
	Answer T
	Error  error
}

type TestDeleter struct {
	err error
}

func (t *TestDeleter) write(ID ArticleID) error {
	return t.err
}

func (g *GeneralAnswer[T]) Present(answer cleango.Output[T]) {
	g.Error = answer.Err
	g.Answer = answer.Answer
}

func TestDelete(t *testing.T) {
	var expected error = &cleango.DomainError{
		Kind:            cleango.NotFound,
		Message:         "unknown blogpost",
		UnderlyingCause: nil,
		Issues:          nil,
	}

	deleter := &TestDeleter{
		err: expected,
	}
	useCase := DeletePost{
		deleter: deleter,
	}

	p := &GeneralAnswer[ArticleID]{}
	id := ArticleID("hello")
	useCase.Execute(id, p)
	if p.Error != expected {
		t.Error("should have matched expected %w", p.Error)
	}

	deleter.err = nil
	p.Error = nil
	useCase.Execute(id, p)
	if p.Error != nil {
		t.Error("should not have an err %w", p.Error)
	}
	if p.Answer != id {
		t.Error("should have matched ids %w", p.Answer)
	}

	useCase.Execute(ArticleID(""), p)
	if p.Error == nil {
		t.Error("should not have an err %w", p.Error)
	}
	err, ok := p.Error.(*cleango.DomainError)
	if !ok {
		t.Fatalf("wrong type of error %s", err)
		return
	}
	if err.Kind != cleango.InvalidInput {
		t.Errorf("invalid kind %d", err.Kind)
	}
}
