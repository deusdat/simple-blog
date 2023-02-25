package web

import (
	_ "embed"
	"net/http"
)

//go:embed factory.html
var body string

func FactoryError(writer http.ResponseWriter, request *http.Request) {
	_, _ = writer.Write([]byte(body))
}
