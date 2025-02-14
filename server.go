package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"streamserver/internal/download"
	"streamserver/internal/google"
	"streamserver/internal/soundcloud"
)

func Server() http.Handler {
	mux := http.NewServeMux()

	soundcloud, err := soundcloud.New()
	if err != nil {
		log.Fatal(err)
	}

	mux.HandleFunc("/yt/{query}", func(w http.ResponseWriter, r *http.Request) {
		query := r.PathValue("query")

		if query == "" {
			http.Error(w, "invalid query", http.StatusNotAcceptable)
			return
		}

		results, err := google.Search(query)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(results)
	})

	mux.HandleFunc("/sc/{query}", func(w http.ResponseWriter, r *http.Request) {
		query := r.PathValue("query")

		if query == "" {
			http.Error(w, "invalid query", http.StatusNotAcceptable)
			return
		}

		results, err := soundcloud.Search(query)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(results)
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithCancel(r.Context())
		defer cancel()

		url := r.URL.Query().Get("url")
		if url == "" {
			http.Error(w, "invalid url", http.StatusNotAcceptable)
			return
		}

		w.Header().Set("Content-Type", "audio/mpeg")
		w.Header().Set("Transfer-Encoding", "chunked")

		if err := download.Download(url, ctx, w); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	return mux
}
