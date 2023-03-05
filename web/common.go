package web

import "net/http"

type NeedsFactory func(f Factory) http.HandlerFunc
