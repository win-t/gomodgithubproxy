package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	account := os.Args[1]
	panic(http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.EscapedPath()
		path = strings.TrimPrefix(path, "/")

		if strings.Contains(path, "/") {
			http.Error(w, "Not found", 404)
			return
		}

		if strings.Contains(path, `"`) {
			http.Error(w, "Bad request", 400)
			return
		}

		w.Header().Set("Cache-Control", "max-age=31536000, public, immutable")

		if r.URL.Query().Get("go-get") != "1" {
			http.Redirect(w, r, "https://github.com/"+account+"/"+path, 302)
			return
		}

		fmt.Fprintf(w, `<html><head><meta name="go-import" content="%s/%s git https://github.com/%s/%s"></head></html>`,
			r.Host, path,
			account, path,
		)
	})))
}
