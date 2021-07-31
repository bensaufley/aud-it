package graphiql

import (
	_ "embed"

	"net/http"
)

//go:embed index.html
var html []byte

var Serve = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write(html)
})
