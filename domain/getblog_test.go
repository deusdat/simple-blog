package domain

import (
	"errors"
	"fmt"
	"github.com/deusdat/cleango"
	"testing"
)

type PresenterLogic[T any] func(cleango.Output[T])
type TestingPresenter struct {
	logic PresenterLogic[Article]
}

func (t TestingPresenter) Present(answer cleango.Output[Article]) {
	t.logic(answer)
}

type articleReader struct {
	err      error
	articles []Article
}

func (a articleReader) Read(paging ArticleReaderSearch) ([]Article, error) {
	if a.err != nil {
		return nil, a.err
	}
	return a.articles, nil
}

func TestGetSingleArticleUseCase_Execute(t *testing.T) {
	type expected func(answer cleango.Output[Article])
	type TestRecords struct {
		reader    ArticleReader
		presenter PresenterLogic[Article]
	}
	tests := []TestRecords{
		{
			reader: articleReader{
				err: cleango.ToDomainError("system err %w", nil),
			},
			presenter: func(output cleango.Output[Article]) {
				if output.Err == nil {
					t.Fatalf("should have produced error\n")
				}
				var domainErr *cleango.DomainError
				if !errors.As(output.Err, &domainErr) ||
					domainErr.Kind != cleango.System {
					fmt.Printf("error was %s", output.Err)
					t.Fatalf("did not get expected error message")
				}
			},
		},
	}
	for _, t := range tests {
		useCase := &GetSingleArticleUseCase{ArticleReader: t.reader}
		presenter := TestingPresenter{PresenterLogic[Article](t.presenter)}
		useCase.Execute(ArticleID("yup"), presenter)
	}
}
