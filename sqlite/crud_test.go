package sqlite

import (
	"context"
	_ "embed"
	"fmt"
	"log"
	"os"
	"simple-blog/domain"
	"testing"
	"time"
)

var blogCrud *BlogCrud

//go:embed init.sql
var initFile string

func init() {
	fileName := "test.db"
	_ = os.Remove(fileName)
	log.Printf("Creating %s...\n", fileName)
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	_ = file.Close()
	if _, err = InitDB(fmt.Sprintf("./%s", fileName)); err != nil {
		panic(err)
	}
	if _, err = db.Exec(initFile); err != nil {
		panic(err)
	}
}

func newCrud() {
	tx, err := db.BeginTx(context.TODO(), nil)
	if err != nil {
		panic(err)
	}
	blogCrud = &BlogCrud{
		Tx: tx,
	}
}

func TestName(t *testing.T) {
	newCrud()
	id, err := blogCrud.Write(struct {
		PostID      domain.ArticleID
		Title       string
		Content     string
		CreatedDate *time.Time
		Author      string
	}{
		PostID:      "",
		Title:       "My First SQLite Post",
		Content:     "This is a <strong>trusted</strong> HTML file",
		CreatedDate: nil,
		Author:      "J Patrick Davenport"})

	if err != nil {
		t.Fatalf("failed to write the post")
	}

	if id != "1" {
		t.Fatalf("invalid id for blog post %s", id)
	}
}
